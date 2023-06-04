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

// WirelessWirelessLansBulkUpdateReader is a Reader for the WirelessWirelessLansBulkUpdate structure.
type WirelessWirelessLansBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansBulkUpdateOK creates a WirelessWirelessLansBulkUpdateOK with default headers values
func NewWirelessWirelessLansBulkUpdateOK() *WirelessWirelessLansBulkUpdateOK {
	return &WirelessWirelessLansBulkUpdateOK{}
}

/*
WirelessWirelessLansBulkUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLansBulkUpdateOK wireless wireless lans bulk update o k
*/
type WirelessWirelessLansBulkUpdateOK struct {
	Payload *models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans bulk update o k response has a 2xx status code
func (o *WirelessWirelessLansBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans bulk update o k response has a 3xx status code
func (o *WirelessWirelessLansBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans bulk update o k response has a 4xx status code
func (o *WirelessWirelessLansBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans bulk update o k response has a 5xx status code
func (o *WirelessWirelessLansBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans bulk update o k response a status code equal to that given
func (o *WirelessWirelessLansBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless lans bulk update o k response
func (o *WirelessWirelessLansBulkUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLansBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lans/][%d] wirelessWirelessLansBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lans/][%d] wirelessWirelessLansBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansBulkUpdateOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLansBulkUpdateDefault creates a WirelessWirelessLansBulkUpdateDefault with default headers values
func NewWirelessWirelessLansBulkUpdateDefault(code int) *WirelessWirelessLansBulkUpdateDefault {
	return &WirelessWirelessLansBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLansBulkUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLansBulkUpdateDefault wireless wireless lans bulk update default
*/
type WirelessWirelessLansBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless lans bulk update default response has a 2xx status code
func (o *WirelessWirelessLansBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lans bulk update default response has a 3xx status code
func (o *WirelessWirelessLansBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lans bulk update default response has a 4xx status code
func (o *WirelessWirelessLansBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lans bulk update default response has a 5xx status code
func (o *WirelessWirelessLansBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lans bulk update default response a status code equal to that given
func (o *WirelessWirelessLansBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless lans bulk update default response
func (o *WirelessWirelessLansBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLansBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lans/][%d] wireless_wireless-lans_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lans/][%d] wireless_wireless-lans_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
