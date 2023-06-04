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

// IpamAsnsPartialUpdateReader is a Reader for the IpamAsnsPartialUpdate structure.
type IpamAsnsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAsnsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAsnsPartialUpdateOK creates a IpamAsnsPartialUpdateOK with default headers values
func NewIpamAsnsPartialUpdateOK() *IpamAsnsPartialUpdateOK {
	return &IpamAsnsPartialUpdateOK{}
}

/*
IpamAsnsPartialUpdateOK describes a response with status code 200, with default header values.

IpamAsnsPartialUpdateOK ipam asns partial update o k
*/
type IpamAsnsPartialUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns partial update o k response has a 2xx status code
func (o *IpamAsnsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns partial update o k response has a 3xx status code
func (o *IpamAsnsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns partial update o k response has a 4xx status code
func (o *IpamAsnsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns partial update o k response has a 5xx status code
func (o *IpamAsnsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns partial update o k response a status code equal to that given
func (o *IpamAsnsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam asns partial update o k response
func (o *IpamAsnsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamAsnsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipamAsnsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipamAsnsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsPartialUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAsnsPartialUpdateDefault creates a IpamAsnsPartialUpdateDefault with default headers values
func NewIpamAsnsPartialUpdateDefault(code int) *IpamAsnsPartialUpdateDefault {
	return &IpamAsnsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAsnsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamAsnsPartialUpdateDefault ipam asns partial update default
*/
type IpamAsnsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam asns partial update default response has a 2xx status code
func (o *IpamAsnsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam asns partial update default response has a 3xx status code
func (o *IpamAsnsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam asns partial update default response has a 4xx status code
func (o *IpamAsnsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam asns partial update default response has a 5xx status code
func (o *IpamAsnsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam asns partial update default response a status code equal to that given
func (o *IpamAsnsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam asns partial update default response
func (o *IpamAsnsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamAsnsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipam_asns_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipam_asns_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAsnsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
