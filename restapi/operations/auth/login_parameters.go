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

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// LoginMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var LoginMaxParseMemory int64 = 32 << 20

// NewLoginParams creates a new LoginParams object
//
// There are no default values defined in the spec.
func NewLoginParams() LoginParams {

	return LoginParams{}
}

// LoginParams contains all the bound params for the login operation
// typically these are obtained from a http.Request
//
// swagger:parameters Login
type LoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: formData
	*/
	APIKey *string
	/*
	  Required: true
	  In: formData
	*/
	Cert io.ReadCloser
	/*
	  In: formData
	*/
	Insecure *string
	/*
	  Required: true
	  In: formData
	*/
	Key io.ReadCloser
	/*
	  In: formData
	*/
	Password *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLoginParams() beforehand.
func (o *LoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(LoginMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdAPIKey, fdhkAPIKey, _ := fds.GetOK("apiKey")
	if err := o.bindAPIKey(fdAPIKey, fdhkAPIKey, route.Formats); err != nil {
		res = append(res, err)
	}

	cert, certHeader, err := r.FormFile("cert")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "cert", err))
	} else if err := o.bindCert(cert, certHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Cert = &runtime.File{Data: cert, Header: certHeader}
	}

	fdInsecure, fdhkInsecure, _ := fds.GetOK("insecure")
	if err := o.bindInsecure(fdInsecure, fdhkInsecure, route.Formats); err != nil {
		res = append(res, err)
	}

	key, keyHeader, err := r.FormFile("key")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "key", err))
	} else if err := o.bindKey(key, keyHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Key = &runtime.File{Data: key, Header: keyHeader}
	}

	fdPassword, fdhkPassword, _ := fds.GetOK("password")
	if err := o.bindPassword(fdPassword, fdhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAPIKey binds and validates parameter APIKey from formData.
func (o *LoginParams) bindAPIKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.APIKey = &raw

	return nil
}

// bindCert binds file parameter Cert.
//
// The only supported validations on files are MinLength and MaxLength
func (o *LoginParams) bindCert(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindInsecure binds and validates parameter Insecure from formData.
func (o *LoginParams) bindInsecure(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Insecure = &raw

	return nil
}

// bindKey binds file parameter Key.
//
// The only supported validations on files are MinLength and MaxLength
func (o *LoginParams) bindKey(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindPassword binds and validates parameter Password from formData.
func (o *LoginParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Password = &raw

	return nil
}