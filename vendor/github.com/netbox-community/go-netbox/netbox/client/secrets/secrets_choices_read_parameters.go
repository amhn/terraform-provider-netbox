// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSecretsChoicesReadParams creates a new SecretsChoicesReadParams object
// with the default values initialized.
func NewSecretsChoicesReadParams() *SecretsChoicesReadParams {
	var ()
	return &SecretsChoicesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsChoicesReadParamsWithTimeout creates a new SecretsChoicesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsChoicesReadParamsWithTimeout(timeout time.Duration) *SecretsChoicesReadParams {
	var ()
	return &SecretsChoicesReadParams{

		timeout: timeout,
	}
}

// NewSecretsChoicesReadParamsWithContext creates a new SecretsChoicesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsChoicesReadParamsWithContext(ctx context.Context) *SecretsChoicesReadParams {
	var ()
	return &SecretsChoicesReadParams{

		Context: ctx,
	}
}

// NewSecretsChoicesReadParamsWithHTTPClient creates a new SecretsChoicesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsChoicesReadParamsWithHTTPClient(client *http.Client) *SecretsChoicesReadParams {
	var ()
	return &SecretsChoicesReadParams{
		HTTPClient: client,
	}
}

/*SecretsChoicesReadParams contains all the parameters to send to the API endpoint
for the secrets choices read operation typically these are written to a http.Request
*/
type SecretsChoicesReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets choices read params
func (o *SecretsChoicesReadParams) WithTimeout(timeout time.Duration) *SecretsChoicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets choices read params
func (o *SecretsChoicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets choices read params
func (o *SecretsChoicesReadParams) WithContext(ctx context.Context) *SecretsChoicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets choices read params
func (o *SecretsChoicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets choices read params
func (o *SecretsChoicesReadParams) WithHTTPClient(client *http.Client) *SecretsChoicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets choices read params
func (o *SecretsChoicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the secrets choices read params
func (o *SecretsChoicesReadParams) WithID(id string) *SecretsChoicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets choices read params
func (o *SecretsChoicesReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsChoicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
