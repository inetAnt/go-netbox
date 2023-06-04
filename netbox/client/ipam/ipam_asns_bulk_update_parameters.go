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

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// NewIpamAsnsBulkUpdateParams creates a new IpamAsnsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAsnsBulkUpdateParams() *IpamAsnsBulkUpdateParams {
	return &IpamAsnsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAsnsBulkUpdateParamsWithTimeout creates a new IpamAsnsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAsnsBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamAsnsBulkUpdateParams {
	return &IpamAsnsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAsnsBulkUpdateParamsWithContext creates a new IpamAsnsBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamAsnsBulkUpdateParamsWithContext(ctx context.Context) *IpamAsnsBulkUpdateParams {
	return &IpamAsnsBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamAsnsBulkUpdateParamsWithHTTPClient creates a new IpamAsnsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAsnsBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamAsnsBulkUpdateParams {
	return &IpamAsnsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamAsnsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam asns bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamAsnsBulkUpdateParams struct {

	// Data.
	Data *models.WritableASN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam asns bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsBulkUpdateParams) WithDefaults() *IpamAsnsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam asns bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamAsnsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) WithContext(ctx context.Context) *IpamAsnsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamAsnsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) WithData(data *models.WritableASN) *IpamAsnsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam asns bulk update params
func (o *IpamAsnsBulkUpdateParams) SetData(data *models.WritableASN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAsnsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
