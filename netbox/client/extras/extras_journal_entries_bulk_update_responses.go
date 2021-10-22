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

	"github.com/digitalocean/go-netbox/netbox/models"
)

// ExtrasJournalEntriesBulkUpdateReader is a Reader for the ExtrasJournalEntriesBulkUpdate structure.
type ExtrasJournalEntriesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJournalEntriesBulkUpdateOK creates a ExtrasJournalEntriesBulkUpdateOK with default headers values
func NewExtrasJournalEntriesBulkUpdateOK() *ExtrasJournalEntriesBulkUpdateOK {
	return &ExtrasJournalEntriesBulkUpdateOK{}
}

/* ExtrasJournalEntriesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesBulkUpdateOK extras journal entries bulk update o k
*/
type ExtrasJournalEntriesBulkUpdateOK struct {
	Payload *models.JournalEntry
}

func (o *ExtrasJournalEntriesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasJournalEntriesBulkUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
