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
)

// DcimRearPortTemplatesDeleteReader is a Reader for the DcimRearPortTemplatesDelete structure.
type DcimRearPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesDeleteNoContent creates a DcimRearPortTemplatesDeleteNoContent with default headers values
func NewDcimRearPortTemplatesDeleteNoContent() *DcimRearPortTemplatesDeleteNoContent {
	return &DcimRearPortTemplatesDeleteNoContent{}
}

/*
DcimRearPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortTemplatesDeleteNoContent dcim rear port templates delete no content
*/
type DcimRearPortTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim rear port templates delete no content response has a 2xx status code
func (o *DcimRearPortTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear port templates delete no content response has a 3xx status code
func (o *DcimRearPortTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear port templates delete no content response has a 4xx status code
func (o *DcimRearPortTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear port templates delete no content response has a 5xx status code
func (o *DcimRearPortTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear port templates delete no content response a status code equal to that given
func (o *DcimRearPortTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim rear port templates delete no content response
func (o *DcimRearPortTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimRearPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimRearPortTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimRearPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRearPortTemplatesDeleteDefault creates a DcimRearPortTemplatesDeleteDefault with default headers values
func NewDcimRearPortTemplatesDeleteDefault(code int) *DcimRearPortTemplatesDeleteDefault {
	return &DcimRearPortTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimRearPortTemplatesDeleteDefault dcim rear port templates delete default
*/
type DcimRearPortTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear port templates delete default response has a 2xx status code
func (o *DcimRearPortTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear port templates delete default response has a 3xx status code
func (o *DcimRearPortTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear port templates delete default response has a 4xx status code
func (o *DcimRearPortTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear port templates delete default response has a 5xx status code
func (o *DcimRearPortTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear port templates delete default response a status code equal to that given
func (o *DcimRearPortTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear port templates delete default response
func (o *DcimRearPortTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
