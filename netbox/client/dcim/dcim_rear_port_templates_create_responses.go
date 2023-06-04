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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// DcimRearPortTemplatesCreateReader is a Reader for the DcimRearPortTemplatesCreate structure.
type DcimRearPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesCreateCreated creates a DcimRearPortTemplatesCreateCreated with default headers values
func NewDcimRearPortTemplatesCreateCreated() *DcimRearPortTemplatesCreateCreated {
	return &DcimRearPortTemplatesCreateCreated{}
}

/*
DcimRearPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimRearPortTemplatesCreateCreated dcim rear port templates create created
*/
type DcimRearPortTemplatesCreateCreated struct {
	Payload *models.RearPortTemplate
}

// IsSuccess returns true when this dcim rear port templates create created response has a 2xx status code
func (o *DcimRearPortTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear port templates create created response has a 3xx status code
func (o *DcimRearPortTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear port templates create created response has a 4xx status code
func (o *DcimRearPortTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear port templates create created response has a 5xx status code
func (o *DcimRearPortTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear port templates create created response a status code equal to that given
func (o *DcimRearPortTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim rear port templates create created response
func (o *DcimRearPortTemplatesCreateCreated) Code() int {
	return 201
}

func (o *DcimRearPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcimRearPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRearPortTemplatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcimRearPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRearPortTemplatesCreateCreated) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesCreateDefault creates a DcimRearPortTemplatesCreateDefault with default headers values
func NewDcimRearPortTemplatesCreateDefault(code int) *DcimRearPortTemplatesCreateDefault {
	return &DcimRearPortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortTemplatesCreateDefault describes a response with status code -1, with default header values.

DcimRearPortTemplatesCreateDefault dcim rear port templates create default
*/
type DcimRearPortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear port templates create default response has a 2xx status code
func (o *DcimRearPortTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear port templates create default response has a 3xx status code
func (o *DcimRearPortTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear port templates create default response has a 4xx status code
func (o *DcimRearPortTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear port templates create default response has a 5xx status code
func (o *DcimRearPortTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear port templates create default response a status code equal to that given
func (o *DcimRearPortTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear port templates create default response
func (o *DcimRearPortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcim_rear-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcim_rear-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
