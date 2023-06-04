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

// NewDcimConsolePortTemplatesCreateParams creates a new DcimConsolePortTemplatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesCreateParams() *DcimConsolePortTemplatesCreateParams {
	return &DcimConsolePortTemplatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesCreateParamsWithTimeout creates a new DcimConsolePortTemplatesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesCreateParams {
	return &DcimConsolePortTemplatesCreateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesCreateParamsWithContext creates a new DcimConsolePortTemplatesCreateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesCreateParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesCreateParams {
	return &DcimConsolePortTemplatesCreateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesCreateParamsWithHTTPClient creates a new DcimConsolePortTemplatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesCreateParams {
	return &DcimConsolePortTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*
DcimConsolePortTemplatesCreateParams contains all the parameters to send to the API endpoint

	for the dcim console port templates create operation.

	Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesCreateParams struct {

	// Data.
	Data *models.WritableConsolePortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesCreateParams) WithDefaults() *DcimConsolePortTemplatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) WithData(data *models.WritableConsolePortTemplate) *DcimConsolePortTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console port templates create params
func (o *DcimConsolePortTemplatesCreateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
