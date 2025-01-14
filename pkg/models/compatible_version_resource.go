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

// CompatibleVersionResource Information about the compatible version.
//
// swagger:model CompatibleVersionResource
type CompatibleVersionResource struct {

	// > WARNING
	// > This endpoint is deprecated and scheduled to be removed in the next major version. This field will soon be removed in favor of having a global capacity constraint for all versions. When specified its value will be ignored.
	//
	// Capacity constraints for the version
	CapacityConstraints *CapacityConstraintsResource `json:"capacity_constraints,omitempty"`

	// Supported node types for the version
	// Required: true
	NodeTypes []*CompatibleNodeTypesResource `json:"node_types"`

	// Compatible version, the key for this resource
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this compatible version resource
func (m *CompatibleVersionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompatibleVersionResource) validateCapacityConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityConstraints) { // not required
		return nil
	}

	if m.CapacityConstraints != nil {
		if err := m.CapacityConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_constraints")
			}
			return err
		}
	}

	return nil
}

func (m *CompatibleVersionResource) validateNodeTypes(formats strfmt.Registry) error {

	if err := validate.Required("node_types", "body", m.NodeTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeTypes); i++ {
		if swag.IsZero(m.NodeTypes[i]) { // not required
			continue
		}

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CompatibleVersionResource) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this compatible version resource based on the context it is used
func (m *CompatibleVersionResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacityConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompatibleVersionResource) contextValidateCapacityConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.CapacityConstraints != nil {
		if err := m.CapacityConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_constraints")
			}
			return err
		}
	}

	return nil
}

func (m *CompatibleVersionResource) contextValidateNodeTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeTypes); i++ {

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompatibleVersionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompatibleVersionResource) UnmarshalBinary(b []byte) error {
	var res CompatibleVersionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
