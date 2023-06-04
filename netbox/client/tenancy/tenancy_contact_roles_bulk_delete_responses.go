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
)

// TenancyContactRolesBulkDeleteReader is a Reader for the TenancyContactRolesBulkDelete structure.
type TenancyContactRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesBulkDeleteNoContent creates a TenancyContactRolesBulkDeleteNoContent with default headers values
func NewTenancyContactRolesBulkDeleteNoContent() *TenancyContactRolesBulkDeleteNoContent {
	return &TenancyContactRolesBulkDeleteNoContent{}
}

/*
TenancyContactRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactRolesBulkDeleteNoContent tenancy contact roles bulk delete no content
*/
type TenancyContactRolesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy contact roles bulk delete no content response has a 2xx status code
func (o *TenancyContactRolesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles bulk delete no content response has a 3xx status code
func (o *TenancyContactRolesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles bulk delete no content response has a 4xx status code
func (o *TenancyContactRolesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles bulk delete no content response has a 5xx status code
func (o *TenancyContactRolesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles bulk delete no content response a status code equal to that given
func (o *TenancyContactRolesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the tenancy contact roles bulk delete no content response
func (o *TenancyContactRolesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *TenancyContactRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/][%d] tenancyContactRolesBulkDeleteNoContent ", 204)
}

func (o *TenancyContactRolesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/][%d] tenancyContactRolesBulkDeleteNoContent ", 204)
}

func (o *TenancyContactRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactRolesBulkDeleteDefault creates a TenancyContactRolesBulkDeleteDefault with default headers values
func NewTenancyContactRolesBulkDeleteDefault(code int) *TenancyContactRolesBulkDeleteDefault {
	return &TenancyContactRolesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesBulkDeleteDefault describes a response with status code -1, with default header values.

TenancyContactRolesBulkDeleteDefault tenancy contact roles bulk delete default
*/
type TenancyContactRolesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact roles bulk delete default response has a 2xx status code
func (o *TenancyContactRolesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles bulk delete default response has a 3xx status code
func (o *TenancyContactRolesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles bulk delete default response has a 4xx status code
func (o *TenancyContactRolesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles bulk delete default response has a 5xx status code
func (o *TenancyContactRolesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles bulk delete default response a status code equal to that given
func (o *TenancyContactRolesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact roles bulk delete default response
func (o *TenancyContactRolesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactRolesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
