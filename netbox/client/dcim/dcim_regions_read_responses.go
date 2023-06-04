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

// DcimRegionsReadReader is a Reader for the DcimRegionsRead structure.
type DcimRegionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsReadOK creates a DcimRegionsReadOK with default headers values
func NewDcimRegionsReadOK() *DcimRegionsReadOK {
	return &DcimRegionsReadOK{}
}

/*
DcimRegionsReadOK describes a response with status code 200, with default header values.

DcimRegionsReadOK dcim regions read o k
*/
type DcimRegionsReadOK struct {
	Payload *models.Region
}

// IsSuccess returns true when this dcim regions read o k response has a 2xx status code
func (o *DcimRegionsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions read o k response has a 3xx status code
func (o *DcimRegionsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions read o k response has a 4xx status code
func (o *DcimRegionsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions read o k response has a 5xx status code
func (o *DcimRegionsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions read o k response a status code equal to that given
func (o *DcimRegionsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim regions read o k response
func (o *DcimRegionsReadOK) Code() int {
	return 200
}

func (o *DcimRegionsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/regions/{id}/][%d] dcimRegionsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/regions/{id}/][%d] dcimRegionsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsReadOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsReadDefault creates a DcimRegionsReadDefault with default headers values
func NewDcimRegionsReadDefault(code int) *DcimRegionsReadDefault {
	return &DcimRegionsReadDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsReadDefault describes a response with status code -1, with default header values.

DcimRegionsReadDefault dcim regions read default
*/
type DcimRegionsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim regions read default response has a 2xx status code
func (o *DcimRegionsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions read default response has a 3xx status code
func (o *DcimRegionsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions read default response has a 4xx status code
func (o *DcimRegionsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions read default response has a 5xx status code
func (o *DcimRegionsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions read default response a status code equal to that given
func (o *DcimRegionsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim regions read default response
func (o *DcimRegionsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRegionsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/regions/{id}/][%d] dcim_regions_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/regions/{id}/][%d] dcim_regions_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
