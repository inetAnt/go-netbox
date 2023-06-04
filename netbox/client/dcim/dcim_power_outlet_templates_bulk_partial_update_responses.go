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

// DcimPowerOutletTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerOutletTemplatesBulkPartialUpdate structure.
type DcimPowerOutletTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesBulkPartialUpdateOK creates a DcimPowerOutletTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletTemplatesBulkPartialUpdateOK() *DcimPowerOutletTemplatesBulkPartialUpdateOK {
	return &DcimPowerOutletTemplatesBulkPartialUpdateOK{}
}

/*
DcimPowerOutletTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesBulkPartialUpdateOK dcim power outlet templates bulk partial update o k
*/
type DcimPowerOutletTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

// IsSuccess returns true when this dcim power outlet templates bulk partial update o k response has a 2xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates bulk partial update o k response has a 3xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates bulk partial update o k response has a 4xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates bulk partial update o k response has a 5xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates bulk partial update o k response a status code equal to that given
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlet templates bulk partial update o k response
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesBulkPartialUpdateDefault creates a DcimPowerOutletTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimPowerOutletTemplatesBulkPartialUpdateDefault(code int) *DcimPowerOutletTemplatesBulkPartialUpdateDefault {
	return &DcimPowerOutletTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesBulkPartialUpdateDefault dcim power outlet templates bulk partial update default
*/
type DcimPowerOutletTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlet templates bulk partial update default response has a 2xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates bulk partial update default response has a 3xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates bulk partial update default response has a 4xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates bulk partial update default response has a 5xx status code
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates bulk partial update default response a status code equal to that given
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlet templates bulk partial update default response
func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
