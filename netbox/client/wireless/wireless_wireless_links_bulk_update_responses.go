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

// WirelessWirelessLinksBulkUpdateReader is a Reader for the WirelessWirelessLinksBulkUpdate structure.
type WirelessWirelessLinksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksBulkUpdateOK creates a WirelessWirelessLinksBulkUpdateOK with default headers values
func NewWirelessWirelessLinksBulkUpdateOK() *WirelessWirelessLinksBulkUpdateOK {
	return &WirelessWirelessLinksBulkUpdateOK{}
}

/*
WirelessWirelessLinksBulkUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLinksBulkUpdateOK wireless wireless links bulk update o k
*/
type WirelessWirelessLinksBulkUpdateOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links bulk update o k response has a 2xx status code
func (o *WirelessWirelessLinksBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links bulk update o k response has a 3xx status code
func (o *WirelessWirelessLinksBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links bulk update o k response has a 4xx status code
func (o *WirelessWirelessLinksBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links bulk update o k response has a 5xx status code
func (o *WirelessWirelessLinksBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links bulk update o k response a status code equal to that given
func (o *WirelessWirelessLinksBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links bulk update o k response
func (o *WirelessWirelessLinksBulkUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/][%d] wirelessWirelessLinksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/][%d] wirelessWirelessLinksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksBulkUpdateOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLinksBulkUpdateDefault creates a WirelessWirelessLinksBulkUpdateDefault with default headers values
func NewWirelessWirelessLinksBulkUpdateDefault(code int) *WirelessWirelessLinksBulkUpdateDefault {
	return &WirelessWirelessLinksBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksBulkUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksBulkUpdateDefault wireless wireless links bulk update default
*/
type WirelessWirelessLinksBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless links bulk update default response has a 2xx status code
func (o *WirelessWirelessLinksBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links bulk update default response has a 3xx status code
func (o *WirelessWirelessLinksBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links bulk update default response has a 4xx status code
func (o *WirelessWirelessLinksBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links bulk update default response has a 5xx status code
func (o *WirelessWirelessLinksBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links bulk update default response a status code equal to that given
func (o *WirelessWirelessLinksBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless links bulk update default response
func (o *WirelessWirelessLinksBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLinksBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/][%d] wireless_wireless-links_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-links/][%d] wireless_wireless-links_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
