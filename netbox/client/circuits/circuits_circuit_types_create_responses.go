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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// CircuitsCircuitTypesCreateReader is a Reader for the CircuitsCircuitTypesCreate structure.
type CircuitsCircuitTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesCreateCreated creates a CircuitsCircuitTypesCreateCreated with default headers values
func NewCircuitsCircuitTypesCreateCreated() *CircuitsCircuitTypesCreateCreated {
	return &CircuitsCircuitTypesCreateCreated{}
}

/*
CircuitsCircuitTypesCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitTypesCreateCreated circuits circuit types create created
*/
type CircuitsCircuitTypesCreateCreated struct {
	Payload *models.CircuitType
}

// IsSuccess returns true when this circuits circuit types create created response has a 2xx status code
func (o *CircuitsCircuitTypesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types create created response has a 3xx status code
func (o *CircuitsCircuitTypesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types create created response has a 4xx status code
func (o *CircuitsCircuitTypesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types create created response has a 5xx status code
func (o *CircuitsCircuitTypesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types create created response a status code equal to that given
func (o *CircuitsCircuitTypesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the circuits circuit types create created response
func (o *CircuitsCircuitTypesCreateCreated) Code() int {
	return 201
}

func (o *CircuitsCircuitTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-types/][%d] circuitsCircuitTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsCircuitTypesCreateCreated) String() string {
	return fmt.Sprintf("[POST /circuits/circuit-types/][%d] circuitsCircuitTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsCircuitTypesCreateCreated) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesCreateDefault creates a CircuitsCircuitTypesCreateDefault with default headers values
func NewCircuitsCircuitTypesCreateDefault(code int) *CircuitsCircuitTypesCreateDefault {
	return &CircuitsCircuitTypesCreateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTypesCreateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTypesCreateDefault circuits circuit types create default
*/
type CircuitsCircuitTypesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuit types create default response has a 2xx status code
func (o *CircuitsCircuitTypesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit types create default response has a 3xx status code
func (o *CircuitsCircuitTypesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit types create default response has a 4xx status code
func (o *CircuitsCircuitTypesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit types create default response has a 5xx status code
func (o *CircuitsCircuitTypesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit types create default response a status code equal to that given
func (o *CircuitsCircuitTypesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuit types create default response
func (o *CircuitsCircuitTypesCreateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTypesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-types/][%d] circuits_circuit-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesCreateDefault) String() string {
	return fmt.Sprintf("[POST /circuits/circuit-types/][%d] circuits_circuit-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
