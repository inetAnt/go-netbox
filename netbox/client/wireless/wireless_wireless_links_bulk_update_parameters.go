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

package wireless

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

// NewWirelessWirelessLinksBulkUpdateParams creates a new WirelessWirelessLinksBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLinksBulkUpdateParams() *WirelessWirelessLinksBulkUpdateParams {
	return &WirelessWirelessLinksBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLinksBulkUpdateParamsWithTimeout creates a new WirelessWirelessLinksBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLinksBulkUpdateParamsWithTimeout(timeout time.Duration) *WirelessWirelessLinksBulkUpdateParams {
	return &WirelessWirelessLinksBulkUpdateParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLinksBulkUpdateParamsWithContext creates a new WirelessWirelessLinksBulkUpdateParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLinksBulkUpdateParamsWithContext(ctx context.Context) *WirelessWirelessLinksBulkUpdateParams {
	return &WirelessWirelessLinksBulkUpdateParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLinksBulkUpdateParamsWithHTTPClient creates a new WirelessWirelessLinksBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLinksBulkUpdateParamsWithHTTPClient(client *http.Client) *WirelessWirelessLinksBulkUpdateParams {
	return &WirelessWirelessLinksBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLinksBulkUpdateParams contains all the parameters to send to the API endpoint

	for the wireless wireless links bulk update operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLinksBulkUpdateParams struct {

	// Data.
	Data *models.WritableWirelessLink

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless links bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLinksBulkUpdateParams) WithDefaults() *WirelessWirelessLinksBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless links bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLinksBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) WithTimeout(timeout time.Duration) *WirelessWirelessLinksBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) WithContext(ctx context.Context) *WirelessWirelessLinksBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) WithHTTPClient(client *http.Client) *WirelessWirelessLinksBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) WithData(data *models.WritableWirelessLink) *WirelessWirelessLinksBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the wireless wireless links bulk update params
func (o *WirelessWirelessLinksBulkUpdateParams) SetData(data *models.WritableWirelessLink) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLinksBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
