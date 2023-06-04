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

// ExtrasJournalEntriesPartialUpdateReader is a Reader for the ExtrasJournalEntriesPartialUpdate structure.
type ExtrasJournalEntriesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesPartialUpdateOK creates a ExtrasJournalEntriesPartialUpdateOK with default headers values
func NewExtrasJournalEntriesPartialUpdateOK() *ExtrasJournalEntriesPartialUpdateOK {
	return &ExtrasJournalEntriesPartialUpdateOK{}
}

/*
ExtrasJournalEntriesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesPartialUpdateOK extras journal entries partial update o k
*/
type ExtrasJournalEntriesPartialUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries partial update o k response has a 2xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries partial update o k response has a 3xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries partial update o k response has a 4xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries partial update o k response has a 5xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries partial update o k response a status code equal to that given
func (o *ExtrasJournalEntriesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries partial update o k response
func (o *ExtrasJournalEntriesPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extrasJournalEntriesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extrasJournalEntriesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesPartialUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesPartialUpdateDefault creates a ExtrasJournalEntriesPartialUpdateDefault with default headers values
func NewExtrasJournalEntriesPartialUpdateDefault(code int) *ExtrasJournalEntriesPartialUpdateDefault {
	return &ExtrasJournalEntriesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasJournalEntriesPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesPartialUpdateDefault extras journal entries partial update default
*/
type ExtrasJournalEntriesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras journal entries partial update default response has a 2xx status code
func (o *ExtrasJournalEntriesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras journal entries partial update default response has a 3xx status code
func (o *ExtrasJournalEntriesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras journal entries partial update default response has a 4xx status code
func (o *ExtrasJournalEntriesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras journal entries partial update default response has a 5xx status code
func (o *ExtrasJournalEntriesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras journal entries partial update default response a status code equal to that given
func (o *ExtrasJournalEntriesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras journal entries partial update default response
func (o *ExtrasJournalEntriesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extras_journal-entries_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extras_journal-entries_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}