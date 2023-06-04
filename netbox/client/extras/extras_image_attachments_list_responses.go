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

// ExtrasImageAttachmentsListReader is a Reader for the ExtrasImageAttachmentsList structure.
type ExtrasImageAttachmentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsListOK creates a ExtrasImageAttachmentsListOK with default headers values
func NewExtrasImageAttachmentsListOK() *ExtrasImageAttachmentsListOK {
	return &ExtrasImageAttachmentsListOK{}
}

/*
ExtrasImageAttachmentsListOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsListOK extras image attachments list o k
*/
type ExtrasImageAttachmentsListOK struct {
	Payload *ExtrasImageAttachmentsListOKBody
}

// IsSuccess returns true when this extras image attachments list o k response has a 2xx status code
func (o *ExtrasImageAttachmentsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments list o k response has a 3xx status code
func (o *ExtrasImageAttachmentsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments list o k response has a 4xx status code
func (o *ExtrasImageAttachmentsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments list o k response has a 5xx status code
func (o *ExtrasImageAttachmentsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments list o k response a status code equal to that given
func (o *ExtrasImageAttachmentsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras image attachments list o k response
func (o *ExtrasImageAttachmentsListOK) Code() int {
	return 200
}

func (o *ExtrasImageAttachmentsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/image-attachments/][%d] extrasImageAttachmentsListOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsListOK) String() string {
	return fmt.Sprintf("[GET /extras/image-attachments/][%d] extrasImageAttachmentsListOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsListOK) GetPayload() *ExtrasImageAttachmentsListOKBody {
	return o.Payload
}

func (o *ExtrasImageAttachmentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasImageAttachmentsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsListDefault creates a ExtrasImageAttachmentsListDefault with default headers values
func NewExtrasImageAttachmentsListDefault(code int) *ExtrasImageAttachmentsListDefault {
	return &ExtrasImageAttachmentsListDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsListDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsListDefault extras image attachments list default
*/
type ExtrasImageAttachmentsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras image attachments list default response has a 2xx status code
func (o *ExtrasImageAttachmentsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments list default response has a 3xx status code
func (o *ExtrasImageAttachmentsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments list default response has a 4xx status code
func (o *ExtrasImageAttachmentsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments list default response has a 5xx status code
func (o *ExtrasImageAttachmentsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments list default response a status code equal to that given
func (o *ExtrasImageAttachmentsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras image attachments list default response
func (o *ExtrasImageAttachmentsListDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsListDefault) Error() string {
	return fmt.Sprintf("[GET /extras/image-attachments/][%d] extras_image-attachments_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsListDefault) String() string {
	return fmt.Sprintf("[GET /extras/image-attachments/][%d] extras_image-attachments_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasImageAttachmentsListOKBody extras image attachments list o k body
swagger:model ExtrasImageAttachmentsListOKBody
*/
type ExtrasImageAttachmentsListOKBody struct {

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
	Results []*models.ImageAttachment `json:"results"`
}

// Validate validates this extras image attachments list o k body
func (o *ExtrasImageAttachmentsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExtrasImageAttachmentsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasImageAttachmentsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasImageAttachmentsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasImageAttachmentsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasImageAttachmentsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasImageAttachmentsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasImageAttachmentsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasImageAttachmentsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasImageAttachmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasImageAttachmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras image attachments list o k body based on the context it is used
func (o *ExtrasImageAttachmentsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasImageAttachmentsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasImageAttachmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasImageAttachmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasImageAttachmentsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasImageAttachmentsListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasImageAttachmentsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
