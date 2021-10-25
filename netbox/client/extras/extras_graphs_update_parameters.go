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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/inetAnt/go-netbox/netbox/models"
)

// NewExtrasGraphsUpdateParams creates a new ExtrasGraphsUpdateParams object
// with the default values initialized.
func NewExtrasGraphsUpdateParams() *ExtrasGraphsUpdateParams {
	var ()
	return &ExtrasGraphsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsUpdateParamsWithTimeout creates a new ExtrasGraphsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphsUpdateParams {
	var ()
	return &ExtrasGraphsUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsUpdateParamsWithContext creates a new ExtrasGraphsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsUpdateParamsWithContext(ctx context.Context) *ExtrasGraphsUpdateParams {
	var ()
	return &ExtrasGraphsUpdateParams{

		Context: ctx,
	}
}

// NewExtrasGraphsUpdateParamsWithHTTPClient creates a new ExtrasGraphsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphsUpdateParams {
	var ()
	return &ExtrasGraphsUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsUpdateParams contains all the parameters to send to the API endpoint
for the extras graphs update operation typically these are written to a http.Request
*/
type ExtrasGraphsUpdateParams struct {

	/*Data*/
	Data *models.Graph
	/*ID
	  A unique integer value identifying this graph.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) WithContext(ctx context.Context) *ExtrasGraphsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) WithData(data *models.Graph) *ExtrasGraphsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) SetData(data *models.Graph) {
	o.Data = data
}

// WithID adds the id to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) WithID(id int64) *ExtrasGraphsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphs update params
func (o *ExtrasGraphsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
