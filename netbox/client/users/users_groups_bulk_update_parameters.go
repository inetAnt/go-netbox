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

package users

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

// NewUsersGroupsBulkUpdateParams creates a new UsersGroupsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsBulkUpdateParams() *UsersGroupsBulkUpdateParams {
	return &UsersGroupsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsBulkUpdateParamsWithTimeout creates a new UsersGroupsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsBulkUpdateParamsWithTimeout(timeout time.Duration) *UsersGroupsBulkUpdateParams {
	return &UsersGroupsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewUsersGroupsBulkUpdateParamsWithContext creates a new UsersGroupsBulkUpdateParams object
// with the ability to set a context for a request.
func NewUsersGroupsBulkUpdateParamsWithContext(ctx context.Context) *UsersGroupsBulkUpdateParams {
	return &UsersGroupsBulkUpdateParams{
		Context: ctx,
	}
}

// NewUsersGroupsBulkUpdateParamsWithHTTPClient creates a new UsersGroupsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsBulkUpdateParamsWithHTTPClient(client *http.Client) *UsersGroupsBulkUpdateParams {
	return &UsersGroupsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
UsersGroupsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the users groups bulk update operation.

	Typically these are written to a http.Request.
*/
type UsersGroupsBulkUpdateParams struct {

	// Data.
	Data *models.Group

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsBulkUpdateParams) WithDefaults() *UsersGroupsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) WithTimeout(timeout time.Duration) *UsersGroupsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) WithContext(ctx context.Context) *UsersGroupsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) WithHTTPClient(client *http.Client) *UsersGroupsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) WithData(data *models.Group) *UsersGroupsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users groups bulk update params
func (o *UsersGroupsBulkUpdateParams) SetData(data *models.Group) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
