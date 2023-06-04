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
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// TenancyContactsListReader is a Reader for the TenancyContactsList structure.
type TenancyContactsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsListOK creates a TenancyContactsListOK with default headers values
func NewTenancyContactsListOK() *TenancyContactsListOK {
	return &TenancyContactsListOK{}
}

/*
TenancyContactsListOK describes a response with status code 200, with default header values.

TenancyContactsListOK tenancy contacts list o k
*/
type TenancyContactsListOK struct {
	Payload *TenancyContactsListOKBody
}

// IsSuccess returns true when this tenancy contacts list o k response has a 2xx status code
func (o *TenancyContactsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts list o k response has a 3xx status code
func (o *TenancyContactsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts list o k response has a 4xx status code
func (o *TenancyContactsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts list o k response has a 5xx status code
func (o *TenancyContactsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts list o k response a status code equal to that given
func (o *TenancyContactsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contacts list o k response
func (o *TenancyContactsListOK) Code() int {
	return 200
}

func (o *TenancyContactsListOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/contacts/][%d] tenancyContactsListOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsListOK) String() string {
	return fmt.Sprintf("[GET /tenancy/contacts/][%d] tenancyContactsListOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsListOK) GetPayload() *TenancyContactsListOKBody {
	return o.Payload
}

func (o *TenancyContactsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TenancyContactsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsListDefault creates a TenancyContactsListDefault with default headers values
func NewTenancyContactsListDefault(code int) *TenancyContactsListDefault {
	return &TenancyContactsListDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsListDefault describes a response with status code -1, with default header values.

TenancyContactsListDefault tenancy contacts list default
*/
type TenancyContactsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contacts list default response has a 2xx status code
func (o *TenancyContactsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts list default response has a 3xx status code
func (o *TenancyContactsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts list default response has a 4xx status code
func (o *TenancyContactsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts list default response has a 5xx status code
func (o *TenancyContactsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts list default response a status code equal to that given
func (o *TenancyContactsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contacts list default response
func (o *TenancyContactsListDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsListDefault) Error() string {
	return fmt.Sprintf("[GET /tenancy/contacts/][%d] tenancy_contacts_list default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsListDefault) String() string {
	return fmt.Sprintf("[GET /tenancy/contacts/][%d] tenancy_contacts_list default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
TenancyContactsListOKBody tenancy contacts list o k body
swagger:model TenancyContactsListOKBody
*/
type TenancyContactsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Contact `json:"results"`
}

// Validate validates this tenancy contacts list o k body
func (o *TenancyContactsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TenancyContactsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenancy contacts list o k body based on the context it is used
func (o *TenancyContactsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TenancyContactsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TenancyContactsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TenancyContactsListOKBody) UnmarshalBinary(b []byte) error {
	var res TenancyContactsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
