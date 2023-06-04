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
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// NewDcimFrontPortsBulkPartialUpdateParams creates a new DcimFrontPortsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsBulkPartialUpdateParams() *DcimFrontPortsBulkPartialUpdateParams {
	return &DcimFrontPortsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsBulkPartialUpdateParamsWithTimeout creates a new DcimFrontPortsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsBulkPartialUpdateParams {
	return &DcimFrontPortsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsBulkPartialUpdateParamsWithContext creates a new DcimFrontPortsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsBulkPartialUpdateParams {
	return &DcimFrontPortsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsBulkPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsBulkPartialUpdateParams {
	return &DcimFrontPortsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimFrontPortsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim front ports bulk partial update operation.

	Typically these are written to a http.Request.
*/
type DcimFrontPortsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableFrontPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkPartialUpdateParams) WithDefaults() *DcimFrontPortsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports bulk partial update params
func (o *DcimFrontPortsBulkPartialUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
