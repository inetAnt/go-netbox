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

// IpamPrefixesAvailableIpsListReader is a Reader for the IpamPrefixesAvailableIpsList structure.
type IpamPrefixesAvailableIpsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailableIpsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesAvailableIpsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesAvailableIpsListOK creates a IpamPrefixesAvailableIpsListOK with default headers values
func NewIpamPrefixesAvailableIpsListOK() *IpamPrefixesAvailableIpsListOK {
	return &IpamPrefixesAvailableIpsListOK{}
}

/*
IpamPrefixesAvailableIpsListOK describes a response with status code 200, with default header values.

IpamPrefixesAvailableIpsListOK ipam prefixes available ips list o k
*/
type IpamPrefixesAvailableIpsListOK struct {
	Payload []*models.AvailableIP
}

// IsSuccess returns true when this ipam prefixes available ips list o k response has a 2xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes available ips list o k response has a 3xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes available ips list o k response has a 4xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes available ips list o k response has a 5xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes available ips list o k response a status code equal to that given
func (o *IpamPrefixesAvailableIpsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam prefixes available ips list o k response
func (o *IpamPrefixesAvailableIpsListOK) Code() int {
	return 200
}

func (o *IpamPrefixesAvailableIpsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsListOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesAvailableIpsListOK) String() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsListOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesAvailableIpsListOK) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesAvailableIpsListDefault creates a IpamPrefixesAvailableIpsListDefault with default headers values
func NewIpamPrefixesAvailableIpsListDefault(code int) *IpamPrefixesAvailableIpsListDefault {
	return &IpamPrefixesAvailableIpsListDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesAvailableIpsListDefault describes a response with status code -1, with default header values.

IpamPrefixesAvailableIpsListDefault ipam prefixes available ips list default
*/
type IpamPrefixesAvailableIpsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes available ips list default response has a 2xx status code
func (o *IpamPrefixesAvailableIpsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes available ips list default response has a 3xx status code
func (o *IpamPrefixesAvailableIpsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes available ips list default response has a 4xx status code
func (o *IpamPrefixesAvailableIpsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes available ips list default response has a 5xx status code
func (o *IpamPrefixesAvailableIpsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes available ips list default response a status code equal to that given
func (o *IpamPrefixesAvailableIpsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam prefixes available ips list default response
func (o *IpamPrefixesAvailableIpsListDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesAvailableIpsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipam_prefixes_available-ips_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailableIpsListDefault) String() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipam_prefixes_available-ips_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailableIpsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}