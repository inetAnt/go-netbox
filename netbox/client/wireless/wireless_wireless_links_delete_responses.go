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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLinksDeleteReader is a Reader for the WirelessWirelessLinksDelete structure.
type WirelessWirelessLinksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLinksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksDeleteNoContent creates a WirelessWirelessLinksDeleteNoContent with default headers values
func NewWirelessWirelessLinksDeleteNoContent() *WirelessWirelessLinksDeleteNoContent {
	return &WirelessWirelessLinksDeleteNoContent{}
}

/*
WirelessWirelessLinksDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLinksDeleteNoContent wireless wireless links delete no content
*/
type WirelessWirelessLinksDeleteNoContent struct {
}

// IsSuccess returns true when this wireless wireless links delete no content response has a 2xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links delete no content response has a 3xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links delete no content response has a 4xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links delete no content response has a 5xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links delete no content response a status code equal to that given
func (o *WirelessWirelessLinksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the wireless wireless links delete no content response
func (o *WirelessWirelessLinksDeleteNoContent) Code() int {
	return 204
}

func (o *WirelessWirelessLinksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksDeleteNoContent ", 204)
}

func (o *WirelessWirelessLinksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksDeleteNoContent ", 204)
}

func (o *WirelessWirelessLinksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWirelessWirelessLinksDeleteDefault creates a WirelessWirelessLinksDeleteDefault with default headers values
func NewWirelessWirelessLinksDeleteDefault(code int) *WirelessWirelessLinksDeleteDefault {
	return &WirelessWirelessLinksDeleteDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksDeleteDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksDeleteDefault wireless wireless links delete default
*/
type WirelessWirelessLinksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless links delete default response has a 2xx status code
func (o *WirelessWirelessLinksDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links delete default response has a 3xx status code
func (o *WirelessWirelessLinksDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links delete default response has a 4xx status code
func (o *WirelessWirelessLinksDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links delete default response has a 5xx status code
func (o *WirelessWirelessLinksDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links delete default response a status code equal to that given
func (o *WirelessWirelessLinksDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless links delete default response
func (o *WirelessWirelessLinksDeleteDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLinksDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wireless_wireless-links_delete default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wireless_wireless-links_delete default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
