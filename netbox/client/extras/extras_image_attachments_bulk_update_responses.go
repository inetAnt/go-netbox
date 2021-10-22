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

// ExtrasImageAttachmentsBulkUpdateReader is a Reader for the ExtrasImageAttachmentsBulkUpdate structure.
type ExtrasImageAttachmentsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsBulkUpdateOK creates a ExtrasImageAttachmentsBulkUpdateOK with default headers values
func NewExtrasImageAttachmentsBulkUpdateOK() *ExtrasImageAttachmentsBulkUpdateOK {
	return &ExtrasImageAttachmentsBulkUpdateOK{}
}

/* ExtrasImageAttachmentsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsBulkUpdateOK extras image attachments bulk update o k
*/
type ExtrasImageAttachmentsBulkUpdateOK struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extrasImageAttachmentsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasImageAttachmentsBulkUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
