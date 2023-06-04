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

// DcimConsolePortsBulkUpdateReader is a Reader for the DcimConsolePortsBulkUpdate structure.
type DcimConsolePortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsBulkUpdateOK creates a DcimConsolePortsBulkUpdateOK with default headers values
func NewDcimConsolePortsBulkUpdateOK() *DcimConsolePortsBulkUpdateOK {
	return &DcimConsolePortsBulkUpdateOK{}
}

/*
DcimConsolePortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsBulkUpdateOK dcim console ports bulk update o k
*/
type DcimConsolePortsBulkUpdateOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports bulk update o k response has a 2xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports bulk update o k response has a 3xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports bulk update o k response has a 4xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports bulk update o k response has a 5xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports bulk update o k response a status code equal to that given
func (o *DcimConsolePortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console ports bulk update o k response
func (o *DcimConsolePortsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimConsolePortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcimConsolePortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcimConsolePortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsBulkUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsBulkUpdateDefault creates a DcimConsolePortsBulkUpdateDefault with default headers values
func NewDcimConsolePortsBulkUpdateDefault(code int) *DcimConsolePortsBulkUpdateDefault {
	return &DcimConsolePortsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortsBulkUpdateDefault dcim console ports bulk update default
*/
type DcimConsolePortsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console ports bulk update default response has a 2xx status code
func (o *DcimConsolePortsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console ports bulk update default response has a 3xx status code
func (o *DcimConsolePortsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console ports bulk update default response has a 4xx status code
func (o *DcimConsolePortsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console ports bulk update default response has a 5xx status code
func (o *DcimConsolePortsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console ports bulk update default response a status code equal to that given
func (o *DcimConsolePortsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console ports bulk update default response
func (o *DcimConsolePortsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcim_console-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcim_console-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
