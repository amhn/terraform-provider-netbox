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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimRearPortTemplatesListParams creates a new DcimRearPortTemplatesListParams object
// with the default values initialized.
func NewDcimRearPortTemplatesListParams() *DcimRearPortTemplatesListParams {
	var ()
	return &DcimRearPortTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesListParamsWithTimeout creates a new DcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRearPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesListParams {
	var ()
	return &DcimRearPortTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesListParamsWithContext creates a new DcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRearPortTemplatesListParamsWithContext(ctx context.Context) *DcimRearPortTemplatesListParams {
	var ()
	return &DcimRearPortTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimRearPortTemplatesListParamsWithHTTPClient creates a new DcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRearPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesListParams {
	var ()
	return &DcimRearPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimRearPortTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim rear port templates list operation typically these are written to a http.Request
*/
type DcimRearPortTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
	/*ID*/
	ID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Positions*/
	Positions *string
	/*Q*/
	Q *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithContext(ctx context.Context) *DcimRearPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimRearPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithID adds the id to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithID(id *string) *DcimRearPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithLimit(limit *int64) *DcimRearPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithName(name *string) *DcimRearPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithOffset(offset *int64) *DcimRearPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPositions adds the positions to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithPositions(positions *string) *DcimRearPortTemplatesListParams {
	o.SetPositions(positions)
	return o
}

// SetPositions adds the positions to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetPositions(positions *string) {
	o.Positions = positions
}

// WithQ adds the q to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithQ(q *string) *DcimRearPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) WithType(typeVar *string) *DcimRearPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim rear port templates list params
func (o *DcimRearPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Positions != nil {

		// query param positions
		var qrPositions string
		if o.Positions != nil {
			qrPositions = *o.Positions
		}
		qPositions := qrPositions
		if qPositions != "" {
			if err := r.SetQueryParam("positions", qPositions); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
