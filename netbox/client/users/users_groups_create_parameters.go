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

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewUsersGroupsCreateParams creates a new UsersGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsCreateParams() *UsersGroupsCreateParams {
	return &UsersGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsCreateParamsWithTimeout creates a new UsersGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsCreateParamsWithTimeout(timeout time.Duration) *UsersGroupsCreateParams {
	return &UsersGroupsCreateParams{
		timeout: timeout,
	}
}

// NewUsersGroupsCreateParamsWithContext creates a new UsersGroupsCreateParams object
// with the ability to set a context for a request.
func NewUsersGroupsCreateParamsWithContext(ctx context.Context) *UsersGroupsCreateParams {
	return &UsersGroupsCreateParams{
		Context: ctx,
	}
}

// NewUsersGroupsCreateParamsWithHTTPClient creates a new UsersGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsCreateParamsWithHTTPClient(client *http.Client) *UsersGroupsCreateParams {
	return &UsersGroupsCreateParams{
		HTTPClient: client,
	}
}

/* UsersGroupsCreateParams contains all the parameters to send to the API endpoint
   for the users groups create operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsCreateParams struct {

	// Data.
	Data *models.Group

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsCreateParams) WithDefaults() *UsersGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups create params
func (o *UsersGroupsCreateParams) WithTimeout(timeout time.Duration) *UsersGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups create params
func (o *UsersGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups create params
func (o *UsersGroupsCreateParams) WithContext(ctx context.Context) *UsersGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups create params
func (o *UsersGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups create params
func (o *UsersGroupsCreateParams) WithHTTPClient(client *http.Client) *UsersGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups create params
func (o *UsersGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users groups create params
func (o *UsersGroupsCreateParams) WithData(data *models.Group) *UsersGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users groups create params
func (o *UsersGroupsCreateParams) SetData(data *models.Group) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
