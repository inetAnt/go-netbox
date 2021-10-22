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

package circuits

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

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewCircuitsProvidersBulkPartialUpdateParams creates a new CircuitsProvidersBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersBulkPartialUpdateParams() *CircuitsProvidersBulkPartialUpdateParams {
	return &CircuitsProvidersBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersBulkPartialUpdateParamsWithTimeout creates a new CircuitsProvidersBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersBulkPartialUpdateParams {
	return &CircuitsProvidersBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersBulkPartialUpdateParamsWithContext creates a new CircuitsProvidersBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersBulkPartialUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersBulkPartialUpdateParams {
	return &CircuitsProvidersBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersBulkPartialUpdateParamsWithHTTPClient creates a new CircuitsProvidersBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersBulkPartialUpdateParams {
	return &CircuitsProvidersBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits providers bulk partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersBulkPartialUpdateParams struct {

	// Data.
	Data *models.Provider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkPartialUpdateParams) WithDefaults() *CircuitsProvidersBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) WithData(data *models.Provider) *CircuitsProvidersBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers bulk partial update params
func (o *CircuitsProvidersBulkPartialUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
