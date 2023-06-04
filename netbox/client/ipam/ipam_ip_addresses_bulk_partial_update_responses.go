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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// IpamIPAddressesBulkPartialUpdateReader is a Reader for the IpamIPAddressesBulkPartialUpdate structure.
type IpamIPAddressesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesBulkPartialUpdateOK creates a IpamIPAddressesBulkPartialUpdateOK with default headers values
func NewIpamIPAddressesBulkPartialUpdateOK() *IpamIPAddressesBulkPartialUpdateOK {
	return &IpamIPAddressesBulkPartialUpdateOK{}
}

/*
IpamIPAddressesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkPartialUpdateOK ipam Ip addresses bulk partial update o k
*/
type IpamIPAddressesBulkPartialUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses bulk partial update o k response has a 2xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses bulk partial update o k response has a 3xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses bulk partial update o k response has a 4xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses bulk partial update o k response has a 5xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses bulk partial update o k response a status code equal to that given
func (o *IpamIPAddressesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip addresses bulk partial update o k response
func (o *IpamIPAddressesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamIPAddressesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipamIpAddressesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipamIpAddressesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesBulkPartialUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesBulkPartialUpdateDefault creates a IpamIPAddressesBulkPartialUpdateDefault with default headers values
func NewIpamIPAddressesBulkPartialUpdateDefault(code int) *IpamIPAddressesBulkPartialUpdateDefault {
	return &IpamIPAddressesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamIPAddressesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamIPAddressesBulkPartialUpdateDefault ipam ip addresses bulk partial update default
*/
type IpamIPAddressesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam ip addresses bulk partial update default response has a 2xx status code
func (o *IpamIPAddressesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip addresses bulk partial update default response has a 3xx status code
func (o *IpamIPAddressesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip addresses bulk partial update default response has a 4xx status code
func (o *IpamIPAddressesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip addresses bulk partial update default response has a 5xx status code
func (o *IpamIPAddressesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip addresses bulk partial update default response a status code equal to that given
func (o *IpamIPAddressesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam ip addresses bulk partial update default response
func (o *IpamIPAddressesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPAddressesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipam_ip-addresses_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipam_ip-addresses_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
