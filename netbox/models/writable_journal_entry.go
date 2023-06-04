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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableJournalEntry writable journal entry
//
// swagger:model WritableJournalEntry
type WritableJournalEntry struct {

	// Assigned object
	// Read Only: true
	AssignedObject map[string]*string `json:"assigned_object,omitempty"`

	// Assigned object id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	AssignedObjectID *int64 `json:"assigned_object_id"`

	// Assigned object type
	// Required: true
	AssignedObjectType *string `json:"assigned_object_type"`

	// Comments
	// Required: true
	// Min Length: 1
	Comments *string `json:"comments"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Created by
	CreatedBy *int64 `json:"created_by,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Kind
	// Enum: [info success warning danger]
	Kind string `json:"kind,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable journal entry
func (m *WritableJournalEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableJournalEntry) validateAssignedObjectID(formats strfmt.Registry) error {

	if err := validate.Required("assigned_object_id", "body", m.AssignedObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("assigned_object_id", "body", *m.AssignedObjectID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("assigned_object_id", "body", *m.AssignedObjectID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) validateAssignedObjectType(formats strfmt.Registry) error {

	if err := validate.Required("assigned_object_type", "body", m.AssignedObjectType); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) validateComments(formats strfmt.Registry) error {

	if err := validate.Required("comments", "body", m.Comments); err != nil {
		return err
	}

	if err := validate.MinLength("comments", "body", *m.Comments, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

var writableJournalEntryTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","success","warning","danger"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableJournalEntryTypeKindPropEnum = append(writableJournalEntryTypeKindPropEnum, v)
	}
}

const (

	// WritableJournalEntryKindInfo captures enum value "info"
	WritableJournalEntryKindInfo string = "info"

	// WritableJournalEntryKindSuccess captures enum value "success"
	WritableJournalEntryKindSuccess string = "success"

	// WritableJournalEntryKindWarning captures enum value "warning"
	WritableJournalEntryKindWarning string = "warning"

	// WritableJournalEntryKindDanger captures enum value "danger"
	WritableJournalEntryKindDanger string = "danger"
)

// prop value enum
func (m *WritableJournalEntry) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableJournalEntryTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableJournalEntry) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable journal entry based on the context it is used
func (m *WritableJournalEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableJournalEntry) contextValidateAssignedObject(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableJournalEntry) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJournalEntry) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableJournalEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableJournalEntry) UnmarshalBinary(b []byte) error {
	var res WritableJournalEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
