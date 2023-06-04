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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// VirtualizationClustersBulkUpdateReader is a Reader for the VirtualizationClustersBulkUpdate structure.
type VirtualizationClustersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersBulkUpdateOK creates a VirtualizationClustersBulkUpdateOK with default headers values
func NewVirtualizationClustersBulkUpdateOK() *VirtualizationClustersBulkUpdateOK {
	return &VirtualizationClustersBulkUpdateOK{}
}

/*
VirtualizationClustersBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersBulkUpdateOK virtualization clusters bulk update o k
*/
type VirtualizationClustersBulkUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters bulk update o k response has a 2xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters bulk update o k response has a 3xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters bulk update o k response has a 4xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters bulk update o k response has a 5xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters bulk update o k response a status code equal to that given
func (o *VirtualizationClustersBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization clusters bulk update o k response
func (o *VirtualizationClustersBulkUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClustersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersBulkUpdateDefault creates a VirtualizationClustersBulkUpdateDefault with default headers values
func NewVirtualizationClustersBulkUpdateDefault(code int) *VirtualizationClustersBulkUpdateDefault {
	return &VirtualizationClustersBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersBulkUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClustersBulkUpdateDefault virtualization clusters bulk update default
*/
type VirtualizationClustersBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization clusters bulk update default response has a 2xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters bulk update default response has a 3xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters bulk update default response has a 4xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters bulk update default response has a 5xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters bulk update default response a status code equal to that given
func (o *VirtualizationClustersBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization clusters bulk update default response
func (o *VirtualizationClustersBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClustersBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualization_clusters_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualization_clusters_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
