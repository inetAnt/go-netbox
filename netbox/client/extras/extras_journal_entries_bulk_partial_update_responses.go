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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// ExtrasJournalEntriesBulkPartialUpdateReader is a Reader for the ExtrasJournalEntriesBulkPartialUpdate structure.
type ExtrasJournalEntriesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesBulkPartialUpdateOK creates a ExtrasJournalEntriesBulkPartialUpdateOK with default headers values
func NewExtrasJournalEntriesBulkPartialUpdateOK() *ExtrasJournalEntriesBulkPartialUpdateOK {
	return &ExtrasJournalEntriesBulkPartialUpdateOK{}
}

/*
ExtrasJournalEntriesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesBulkPartialUpdateOK extras journal entries bulk partial update o k
*/
type ExtrasJournalEntriesBulkPartialUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries bulk partial update o k response has a 2xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries bulk partial update o k response has a 3xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries bulk partial update o k response has a 4xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries bulk partial update o k response has a 5xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries bulk partial update o k response a status code equal to that given
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries bulk partial update o k response
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extrasJournalEntriesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extrasJournalEntriesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesBulkPartialUpdateDefault creates a ExtrasJournalEntriesBulkPartialUpdateDefault with default headers values
func NewExtrasJournalEntriesBulkPartialUpdateDefault(code int) *ExtrasJournalEntriesBulkPartialUpdateDefault {
	return &ExtrasJournalEntriesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasJournalEntriesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesBulkPartialUpdateDefault extras journal entries bulk partial update default
*/
type ExtrasJournalEntriesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras journal entries bulk partial update default response has a 2xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras journal entries bulk partial update default response has a 3xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras journal entries bulk partial update default response has a 4xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras journal entries bulk partial update default response has a 5xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras journal entries bulk partial update default response a status code equal to that given
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras journal entries bulk partial update default response
func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extras_journal-entries_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extras_journal-entries_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
