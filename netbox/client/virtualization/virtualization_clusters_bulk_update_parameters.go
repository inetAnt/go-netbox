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

package virtualization

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

// NewVirtualizationClustersBulkUpdateParams creates a new VirtualizationClustersBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersBulkUpdateParams() *VirtualizationClustersBulkUpdateParams {
	return &VirtualizationClustersBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersBulkUpdateParamsWithTimeout creates a new VirtualizationClustersBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersBulkUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersBulkUpdateParams {
	return &VirtualizationClustersBulkUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersBulkUpdateParamsWithContext creates a new VirtualizationClustersBulkUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersBulkUpdateParamsWithContext(ctx context.Context) *VirtualizationClustersBulkUpdateParams {
	return &VirtualizationClustersBulkUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersBulkUpdateParamsWithHTTPClient creates a new VirtualizationClustersBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersBulkUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersBulkUpdateParams {
	return &VirtualizationClustersBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
VirtualizationClustersBulkUpdateParams contains all the parameters to send to the API endpoint

	for the virtualization clusters bulk update operation.

	Typically these are written to a http.Request.
*/
type VirtualizationClustersBulkUpdateParams struct {

	// Data.
	Data *models.WritableCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersBulkUpdateParams) WithDefaults() *VirtualizationClustersBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) WithContext(ctx context.Context) *VirtualizationClustersBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) WithData(data *models.WritableCluster) *VirtualizationClustersBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters bulk update params
func (o *VirtualizationClustersBulkUpdateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
