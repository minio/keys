// Copyright 2022 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package sys

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"os"

	"github.com/minio/kes"
	"github.com/minio/kes/internal/key"
)

// ErrMoreUnsealKeysRequired is an error indicating that the Unsealer
// requires more UnsealKeys to decrypt the sealed key.
var ErrMoreUnsealKeysRequired = errors.New("sys: more unseal keys required")

// A Sealer seals secrets, i.e. cryptographic keys.
type Sealer interface {
	// Seal encrypts the given plaintext and
	// returns a Stanza and a set of UnsealKeys.
	//
	// The Stanza contains the sealed plaintext
	// data. The set of UnsealKeys are used by a
	// corresponding Unsealer to decrypt the stanza
	// and obtain the original plaintext.
	Seal(plaintext []byte) (*Stanza, []UnsealKey, error)
}

// An Unsealer decrypts sealed secrets encrypted by
// the corresponding Sealer.
type Unsealer interface {
	// Unseal decrypts the given stanza with a set of
	// UnsealKeys and returns the plaintext data.
	//
	// Unseal returns MoreUnsealKeysRequired if at
	// least one more UnsealKey is required to decrypt
	// the Stanza.
	Unseal(*Stanza, ...UnsealKey) ([]byte, error)
}

// An UnsealKey is a key generated by a Sealer that
// can (partially) unseal a Stanza.
type UnsealKey interface {
	String() string
}

const (
	envUnsealerType = "ENVIRONMENT_SECRET"
)

// SealFromEnvironment returns a new Sealer that
// encrypts secrets using a key from the named
// environment variable.
// It returns an error if it fails to read a
// key from the named environment variable.
func SealFromEnvironment(name string) (Sealer, error) {
	b64Key, ok := os.LookupEnv(name)
	if !ok {
		return nil, fmt.Errorf("sys: seal env. variable '%s' not found", name)
	}
	b, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, err
	}
	key, err := key.New(kes.AES256_GCM_SHA256, b, "")
	if err != nil {
		return nil, err
	}
	return &envSealer{
		name: name,
		key:  key,
	}, nil
}

// UnsealFromEnvironment returns a new Unsealer that
// decrypts sealed secrets using a key from the
// environment.
func UnsealFromEnvironment() Unsealer {
	return new(envUnsealer)
}

type envSealer struct {
	name string
	key  key.Key
}

func (s *envSealer) Seal(plaintext []byte) (*Stanza, []UnsealKey, error) {
	type Body struct {
		Name       string
		Ciphertext []byte
	}

	ciphertext, err := s.key.Wrap(plaintext, nil)
	if err != nil {
		return nil, nil, err
	}
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(Body{Name: s.name, Ciphertext: ciphertext}); err != nil {
		return nil, nil, err
	}
	return &Stanza{
		Type: envUnsealerType,
		Body: buffer.Bytes(),
	}, []UnsealKey{}, nil
}

type envUnsealer struct{}

func (envUnsealer) Unseal(stanza *Stanza, keys ...UnsealKey) ([]byte, error) {
	type Body struct {
		Name       string
		Ciphertext []byte
	}

	if stanza.Type != envUnsealerType {
		return nil, fmt.Errorf("sys: incompatible stanza '%s'", stanza.Type)
	}
	if len(keys) != 0 {
	}

	var body Body
	if err := gob.NewDecoder(bytes.NewReader(stanza.Body)).Decode(&body); err != nil {
		return nil, err
	}
	b64Key, ok := os.LookupEnv(body.Name)
	if !ok {
		return nil, fmt.Errorf("sys: unseal env. variable '%s' not found", body.Name)
	}
	b, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, err
	}
	key, err := key.New(kes.AES256_GCM_SHA256, b, "")
	if err != nil {
		return nil, err
	}
	return key.Unwrap(body.Ciphertext, nil)
}
