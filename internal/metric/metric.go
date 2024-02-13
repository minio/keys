// Copyright 2020 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package metric

import (
	"io"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/minio/kes/internal/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/expfmt"
)

// New returns a new Metrics that gathers and exposes various
// metrics about the application.
func New() *Metrics {
	requestStatusLabels := []string{"code"}

	registry := prometheus.NewRegistry()
	promFactory := promauto.With(registry)

	metrics := &Metrics{
		gatherer: registry,
		requestSucceeded: promFactory.NewCounterVec(prometheus.CounterOpts{
			Namespace: "kes",
			Subsystem: "http",
			Name:      "request_success",
			Help:      "Number of requests that have been served successfully.",
		}, requestStatusLabels),
		requestErrored: promFactory.NewCounterVec(prometheus.CounterOpts{
			Namespace: "kes",
			Subsystem: "http",
			Name:      "request_error",
			Help:      "Number of request that failed due to some error. (HTTP 4xx status code)",
		}, requestStatusLabels),
		requestFailed: promFactory.NewCounterVec(prometheus.CounterOpts{
			Namespace: "kes",
			Subsystem: "http",
			Name:      "request_failure",
			Help:      "Number of request that failed due to some internal failure. (HTTP 5xx status code)",
		}, requestStatusLabels),
		requestActive: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "http",
			Name:      "request_active",
			Help:      "Number of active requests that are not finished, yet.",
		}),
		requestLatency: promFactory.NewHistogram(prometheus.HistogramOpts{
			Namespace: "kes",
			Subsystem: "http",
			Name:      "response_time",
			Buckets:   []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1.0, 1.5, 3.0, 5.0, 10.0}, // from 10ms to 10s
			Help:      "Histogram of request response times spawning from 10ms to 10s.",
		}),

		errorLogEvents: promFactory.NewCounter(prometheus.CounterOpts{
			Namespace: "kes",
			Subsystem: "log",
			Name:      "error_events",
			Help:      "Number of error log events written to the error log targets.",
		}),
		auditLogEvents: promFactory.NewCounter(prometheus.CounterOpts{
			Namespace: "kes",
			Subsystem: "log",
			Name:      "audit_events",
			Help:      "Number of audit log events written to the audit log targets.",
		}),

		startTime: time.Now(),
		upTimeInSeconds: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "up_time",
			Help:      "The time the server has been up and running in seconds.",
		}),

		numCPUs: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "num_cpu",
			Help:      "The number of logical CPUs available on the system. It may be larger than the number of usable CPUs.",
		}),
		numUsableCPUs: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "num_cpu_used",
			Help:      "The number of logical CPUs usable by the server. It may be smaller than the number of available CPUs.",
		}),
		numThreads: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "num_threads",
			Help:      "The number of concurrent co-routines/threads that currently exists.",
		}),

		memHeapUsed: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "mem_heap_used",
			Help:      "The number of bytes that are currently allocated on the heap memory.",
		}),
		memHeapObjects: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "mem_heap_objects",
			Help:      "The number of objects that are currently allocated on the heap memory.",
		}),
		memStackUsed: promFactory.NewGauge(prometheus.GaugeOpts{
			Namespace: "kes",
			Subsystem: "system",
			Name:      "mem_stack_used",
			Help:      "The bytes of stack memory obtained from the OS.",
		}),
	}

	return metrics
}

// Metrics is a type that gathers various metrics and information
// about an application.
type Metrics struct {
	gatherer prometheus.Gatherer

	requestSucceeded *prometheus.CounterVec
	requestFailed    *prometheus.CounterVec
	requestErrored   *prometheus.CounterVec
	requestActive    prometheus.Gauge
	requestLatency   prometheus.Histogram

	errorLogEvents prometheus.Counter
	auditLogEvents prometheus.Counter

	startTime       time.Time // Used to compute the up time as upTime = now - startTime
	upTimeInSeconds prometheus.Gauge
	numCPUs         prometheus.Gauge
	numUsableCPUs   prometheus.Gauge
	numThreads      prometheus.Gauge

	memHeapUsed    prometheus.Gauge
	memHeapObjects prometheus.Gauge
	memStackUsed   prometheus.Gauge
}

// EncodeTo collects all outstanding metrics information
// about the application and writes it to encoder.
func (m *Metrics) EncodeTo(encoder expfmt.Encoder) error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.upTimeInSeconds.Set(time.Since(m.startTime).Truncate(10 * time.Millisecond).Seconds())
	m.numCPUs.Set(float64(runtime.NumCPU()))
	m.numUsableCPUs.Set(float64(runtime.GOMAXPROCS(0)))
	m.numThreads.Set(float64(runtime.NumGoroutine()))
	m.memHeapUsed.Set(float64(memStats.HeapAlloc))
	m.memHeapObjects.Set(float64(memStats.HeapObjects))
	m.memStackUsed.Set(float64(memStats.StackSys))

	metrics, err := m.gatherer.Gather()
	if err != nil {
		return err
	}
	for _, metric := range metrics {
		if err := encoder.Encode(metric); err != nil {
			return err
		}
	}
	return nil
}

// Count returns a HandlerFunc that wraps h and counts the
// how many requests succeeded (HTTP 200 OK) and how many
// failed.
//
// Count distingushes requests that fail with some sort of
// well-defined error (HTTP 4xx) and requests that fail due
// to some internal error (HTTP 5xx).
func (m *Metrics) Count(h api.Handler) api.Handler {
	return api.HandlerFunc(func(resp *api.Response, req *api.Request) {
		m.requestActive.Inc()
		defer m.requestActive.Dec()

		rw := &countResponseWriter{
			ResponseWriter: resp.ResponseWriter,
			succeeded:      m.requestSucceeded,
			errored:        m.requestErrored,
			failed:         m.requestFailed,
		}
		if flusher, ok := resp.ResponseWriter.(http.Flusher); ok {
			rw.flusher = flusher
		}
		resp.ResponseWriter = rw
		h.ServeAPI(resp, req)
	})
}

// Latency returns a HandlerFunc that wraps h and measures the
// internal request-response latency.
//
// The internal request-response latency is the time the
// application takes to generate and send a response after
// receiving a request. It basically shows how many request
// the application can handle.
func (m *Metrics) Latency(h api.Handler) api.Handler {
	return api.HandlerFunc(func(resp *api.Response, req *api.Request) {
		rw := &latencyResponseWriter{
			ResponseWriter: resp.ResponseWriter,
			start:          time.Now(),
			histogram:      m.requestLatency,
		}
		if flusher, ok := rw.ResponseWriter.(http.Flusher); ok {
			rw.flusher = flusher
		}
		resp.ResponseWriter = rw
		h.ServeAPI(resp, req)
	})
}

// ErrorEventCounter returns an io.Writer that increments
// the error event log counter on each write call.
//
// The returned io.Writer never returns an error on writes.
func (m *Metrics) ErrorEventCounter() io.Writer {
	return eventCounter{metric: m.errorLogEvents}
}

// AuditEventCounter returns an io.Writer that increments
// the audit event log counter on each write call.
//
// The returned io.Writer never returns an error on writes.
func (m *Metrics) AuditEventCounter() io.Writer {
	return eventCounter{metric: m.auditLogEvents}
}

type eventCounter struct {
	metric prometheus.Counter
}

func (w eventCounter) Write(p []byte) (int, error) {
	w.metric.Inc()
	return len(p), nil
}

// latencyResponseWriter is an http.ResponseWriter that
// measures the internal request-response latency.
type latencyResponseWriter struct {
	http.ResponseWriter
	flusher http.Flusher

	start     time.Time            // The point in time when the request was received
	histogram prometheus.Histogram // The latency histogram
	written   bool                 // Inidicates whether the HTTP headers have been written
}

var (
	_ http.ResponseWriter = (*latencyResponseWriter)(nil)
	_ http.Flusher        = (*latencyResponseWriter)(nil)
)

func (w *latencyResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	if !w.written {
		w.histogram.Observe(time.Since(w.start).Seconds())
		w.written = true
	}
}

func (w *latencyResponseWriter) Flush() {
	if w.flusher != nil {
		w.flusher.Flush()
	}
}

// Unwrap returns the underlying http.ResponseWriter.
//
// This method is implemented for http.ResponseController.
func (w *latencyResponseWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

// countResponseWriter is an http.ResponseWriter that
// counts the number of requests partition by requests
// that:
//   - Succeeded (HTTP 200 OK)
//   - Errored   (HTTP 4xx)
//   - Failed    (HTTP 5xx)
type countResponseWriter struct {
	http.ResponseWriter
	flusher http.Flusher

	succeeded *prometheus.CounterVec
	errored   *prometheus.CounterVec
	failed    *prometheus.CounterVec
	prometheus.Metric
	written bool // Inidicates whether the HTTP headers have been written
}

var (
	_ http.ResponseWriter = (*countResponseWriter)(nil)
	_ http.Flusher        = (*countResponseWriter)(nil)
)

func (w *countResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	if !w.written {
		switch {
		case status >= 200 && status < 300:
			w.succeeded.WithLabelValues(strconv.Itoa(status)).Inc()
		case status >= 400 && status < 500:
			w.errored.WithLabelValues(strconv.Itoa(status)).Inc()
		case status >= 500 && status < 600:
			w.failed.WithLabelValues(strconv.Itoa(status)).Inc()
		default:
			// We panic to signal that the server returned a status code
			// that is not tracked. If, in the future, the application
			// returns a new (kind of) status code it should be collected
			// as well.
			// Otherwise, we would silently ignore new status codes and the
			// metrics would be incomplete.
			panic("metrics: unexpected response status code " + strconv.Itoa(status))
		}
		w.written = true
	}
}

func (w *countResponseWriter) Write(b []byte) (int, error) { return w.ResponseWriter.Write(b) }

func (w *countResponseWriter) Flush() {
	if w.flusher != nil {
		w.flusher.Flush()
	}
}

// Unwrap returns the underlying http.ResponseWriter.
//
// This method is implemented for http.ResponseController.
func (w *countResponseWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }
