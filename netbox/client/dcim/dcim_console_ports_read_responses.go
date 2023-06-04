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

// DcimConsolePortsReadReader is a Reader for the DcimConsolePortsRead structure.
type DcimConsolePortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsReadOK creates a DcimConsolePortsReadOK with default headers values
func NewDcimConsolePortsReadOK() *DcimConsolePortsReadOK {
	return &DcimConsolePortsReadOK{}
}

/*
DcimConsolePortsReadOK describes a response with status code 200, with default header values.

DcimConsolePortsReadOK dcim console ports read o k
*/
type DcimConsolePortsReadOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports read o k response has a 2xx status code
func (o *DcimConsolePortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports read o k response has a 3xx status code
func (o *DcimConsolePortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports read o k response has a 4xx status code
func (o *DcimConsolePortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports read o k response has a 5xx status code
func (o *DcimConsolePortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports read o k response a status code equal to that given
func (o *DcimConsolePortsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console ports read o k response
func (o *DcimConsolePortsReadOK) Code() int {
	return 200
}

func (o *DcimConsolePortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsReadOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsReadDefault creates a DcimConsolePortsReadDefault with default headers values
func NewDcimConsolePortsReadDefault(code int) *DcimConsolePortsReadDefault {
	return &DcimConsolePortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortsReadDefault describes a response with status code -1, with default header values.

DcimConsolePortsReadDefault dcim console ports read default
*/
type DcimConsolePortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console ports read default response has a 2xx status code
func (o *DcimConsolePortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console ports read default response has a 3xx status code
func (o *DcimConsolePortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console ports read default response has a 4xx status code
func (o *DcimConsolePortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console ports read default response has a 5xx status code
func (o *DcimConsolePortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console ports read default response a status code equal to that given
func (o *DcimConsolePortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console ports read default response
func (o *DcimConsolePortsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcim_console-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcim_console-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
