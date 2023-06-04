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

// ExtrasTagsUpdateReader is a Reader for the ExtrasTagsUpdate structure.
type ExtrasTagsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsUpdateOK creates a ExtrasTagsUpdateOK with default headers values
func NewExtrasTagsUpdateOK() *ExtrasTagsUpdateOK {
	return &ExtrasTagsUpdateOK{}
}

/*
ExtrasTagsUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsUpdateOK extras tags update o k
*/
type ExtrasTagsUpdateOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this extras tags update o k response has a 2xx status code
func (o *ExtrasTagsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags update o k response has a 3xx status code
func (o *ExtrasTagsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags update o k response has a 4xx status code
func (o *ExtrasTagsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags update o k response has a 5xx status code
func (o *ExtrasTagsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags update o k response a status code equal to that given
func (o *ExtrasTagsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras tags update o k response
func (o *ExtrasTagsUpdateOK) Code() int {
	return 200
}

func (o *ExtrasTagsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extrasTagsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extrasTagsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsUpdateDefault creates a ExtrasTagsUpdateDefault with default headers values
func NewExtrasTagsUpdateDefault(code int) *ExtrasTagsUpdateDefault {
	return &ExtrasTagsUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasTagsUpdateDefault describes a response with status code -1, with default header values.

ExtrasTagsUpdateDefault extras tags update default
*/
type ExtrasTagsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras tags update default response has a 2xx status code
func (o *ExtrasTagsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras tags update default response has a 3xx status code
func (o *ExtrasTagsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras tags update default response has a 4xx status code
func (o *ExtrasTagsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras tags update default response has a 5xx status code
func (o *ExtrasTagsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras tags update default response a status code equal to that given
func (o *ExtrasTagsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras tags update default response
func (o *ExtrasTagsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasTagsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extras_tags_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extras_tags_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
