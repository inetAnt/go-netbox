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

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewDcimRackReservationsBulkPartialUpdateParams creates a new DcimRackReservationsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsBulkPartialUpdateParams() *DcimRackReservationsBulkPartialUpdateParams {
	return &DcimRackReservationsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsBulkPartialUpdateParamsWithTimeout creates a new DcimRackReservationsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackReservationsBulkPartialUpdateParams {
	return &DcimRackReservationsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsBulkPartialUpdateParamsWithContext creates a new DcimRackReservationsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRackReservationsBulkPartialUpdateParams {
	return &DcimRackReservationsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsBulkPartialUpdateParamsWithHTTPClient creates a new DcimRackReservationsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackReservationsBulkPartialUpdateParams {
	return &DcimRackReservationsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRackReservation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkPartialUpdateParams) WithDefaults() *DcimRackReservationsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackReservationsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRackReservationsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackReservationsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) WithData(data *models.WritableRackReservation) *DcimRackReservationsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack reservations bulk partial update params
func (o *DcimRackReservationsBulkPartialUpdateParams) SetData(data *models.WritableRackReservation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
