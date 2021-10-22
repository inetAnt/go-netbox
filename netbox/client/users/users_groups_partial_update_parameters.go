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
	"github.com/go-openapi/swag"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewUsersGroupsPartialUpdateParams creates a new UsersGroupsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsPartialUpdateParams() *UsersGroupsPartialUpdateParams {
	return &UsersGroupsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsPartialUpdateParamsWithTimeout creates a new UsersGroupsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersGroupsPartialUpdateParams {
	return &UsersGroupsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewUsersGroupsPartialUpdateParamsWithContext creates a new UsersGroupsPartialUpdateParams object
// with the ability to set a context for a request.
func NewUsersGroupsPartialUpdateParamsWithContext(ctx context.Context) *UsersGroupsPartialUpdateParams {
	return &UsersGroupsPartialUpdateParams{
		Context: ctx,
	}
}

// NewUsersGroupsPartialUpdateParamsWithHTTPClient creates a new UsersGroupsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersGroupsPartialUpdateParams {
	return &UsersGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* UsersGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the users groups partial update operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsPartialUpdateParams struct {

	// Data.
	Data *models.Group

	/* ID.

	   A unique integer value identifying this group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsPartialUpdateParams) WithDefaults() *UsersGroupsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) WithContext(ctx context.Context) *UsersGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) WithData(data *models.Group) *UsersGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) SetData(data *models.Group) {
	o.Data = data
}

// WithID adds the id to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) WithID(id int64) *UsersGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users groups partial update params
func (o *UsersGroupsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
