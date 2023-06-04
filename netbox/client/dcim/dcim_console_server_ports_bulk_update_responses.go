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

// DcimConsoleServerPortsBulkUpdateReader is a Reader for the DcimConsoleServerPortsBulkUpdate structure.
type DcimConsoleServerPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsBulkUpdateOK creates a DcimConsoleServerPortsBulkUpdateOK with default headers values
func NewDcimConsoleServerPortsBulkUpdateOK() *DcimConsoleServerPortsBulkUpdateOK {
	return &DcimConsoleServerPortsBulkUpdateOK{}
}

/*
DcimConsoleServerPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsBulkUpdateOK dcim console server ports bulk update o k
*/
type DcimConsoleServerPortsBulkUpdateOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports bulk update o k response has a 2xx status code
func (o *DcimConsoleServerPortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports bulk update o k response has a 3xx status code
func (o *DcimConsoleServerPortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports bulk update o k response has a 4xx status code
func (o *DcimConsoleServerPortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports bulk update o k response has a 5xx status code
func (o *DcimConsoleServerPortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports bulk update o k response a status code equal to that given
func (o *DcimConsoleServerPortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports bulk update o k response
func (o *DcimConsoleServerPortsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsBulkUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsBulkUpdateDefault creates a DcimConsoleServerPortsBulkUpdateDefault with default headers values
func NewDcimConsoleServerPortsBulkUpdateDefault(code int) *DcimConsoleServerPortsBulkUpdateDefault {
	return &DcimConsoleServerPortsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsBulkUpdateDefault dcim console server ports bulk update default
*/
type DcimConsoleServerPortsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console server ports bulk update default response has a 2xx status code
func (o *DcimConsoleServerPortsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports bulk update default response has a 3xx status code
func (o *DcimConsoleServerPortsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports bulk update default response has a 4xx status code
func (o *DcimConsoleServerPortsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports bulk update default response has a 5xx status code
func (o *DcimConsoleServerPortsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports bulk update default response a status code equal to that given
func (o *DcimConsoleServerPortsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console server ports bulk update default response
func (o *DcimConsoleServerPortsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/][%d] dcim_console-server-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/][%d] dcim_console-server-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
