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

// ExtrasExportTemplatesReadReader is a Reader for the ExtrasExportTemplatesRead structure.
type ExtrasExportTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesReadOK creates a ExtrasExportTemplatesReadOK with default headers values
func NewExtrasExportTemplatesReadOK() *ExtrasExportTemplatesReadOK {
	return &ExtrasExportTemplatesReadOK{}
}

/*
ExtrasExportTemplatesReadOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesReadOK extras export templates read o k
*/
type ExtrasExportTemplatesReadOK struct {
	Payload *models.ExportTemplate
}

// IsSuccess returns true when this extras export templates read o k response has a 2xx status code
func (o *ExtrasExportTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates read o k response has a 3xx status code
func (o *ExtrasExportTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates read o k response has a 4xx status code
func (o *ExtrasExportTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates read o k response has a 5xx status code
func (o *ExtrasExportTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates read o k response a status code equal to that given
func (o *ExtrasExportTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras export templates read o k response
func (o *ExtrasExportTemplatesReadOK) Code() int {
	return 200
}

func (o *ExtrasExportTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/export-templates/{id}/][%d] extrasExportTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/export-templates/{id}/][%d] extrasExportTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesReadOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesReadDefault creates a ExtrasExportTemplatesReadDefault with default headers values
func NewExtrasExportTemplatesReadDefault(code int) *ExtrasExportTemplatesReadDefault {
	return &ExtrasExportTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesReadDefault describes a response with status code -1, with default header values.

ExtrasExportTemplatesReadDefault extras export templates read default
*/
type ExtrasExportTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras export templates read default response has a 2xx status code
func (o *ExtrasExportTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates read default response has a 3xx status code
func (o *ExtrasExportTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates read default response has a 4xx status code
func (o *ExtrasExportTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates read default response has a 5xx status code
func (o *ExtrasExportTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates read default response a status code equal to that given
func (o *ExtrasExportTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras export templates read default response
func (o *ExtrasExportTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/export-templates/{id}/][%d] extras_export-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/export-templates/{id}/][%d] extras_export-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
