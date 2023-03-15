// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO KES
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/minio/kes/models"
	"github.com/minio/kes/restapi/operations/auth"
	"github.com/minio/kes/restapi/operations/encryption"
)

// NewKesAPI creates a new Kes instance
func NewKesAPI(spec *loads.Document) *KesAPI {
	return &KesAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		EncryptionAPIsHandler: encryption.APIsHandlerFunc(func(params encryption.APIsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.APIs has not yet been implemented")
		}),
		EncryptionAssignPolicyHandler: encryption.AssignPolicyHandlerFunc(func(params encryption.AssignPolicyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.AssignPolicy has not yet been implemented")
		}),
		EncryptionCreateKeyHandler: encryption.CreateKeyHandlerFunc(func(params encryption.CreateKeyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.CreateKey has not yet been implemented")
		}),
		EncryptionDeleteIdentityHandler: encryption.DeleteIdentityHandlerFunc(func(params encryption.DeleteIdentityParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DeleteIdentity has not yet been implemented")
		}),
		EncryptionDeleteKeyHandler: encryption.DeleteKeyHandlerFunc(func(params encryption.DeleteKeyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DeleteKey has not yet been implemented")
		}),
		EncryptionDeletePolicyHandler: encryption.DeletePolicyHandlerFunc(func(params encryption.DeletePolicyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DeletePolicy has not yet been implemented")
		}),
		EncryptionDescribeIdentityHandler: encryption.DescribeIdentityHandlerFunc(func(params encryption.DescribeIdentityParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DescribeIdentity has not yet been implemented")
		}),
		EncryptionDescribeKeyHandler: encryption.DescribeKeyHandlerFunc(func(params encryption.DescribeKeyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DescribeKey has not yet been implemented")
		}),
		EncryptionDescribePolicyHandler: encryption.DescribePolicyHandlerFunc(func(params encryption.DescribePolicyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DescribePolicy has not yet been implemented")
		}),
		EncryptionDescribeSelfIdentityHandler: encryption.DescribeSelfIdentityHandlerFunc(func(params encryption.DescribeSelfIdentityParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.DescribeSelfIdentity has not yet been implemented")
		}),
		EncryptionGetPolicyHandler: encryption.GetPolicyHandlerFunc(func(params encryption.GetPolicyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.GetPolicy has not yet been implemented")
		}),
		EncryptionImportKeyHandler: encryption.ImportKeyHandlerFunc(func(params encryption.ImportKeyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.ImportKey has not yet been implemented")
		}),
		EncryptionListIdentitiesHandler: encryption.ListIdentitiesHandlerFunc(func(params encryption.ListIdentitiesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.ListIdentities has not yet been implemented")
		}),
		EncryptionListKeysHandler: encryption.ListKeysHandlerFunc(func(params encryption.ListKeysParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.ListKeys has not yet been implemented")
		}),
		EncryptionListPoliciesHandler: encryption.ListPoliciesHandlerFunc(func(params encryption.ListPoliciesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.ListPolicies has not yet been implemented")
		}),
		AuthLoginHandler: auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Login has not yet been implemented")
		}),
		AuthLoginDetailHandler: auth.LoginDetailHandlerFunc(func(params auth.LoginDetailParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginDetail has not yet been implemented")
		}),
		AuthLogoutHandler: auth.LogoutHandlerFunc(func(params auth.LogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.Logout has not yet been implemented")
		}),
		EncryptionMetricsHandler: encryption.MetricsHandlerFunc(func(params encryption.MetricsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.Metrics has not yet been implemented")
		}),
		AuthSessionCheckHandler: auth.SessionCheckHandlerFunc(func(params auth.SessionCheckParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.SessionCheck has not yet been implemented")
		}),
		EncryptionSetPolicyHandler: encryption.SetPolicyHandlerFunc(func(params encryption.SetPolicyParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.SetPolicy has not yet been implemented")
		}),
		EncryptionStatusHandler: encryption.StatusHandlerFunc(func(params encryption.StatusParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.Status has not yet been implemented")
		}),
		EncryptionVersionHandler: encryption.VersionHandlerFunc(func(params encryption.VersionParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation encryption.Version has not yet been implemented")
		}),

		KeyAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (key) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*KesAPI the kes API */
type KesAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// KeyAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	KeyAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// EncryptionAPIsHandler sets the operation handler for the a p is operation
	EncryptionAPIsHandler encryption.APIsHandler
	// EncryptionAssignPolicyHandler sets the operation handler for the assign policy operation
	EncryptionAssignPolicyHandler encryption.AssignPolicyHandler
	// EncryptionCreateKeyHandler sets the operation handler for the create key operation
	EncryptionCreateKeyHandler encryption.CreateKeyHandler
	// EncryptionDeleteIdentityHandler sets the operation handler for the delete identity operation
	EncryptionDeleteIdentityHandler encryption.DeleteIdentityHandler
	// EncryptionDeleteKeyHandler sets the operation handler for the delete key operation
	EncryptionDeleteKeyHandler encryption.DeleteKeyHandler
	// EncryptionDeletePolicyHandler sets the operation handler for the delete policy operation
	EncryptionDeletePolicyHandler encryption.DeletePolicyHandler
	// EncryptionDescribeIdentityHandler sets the operation handler for the describe identity operation
	EncryptionDescribeIdentityHandler encryption.DescribeIdentityHandler
	// EncryptionDescribeKeyHandler sets the operation handler for the describe key operation
	EncryptionDescribeKeyHandler encryption.DescribeKeyHandler
	// EncryptionDescribePolicyHandler sets the operation handler for the describe policy operation
	EncryptionDescribePolicyHandler encryption.DescribePolicyHandler
	// EncryptionDescribeSelfIdentityHandler sets the operation handler for the describe self identity operation
	EncryptionDescribeSelfIdentityHandler encryption.DescribeSelfIdentityHandler
	// EncryptionGetPolicyHandler sets the operation handler for the get policy operation
	EncryptionGetPolicyHandler encryption.GetPolicyHandler
	// EncryptionImportKeyHandler sets the operation handler for the import key operation
	EncryptionImportKeyHandler encryption.ImportKeyHandler
	// EncryptionListIdentitiesHandler sets the operation handler for the list identities operation
	EncryptionListIdentitiesHandler encryption.ListIdentitiesHandler
	// EncryptionListKeysHandler sets the operation handler for the list keys operation
	EncryptionListKeysHandler encryption.ListKeysHandler
	// EncryptionListPoliciesHandler sets the operation handler for the list policies operation
	EncryptionListPoliciesHandler encryption.ListPoliciesHandler
	// AuthLoginHandler sets the operation handler for the login operation
	AuthLoginHandler auth.LoginHandler
	// AuthLoginDetailHandler sets the operation handler for the login detail operation
	AuthLoginDetailHandler auth.LoginDetailHandler
	// AuthLogoutHandler sets the operation handler for the logout operation
	AuthLogoutHandler auth.LogoutHandler
	// EncryptionMetricsHandler sets the operation handler for the metrics operation
	EncryptionMetricsHandler encryption.MetricsHandler
	// AuthSessionCheckHandler sets the operation handler for the session check operation
	AuthSessionCheckHandler auth.SessionCheckHandler
	// EncryptionSetPolicyHandler sets the operation handler for the set policy operation
	EncryptionSetPolicyHandler encryption.SetPolicyHandler
	// EncryptionStatusHandler sets the operation handler for the status operation
	EncryptionStatusHandler encryption.StatusHandler
	// EncryptionVersionHandler sets the operation handler for the version operation
	EncryptionVersionHandler encryption.VersionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *KesAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *KesAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *KesAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KesAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KesAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KesAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KesAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KesAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KesAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KesAPI
func (o *KesAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "KeyAuth")
	}

	if o.EncryptionAPIsHandler == nil {
		unregistered = append(unregistered, "encryption.APIsHandler")
	}
	if o.EncryptionAssignPolicyHandler == nil {
		unregistered = append(unregistered, "encryption.AssignPolicyHandler")
	}
	if o.EncryptionCreateKeyHandler == nil {
		unregistered = append(unregistered, "encryption.CreateKeyHandler")
	}
	if o.EncryptionDeleteIdentityHandler == nil {
		unregistered = append(unregistered, "encryption.DeleteIdentityHandler")
	}
	if o.EncryptionDeleteKeyHandler == nil {
		unregistered = append(unregistered, "encryption.DeleteKeyHandler")
	}
	if o.EncryptionDeletePolicyHandler == nil {
		unregistered = append(unregistered, "encryption.DeletePolicyHandler")
	}
	if o.EncryptionDescribeIdentityHandler == nil {
		unregistered = append(unregistered, "encryption.DescribeIdentityHandler")
	}
	if o.EncryptionDescribeKeyHandler == nil {
		unregistered = append(unregistered, "encryption.DescribeKeyHandler")
	}
	if o.EncryptionDescribePolicyHandler == nil {
		unregistered = append(unregistered, "encryption.DescribePolicyHandler")
	}
	if o.EncryptionDescribeSelfIdentityHandler == nil {
		unregistered = append(unregistered, "encryption.DescribeSelfIdentityHandler")
	}
	if o.EncryptionGetPolicyHandler == nil {
		unregistered = append(unregistered, "encryption.GetPolicyHandler")
	}
	if o.EncryptionImportKeyHandler == nil {
		unregistered = append(unregistered, "encryption.ImportKeyHandler")
	}
	if o.EncryptionListIdentitiesHandler == nil {
		unregistered = append(unregistered, "encryption.ListIdentitiesHandler")
	}
	if o.EncryptionListKeysHandler == nil {
		unregistered = append(unregistered, "encryption.ListKeysHandler")
	}
	if o.EncryptionListPoliciesHandler == nil {
		unregistered = append(unregistered, "encryption.ListPoliciesHandler")
	}
	if o.AuthLoginHandler == nil {
		unregistered = append(unregistered, "auth.LoginHandler")
	}
	if o.AuthLoginDetailHandler == nil {
		unregistered = append(unregistered, "auth.LoginDetailHandler")
	}
	if o.AuthLogoutHandler == nil {
		unregistered = append(unregistered, "auth.LogoutHandler")
	}
	if o.EncryptionMetricsHandler == nil {
		unregistered = append(unregistered, "encryption.MetricsHandler")
	}
	if o.AuthSessionCheckHandler == nil {
		unregistered = append(unregistered, "auth.SessionCheckHandler")
	}
	if o.EncryptionSetPolicyHandler == nil {
		unregistered = append(unregistered, "encryption.SetPolicyHandler")
	}
	if o.EncryptionStatusHandler == nil {
		unregistered = append(unregistered, "encryption.StatusHandler")
	}
	if o.EncryptionVersionHandler == nil {
		unregistered = append(unregistered, "encryption.VersionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KesAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KesAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.KeyAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *KesAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *KesAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *KesAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KesAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the kes API
func (o *KesAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KesAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/apis"] = encryption.NewAPIs(o.context, o.EncryptionAPIsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/encryption/policies/{name}/assign"] = encryption.NewAssignPolicy(o.context, o.EncryptionAssignPolicyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/encryption/keys"] = encryption.NewCreateKey(o.context, o.EncryptionCreateKeyHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/encryption/identities/{name}"] = encryption.NewDeleteIdentity(o.context, o.EncryptionDeleteIdentityHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/encryption/keys/{name}"] = encryption.NewDeleteKey(o.context, o.EncryptionDeleteKeyHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/encryption/policies/{name}"] = encryption.NewDeletePolicy(o.context, o.EncryptionDeletePolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/identities/{name}/describe"] = encryption.NewDescribeIdentity(o.context, o.EncryptionDescribeIdentityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/keys/{name}"] = encryption.NewDescribeKey(o.context, o.EncryptionDescribeKeyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/policies/{name}/describe"] = encryption.NewDescribePolicy(o.context, o.EncryptionDescribePolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/describe-self/identity"] = encryption.NewDescribeSelfIdentity(o.context, o.EncryptionDescribeSelfIdentityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/policies/{name}"] = encryption.NewGetPolicy(o.context, o.EncryptionGetPolicyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/encryption/keys/{name}/import"] = encryption.NewImportKey(o.context, o.EncryptionImportKeyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/identities"] = encryption.NewListIdentities(o.context, o.EncryptionListIdentitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/keys"] = encryption.NewListKeys(o.context, o.EncryptionListKeysHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/policies"] = encryption.NewListPolicies(o.context, o.EncryptionListPoliciesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = auth.NewLogin(o.context, o.AuthLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = auth.NewLoginDetail(o.context, o.AuthLoginDetailHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/logout"] = auth.NewLogout(o.context, o.AuthLogoutHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/metrics"] = encryption.NewMetrics(o.context, o.EncryptionMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/session"] = auth.NewSessionCheck(o.context, o.AuthSessionCheckHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/encryption/policies"] = encryption.NewSetPolicy(o.context, o.EncryptionSetPolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/status"] = encryption.NewStatus(o.context, o.EncryptionStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/encryption/version"] = encryption.NewVersion(o.context, o.EncryptionVersionHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KesAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *KesAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *KesAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *KesAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *KesAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}