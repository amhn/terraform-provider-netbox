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

package ipam

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

// NewIpamAggregatesListParams creates a new IpamAggregatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesListParams() *IpamAggregatesListParams {
	return &IpamAggregatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesListParamsWithTimeout creates a new IpamAggregatesListParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesListParamsWithTimeout(timeout time.Duration) *IpamAggregatesListParams {
	return &IpamAggregatesListParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesListParamsWithContext creates a new IpamAggregatesListParams object
// with the ability to set a context for a request.
func NewIpamAggregatesListParamsWithContext(ctx context.Context) *IpamAggregatesListParams {
	return &IpamAggregatesListParams{
		Context: ctx,
	}
}

// NewIpamAggregatesListParamsWithHTTPClient creates a new IpamAggregatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesListParamsWithHTTPClient(client *http.Client) *IpamAggregatesListParams {
	return &IpamAggregatesListParams{
		HTTPClient: client,
	}
}

/* IpamAggregatesListParams contains all the parameters to send to the API endpoint
   for the ipam aggregates list operation.

   Typically these are written to a http.Request.
*/
type IpamAggregatesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DateAdded.
	DateAdded *string

	// DateAddedGt.
	DateAddedGt *string

	// DateAddedGte.
	DateAddedGte *string

	// DateAddedLt.
	DateAddedLt *string

	// DateAddedLte.
	DateAddedLte *string

	// DateAddedn.
	DateAddedn *string

	// Description.
	Description *string

	// DescriptionEmpty.
	DescriptionEmpty *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// Family.
	Family *float64

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Prefix.
	Prefix *string

	// Q.
	Q *string

	// Rir.
	Rir *string

	// Rirn.
	Rirn *string

	// RirID.
	RirID *string

	// RirIDn.
	RirIDn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// Tenantn.
	Tenantn *string

	// TenantGroup.
	TenantGroup *string

	// TenantGroupn.
	TenantGroupn *string

	// TenantGroupID.
	TenantGroupID *string

	// TenantGroupIDn.
	TenantGroupIDn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesListParams) WithDefaults() *IpamAggregatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTimeout(timeout time.Duration) *IpamAggregatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithContext(ctx context.Context) *IpamAggregatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithHTTPClient(client *http.Client) *IpamAggregatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithCreated(created *string) *IpamAggregatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithCreatedGte(createdGte *string) *IpamAggregatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithCreatedLte(createdLte *string) *IpamAggregatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDateAdded adds the dateAdded to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAdded(dateAdded *string) *IpamAggregatesListParams {
	o.SetDateAdded(dateAdded)
	return o
}

// SetDateAdded adds the dateAdded to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAdded(dateAdded *string) {
	o.DateAdded = dateAdded
}

// WithDateAddedGt adds the dateAddedGt to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAddedGt(dateAddedGt *string) *IpamAggregatesListParams {
	o.SetDateAddedGt(dateAddedGt)
	return o
}

// SetDateAddedGt adds the dateAddedGt to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAddedGt(dateAddedGt *string) {
	o.DateAddedGt = dateAddedGt
}

// WithDateAddedGte adds the dateAddedGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAddedGte(dateAddedGte *string) *IpamAggregatesListParams {
	o.SetDateAddedGte(dateAddedGte)
	return o
}

// SetDateAddedGte adds the dateAddedGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAddedGte(dateAddedGte *string) {
	o.DateAddedGte = dateAddedGte
}

// WithDateAddedLt adds the dateAddedLt to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAddedLt(dateAddedLt *string) *IpamAggregatesListParams {
	o.SetDateAddedLt(dateAddedLt)
	return o
}

// SetDateAddedLt adds the dateAddedLt to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAddedLt(dateAddedLt *string) {
	o.DateAddedLt = dateAddedLt
}

// WithDateAddedLte adds the dateAddedLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAddedLte(dateAddedLte *string) *IpamAggregatesListParams {
	o.SetDateAddedLte(dateAddedLte)
	return o
}

// SetDateAddedLte adds the dateAddedLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAddedLte(dateAddedLte *string) {
	o.DateAddedLte = dateAddedLte
}

// WithDateAddedn adds the dateAddedn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAddedn(dateAddedn *string) *IpamAggregatesListParams {
	o.SetDateAddedn(dateAddedn)
	return o
}

// SetDateAddedn adds the dateAddedN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAddedn(dateAddedn *string) {
	o.DateAddedn = dateAddedn
}

// WithDescription adds the description to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescription(description *string) *IpamAggregatesListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionEmpty adds the descriptionEmpty to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionEmpty(descriptionEmpty *string) *IpamAggregatesListParams {
	o.SetDescriptionEmpty(descriptionEmpty)
	return o
}

// SetDescriptionEmpty adds the descriptionEmpty to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionEmpty(descriptionEmpty *string) {
	o.DescriptionEmpty = descriptionEmpty
}

// WithDescriptionIc adds the descriptionIc to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionIc(descriptionIc *string) *IpamAggregatesListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionIe(descriptionIe *string) *IpamAggregatesListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionIew(descriptionIew *string) *IpamAggregatesListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionIsw(descriptionIsw *string) *IpamAggregatesListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionn(descriptionn *string) *IpamAggregatesListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionNic(descriptionNic *string) *IpamAggregatesListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionNie(descriptionNie *string) *IpamAggregatesListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionNiew(descriptionNiew *string) *IpamAggregatesListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDescriptionNisw(descriptionNisw *string) *IpamAggregatesListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithFamily adds the family to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithFamily(family *float64) *IpamAggregatesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetFamily(family *float64) {
	o.Family = family
}

// WithID adds the id to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithID(id *string) *IpamAggregatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDGt(iDGt *string) *IpamAggregatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDGte(iDGte *string) *IpamAggregatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDLt(iDLt *string) *IpamAggregatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDLte(iDLte *string) *IpamAggregatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDn(iDn *string) *IpamAggregatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithLastUpdated(lastUpdated *string) *IpamAggregatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamAggregatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamAggregatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithLimit(limit *int64) *IpamAggregatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithOffset(offset *int64) *IpamAggregatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPrefix adds the prefix to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithPrefix(prefix *string) *IpamAggregatesListParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WithQ adds the q to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithQ(q *string) *IpamAggregatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRir adds the rir to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRir(rir *string) *IpamAggregatesListParams {
	o.SetRir(rir)
	return o
}

// SetRir adds the rir to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRir(rir *string) {
	o.Rir = rir
}

// WithRirn adds the rirn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRirn(rirn *string) *IpamAggregatesListParams {
	o.SetRirn(rirn)
	return o
}

// SetRirn adds the rirN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRirn(rirn *string) {
	o.Rirn = rirn
}

// WithRirID adds the rirID to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRirID(rirID *string) *IpamAggregatesListParams {
	o.SetRirID(rirID)
	return o
}

// SetRirID adds the rirId to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRirID(rirID *string) {
	o.RirID = rirID
}

// WithRirIDn adds the rirIDn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRirIDn(rirIDn *string) *IpamAggregatesListParams {
	o.SetRirIDn(rirIDn)
	return o
}

// SetRirIDn adds the rirIdN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRirIDn(rirIDn *string) {
	o.RirIDn = rirIDn
}

// WithTag adds the tag to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTag(tag *string) *IpamAggregatesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTagn(tagn *string) *IpamAggregatesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenant(tenant *string) *IpamAggregatesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantn(tenantn *string) *IpamAggregatesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantGroup(tenantGroup *string) *IpamAggregatesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantGroupn(tenantGroupn *string) *IpamAggregatesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantGroupID(tenantGroupID *string) *IpamAggregatesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamAggregatesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantID(tenantID *string) *IpamAggregatesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTenantIDn(tenantIDn *string) *IpamAggregatesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DateAdded != nil {

		// query param date_added
		var qrDateAdded string

		if o.DateAdded != nil {
			qrDateAdded = *o.DateAdded
		}
		qDateAdded := qrDateAdded
		if qDateAdded != "" {

			if err := r.SetQueryParam("date_added", qDateAdded); err != nil {
				return err
			}
		}
	}

	if o.DateAddedGt != nil {

		// query param date_added__gt
		var qrDateAddedGt string

		if o.DateAddedGt != nil {
			qrDateAddedGt = *o.DateAddedGt
		}
		qDateAddedGt := qrDateAddedGt
		if qDateAddedGt != "" {

			if err := r.SetQueryParam("date_added__gt", qDateAddedGt); err != nil {
				return err
			}
		}
	}

	if o.DateAddedGte != nil {

		// query param date_added__gte
		var qrDateAddedGte string

		if o.DateAddedGte != nil {
			qrDateAddedGte = *o.DateAddedGte
		}
		qDateAddedGte := qrDateAddedGte
		if qDateAddedGte != "" {

			if err := r.SetQueryParam("date_added__gte", qDateAddedGte); err != nil {
				return err
			}
		}
	}

	if o.DateAddedLt != nil {

		// query param date_added__lt
		var qrDateAddedLt string

		if o.DateAddedLt != nil {
			qrDateAddedLt = *o.DateAddedLt
		}
		qDateAddedLt := qrDateAddedLt
		if qDateAddedLt != "" {

			if err := r.SetQueryParam("date_added__lt", qDateAddedLt); err != nil {
				return err
			}
		}
	}

	if o.DateAddedLte != nil {

		// query param date_added__lte
		var qrDateAddedLte string

		if o.DateAddedLte != nil {
			qrDateAddedLte = *o.DateAddedLte
		}
		qDateAddedLte := qrDateAddedLte
		if qDateAddedLte != "" {

			if err := r.SetQueryParam("date_added__lte", qDateAddedLte); err != nil {
				return err
			}
		}
	}

	if o.DateAddedn != nil {

		// query param date_added__n
		var qrDateAddedn string

		if o.DateAddedn != nil {
			qrDateAddedn = *o.DateAddedn
		}
		qDateAddedn := qrDateAddedn
		if qDateAddedn != "" {

			if err := r.SetQueryParam("date_added__n", qDateAddedn); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionEmpty != nil {

		// query param description__empty
		var qrDescriptionEmpty string

		if o.DescriptionEmpty != nil {
			qrDescriptionEmpty = *o.DescriptionEmpty
		}
		qDescriptionEmpty := qrDescriptionEmpty
		if qDescriptionEmpty != "" {

			if err := r.SetQueryParam("description__empty", qDescriptionEmpty); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.Family != nil {

		// query param family
		var qrFamily float64

		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := swag.FormatFloat64(qrFamily)
		if qFamily != "" {

			if err := r.SetQueryParam("family", qFamily); err != nil {
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

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string

		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {

			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
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

	if o.Rir != nil {

		// query param rir
		var qrRir string

		if o.Rir != nil {
			qrRir = *o.Rir
		}
		qRir := qrRir
		if qRir != "" {

			if err := r.SetQueryParam("rir", qRir); err != nil {
				return err
			}
		}
	}

	if o.Rirn != nil {

		// query param rir__n
		var qrRirn string

		if o.Rirn != nil {
			qrRirn = *o.Rirn
		}
		qRirn := qrRirn
		if qRirn != "" {

			if err := r.SetQueryParam("rir__n", qRirn); err != nil {
				return err
			}
		}
	}

	if o.RirID != nil {

		// query param rir_id
		var qrRirID string

		if o.RirID != nil {
			qrRirID = *o.RirID
		}
		qRirID := qrRirID
		if qRirID != "" {

			if err := r.SetQueryParam("rir_id", qRirID); err != nil {
				return err
			}
		}
	}

	if o.RirIDn != nil {

		// query param rir_id__n
		var qrRirIDn string

		if o.RirIDn != nil {
			qrRirIDn = *o.RirIDn
		}
		qRirIDn := qrRirIDn
		if qRirIDn != "" {

			if err := r.SetQueryParam("rir_id__n", qRirIDn); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string

		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {

			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string

		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {

			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string

		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {

			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string

		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {

			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
