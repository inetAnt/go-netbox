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

// DcimDeviceBayTemplatesPartialUpdateReader is a Reader for the DcimDeviceBayTemplatesPartialUpdate structure.
type DcimDeviceBayTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateOK creates a DcimDeviceBayTemplatesPartialUpdateOK with default headers values
func NewDcimDeviceBayTemplatesPartialUpdateOK() *DcimDeviceBayTemplatesPartialUpdateOK {
	return &DcimDeviceBayTemplatesPartialUpdateOK{}
}

/*
DcimDeviceBayTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesPartialUpdateOK dcim device bay templates partial update o k
*/
type DcimDeviceBayTemplatesPartialUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates partial update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates partial update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates partial update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates partial update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates partial update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device bay templates partial update o k response
func (o *DcimDeviceBayTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesPartialUpdateDefault creates a DcimDeviceBayTemplatesPartialUpdateDefault with default headers values
func NewDcimDeviceBayTemplatesPartialUpdateDefault(code int) *DcimDeviceBayTemplatesPartialUpdateDefault {
	return &DcimDeviceBayTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesPartialUpdateDefault dcim device bay templates partial update default
*/
type DcimDeviceBayTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device bay templates partial update default response has a 2xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates partial update default response has a 3xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates partial update default response has a 4xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates partial update default response has a 5xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates partial update default response a status code equal to that given
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device bay templates partial update default response
func (o *DcimDeviceBayTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
