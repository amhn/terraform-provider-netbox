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

package dcim

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

	models "github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimRegionsCreateParams creates a new DcimRegionsCreateParams object
// with the default values initialized.
func NewDcimRegionsCreateParams() *DcimRegionsCreateParams {
	var ()
	return &DcimRegionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsCreateParamsWithTimeout creates a new DcimRegionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRegionsCreateParamsWithTimeout(timeout time.Duration) *DcimRegionsCreateParams {
	var ()
	return &DcimRegionsCreateParams{

		timeout: timeout,
	}
}

// NewDcimRegionsCreateParamsWithContext creates a new DcimRegionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRegionsCreateParamsWithContext(ctx context.Context) *DcimRegionsCreateParams {
	var ()
	return &DcimRegionsCreateParams{

		Context: ctx,
	}
}

// NewDcimRegionsCreateParamsWithHTTPClient creates a new DcimRegionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRegionsCreateParamsWithHTTPClient(client *http.Client) *DcimRegionsCreateParams {
	var ()
	return &DcimRegionsCreateParams{
		HTTPClient: client,
	}
}

/*DcimRegionsCreateParams contains all the parameters to send to the API endpoint
for the dcim regions create operation typically these are written to a http.Request
*/
type DcimRegionsCreateParams struct {

	/*Data*/
	Data *models.WritableRegion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim regions create params
func (o *DcimRegionsCreateParams) WithTimeout(timeout time.Duration) *DcimRegionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions create params
func (o *DcimRegionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions create params
func (o *DcimRegionsCreateParams) WithContext(ctx context.Context) *DcimRegionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions create params
func (o *DcimRegionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions create params
func (o *DcimRegionsCreateParams) WithHTTPClient(client *http.Client) *DcimRegionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions create params
func (o *DcimRegionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim regions create params
func (o *DcimRegionsCreateParams) WithData(data *models.WritableRegion) *DcimRegionsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim regions create params
func (o *DcimRegionsCreateParams) SetData(data *models.WritableRegion) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
