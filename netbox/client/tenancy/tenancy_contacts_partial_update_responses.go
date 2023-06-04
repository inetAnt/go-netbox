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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// TenancyContactsPartialUpdateReader is a Reader for the TenancyContactsPartialUpdate structure.
type TenancyContactsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsPartialUpdateOK creates a TenancyContactsPartialUpdateOK with default headers values
func NewTenancyContactsPartialUpdateOK() *TenancyContactsPartialUpdateOK {
	return &TenancyContactsPartialUpdateOK{}
}

/*
TenancyContactsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactsPartialUpdateOK tenancy contacts partial update o k
*/
type TenancyContactsPartialUpdateOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts partial update o k response has a 2xx status code
func (o *TenancyContactsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts partial update o k response has a 3xx status code
func (o *TenancyContactsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts partial update o k response has a 4xx status code
func (o *TenancyContactsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts partial update o k response has a 5xx status code
func (o *TenancyContactsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts partial update o k response a status code equal to that given
func (o *TenancyContactsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contacts partial update o k response
func (o *TenancyContactsPartialUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancyContactsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancyContactsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsPartialUpdateOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsPartialUpdateDefault creates a TenancyContactsPartialUpdateDefault with default headers values
func NewTenancyContactsPartialUpdateDefault(code int) *TenancyContactsPartialUpdateDefault {
	return &TenancyContactsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactsPartialUpdateDefault tenancy contacts partial update default
*/
type TenancyContactsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contacts partial update default response has a 2xx status code
func (o *TenancyContactsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts partial update default response has a 3xx status code
func (o *TenancyContactsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts partial update default response has a 4xx status code
func (o *TenancyContactsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts partial update default response has a 5xx status code
func (o *TenancyContactsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts partial update default response a status code equal to that given
func (o *TenancyContactsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contacts partial update default response
func (o *TenancyContactsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancy_contacts_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancy_contacts_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
