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

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// WirelessWirelessLinksCreateReader is a Reader for the WirelessWirelessLinksCreate structure.
type WirelessWirelessLinksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWirelessWirelessLinksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksCreateCreated creates a WirelessWirelessLinksCreateCreated with default headers values
func NewWirelessWirelessLinksCreateCreated() *WirelessWirelessLinksCreateCreated {
	return &WirelessWirelessLinksCreateCreated{}
}

/*
WirelessWirelessLinksCreateCreated describes a response with status code 201, with default header values.

WirelessWirelessLinksCreateCreated wireless wireless links create created
*/
type WirelessWirelessLinksCreateCreated struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links create created response has a 2xx status code
func (o *WirelessWirelessLinksCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links create created response has a 3xx status code
func (o *WirelessWirelessLinksCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links create created response has a 4xx status code
func (o *WirelessWirelessLinksCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links create created response has a 5xx status code
func (o *WirelessWirelessLinksCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links create created response a status code equal to that given
func (o *WirelessWirelessLinksCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the wireless wireless links create created response
func (o *WirelessWirelessLinksCreateCreated) Code() int {
	return 201
}

func (o *WirelessWirelessLinksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-links/][%d] wirelessWirelessLinksCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLinksCreateCreated) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-links/][%d] wirelessWirelessLinksCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLinksCreateCreated) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLinksCreateDefault creates a WirelessWirelessLinksCreateDefault with default headers values
func NewWirelessWirelessLinksCreateDefault(code int) *WirelessWirelessLinksCreateDefault {
	return &WirelessWirelessLinksCreateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksCreateDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksCreateDefault wireless wireless links create default
*/
type WirelessWirelessLinksCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless links create default response has a 2xx status code
func (o *WirelessWirelessLinksCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links create default response has a 3xx status code
func (o *WirelessWirelessLinksCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links create default response has a 4xx status code
func (o *WirelessWirelessLinksCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links create default response has a 5xx status code
func (o *WirelessWirelessLinksCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links create default response a status code equal to that given
func (o *WirelessWirelessLinksCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless links create default response
func (o *WirelessWirelessLinksCreateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLinksCreateDefault) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-links/][%d] wireless_wireless-links_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksCreateDefault) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-links/][%d] wireless_wireless-links_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
