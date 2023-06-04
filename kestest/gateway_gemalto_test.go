// Copyright 2023 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package kestest_test

import (
	"flag"
	"os"
	"testing"

	"github.com/minio/kes/edge/edgeconf"
)

var gemaltoConfigFile = flag.String("gemalto.config", "", "Path to a KES config file with Gemalto SecretsManager config")

func TestGatewayGemalto(t *testing.T) {
	if *gemaltoConfigFile == "" {
		t.Skip("Gemalto tests disabled. Use -gemalto.config=<config file with Gemalto SecretManager config> to enable them")
	}
	file, err := os.Open(*gemaltoConfigFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	ctx, cancel := testingContext(t)
	defer cancel()

	store, _, err := edgeconf.Connect(ctx, file)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Metrics", func(t *testing.T) { testMetrics(ctx, store, t) })
	t.Run("APIs", func(t *testing.T) { testAPIs(ctx, store, t) })
	t.Run("CreateKey", func(t *testing.T) { testCreateKey(ctx, store, t, RandString(ranStringLength)) })
	t.Run("ImportKey", func(t *testing.T) { testImportKey(ctx, store, t, RandString(ranStringLength)) })
	t.Run("GenerateKey", func(t *testing.T) { testGenerateKey(ctx, store, t, RandString(ranStringLength)) })
	t.Run("EncryptKey", func(t *testing.T) { testEncryptKey(ctx, store, t, RandString(ranStringLength)) })
	t.Run("DecryptKey", func(t *testing.T) { testDecryptKey(ctx, store, t, RandString(ranStringLength)) })
	t.Run("DescribePolicy", func(t *testing.T) { testDescribePolicy(ctx, store, t) })
	t.Run("GetPolicy", func(t *testing.T) { testGetPolicy(ctx, store, t) })
	t.Run("SelfDescribe", func(t *testing.T) { testSelfDescribe(ctx, store, t) })
}
