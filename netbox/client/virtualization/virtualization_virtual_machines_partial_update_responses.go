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

// VirtualizationVirtualMachinesPartialUpdateReader is a Reader for the VirtualizationVirtualMachinesPartialUpdate structure.
type VirtualizationVirtualMachinesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesPartialUpdateOK creates a VirtualizationVirtualMachinesPartialUpdateOK with default headers values
func NewVirtualizationVirtualMachinesPartialUpdateOK() *VirtualizationVirtualMachinesPartialUpdateOK {
	return &VirtualizationVirtualMachinesPartialUpdateOK{}
}

/*
VirtualizationVirtualMachinesPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesPartialUpdateOK virtualization virtual machines partial update o k
*/
type VirtualizationVirtualMachinesPartialUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines partial update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines partial update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines partial update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines partial update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines partial update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual machines partial update o k response
func (o *VirtualizationVirtualMachinesPartialUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesPartialUpdateDefault creates a VirtualizationVirtualMachinesPartialUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesPartialUpdateDefault(code int) *VirtualizationVirtualMachinesPartialUpdateDefault {
	return &VirtualizationVirtualMachinesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualMachinesPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesPartialUpdateDefault virtualization virtual machines partial update default
*/
type VirtualizationVirtualMachinesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization virtual machines partial update default response has a 2xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual machines partial update default response has a 3xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual machines partial update default response has a 4xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual machines partial update default response has a 5xx status code
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual machines partial update default response a status code equal to that given
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization virtual machines partial update default response
func (o *VirtualizationVirtualMachinesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
