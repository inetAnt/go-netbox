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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inetAnt/go-netbox/v3/netbox/models"
)

// UsersTokensUpdateReader is a Reader for the UsersTokensUpdate structure.
type UsersTokensUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensUpdateOK creates a UsersTokensUpdateOK with default headers values
func NewUsersTokensUpdateOK() *UsersTokensUpdateOK {
	return &UsersTokensUpdateOK{}
}

/*
UsersTokensUpdateOK describes a response with status code 200, with default header values.

UsersTokensUpdateOK users tokens update o k
*/
type UsersTokensUpdateOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens update o k response has a 2xx status code
func (o *UsersTokensUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens update o k response has a 3xx status code
func (o *UsersTokensUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens update o k response has a 4xx status code
func (o *UsersTokensUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens update o k response has a 5xx status code
func (o *UsersTokensUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens update o k response a status code equal to that given
func (o *UsersTokensUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users tokens update o k response
func (o *UsersTokensUpdateOK) Code() int {
	return 200
}

func (o *UsersTokensUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] usersTokensUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] usersTokensUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensUpdateOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensUpdateDefault creates a UsersTokensUpdateDefault with default headers values
func NewUsersTokensUpdateDefault(code int) *UsersTokensUpdateDefault {
	return &UsersTokensUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersTokensUpdateDefault describes a response with status code -1, with default header values.

UsersTokensUpdateDefault users tokens update default
*/
type UsersTokensUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users tokens update default response has a 2xx status code
func (o *UsersTokensUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens update default response has a 3xx status code
func (o *UsersTokensUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens update default response has a 4xx status code
func (o *UsersTokensUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens update default response has a 5xx status code
func (o *UsersTokensUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens update default response a status code equal to that given
func (o *UsersTokensUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users tokens update default response
func (o *UsersTokensUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] users_tokens_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] users_tokens_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
