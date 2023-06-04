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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// ExtrasWebhooksPartialUpdateReader is a Reader for the ExtrasWebhooksPartialUpdate structure.
type ExtrasWebhooksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksPartialUpdateOK creates a ExtrasWebhooksPartialUpdateOK with default headers values
func NewExtrasWebhooksPartialUpdateOK() *ExtrasWebhooksPartialUpdateOK {
	return &ExtrasWebhooksPartialUpdateOK{}
}

/*
ExtrasWebhooksPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksPartialUpdateOK extras webhooks partial update o k
*/
type ExtrasWebhooksPartialUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks partial update o k response has a 2xx status code
func (o *ExtrasWebhooksPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks partial update o k response has a 3xx status code
func (o *ExtrasWebhooksPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks partial update o k response has a 4xx status code
func (o *ExtrasWebhooksPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks partial update o k response has a 5xx status code
func (o *ExtrasWebhooksPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks partial update o k response a status code equal to that given
func (o *ExtrasWebhooksPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks partial update o k response
func (o *ExtrasWebhooksPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/{id}/][%d] extrasWebhooksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/{id}/][%d] extrasWebhooksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksPartialUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksPartialUpdateDefault creates a ExtrasWebhooksPartialUpdateDefault with default headers values
func NewExtrasWebhooksPartialUpdateDefault(code int) *ExtrasWebhooksPartialUpdateDefault {
	return &ExtrasWebhooksPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasWebhooksPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasWebhooksPartialUpdateDefault extras webhooks partial update default
*/
type ExtrasWebhooksPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras webhooks partial update default response has a 2xx status code
func (o *ExtrasWebhooksPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras webhooks partial update default response has a 3xx status code
func (o *ExtrasWebhooksPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras webhooks partial update default response has a 4xx status code
func (o *ExtrasWebhooksPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras webhooks partial update default response has a 5xx status code
func (o *ExtrasWebhooksPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras webhooks partial update default response a status code equal to that given
func (o *ExtrasWebhooksPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras webhooks partial update default response
func (o *ExtrasWebhooksPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/{id}/][%d] extras_webhooks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/{id}/][%d] extras_webhooks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
