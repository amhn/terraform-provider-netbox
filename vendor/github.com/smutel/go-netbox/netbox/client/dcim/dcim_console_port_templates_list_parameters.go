// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimConsolePortTemplatesListParams creates a new DcimConsolePortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesListParams() *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithTimeout creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithContext creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesListParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesListParamsWithHTTPClient creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim console port templates list operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesListParams) WithDefaults() *DcimConsolePortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithCreated(created *string) *DcimConsolePortTemplatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithCreatedGte(createdGte *string) *DcimConsolePortTemplatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithCreatedLte(createdLte *string) *DcimConsolePortTemplatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevicetypeID adds the devicetypeID to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimConsolePortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimConsolePortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithID(id *string) *DcimConsolePortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDGt(iDGt *string) *DcimConsolePortTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDGte(iDGte *string) *DcimConsolePortTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDLt(iDLt *string) *DcimConsolePortTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDLte(iDLte *string) *DcimConsolePortTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDn(iDn *string) *DcimConsolePortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLastUpdated(lastUpdated *string) *DcimConsolePortTemplatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimConsolePortTemplatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimConsolePortTemplatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLimit(limit *int64) *DcimConsolePortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithName(name *string) *DcimConsolePortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameEmpty(nameEmpty *string) *DcimConsolePortTemplatesListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIc(nameIc *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIe(nameIe *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIew(nameIew *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNamen(namen *string) *DcimConsolePortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNic(nameNic *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNie(nameNie *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithOffset(offset *int64) *DcimConsolePortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithType adds the typeVar to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithType(typeVar *string) *DcimConsolePortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithTypen(typen *string) *DcimConsolePortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

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

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
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

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
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

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
