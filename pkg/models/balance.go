// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Balance The available balance for an organization
//
// swagger:model Balance
type Balance struct {

	// Available balance
	// Required: true
	Available *float64 `json:"available"`

	// A collection of order line items for for an organization
	// Required: true
	LineItems []*SimplifiedLineItem `json:"line_items"`

	// Remaining balance
	// Required: true
	Remaining *float64 `json:"remaining"`
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemaining(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balance) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateLineItems(formats strfmt.Registry) error {

	if err := validate.Required("line_items", "body", m.LineItems); err != nil {
		return err
	}

	for i := 0; i < len(m.LineItems); i++ {
		if swag.IsZero(m.LineItems[i]) { // not required
			continue
		}

		if m.LineItems[i] != nil {
			if err := m.LineItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("line_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Balance) validateRemaining(formats strfmt.Registry) error {

	if err := validate.Required("remaining", "body", m.Remaining); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this balance based on the context it is used
func (m *Balance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balance) contextValidateLineItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineItems); i++ {

		if m.LineItems[i] != nil {
			if err := m.LineItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("line_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}