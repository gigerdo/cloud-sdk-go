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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllocatorHealthStatus The health status of the allocator.
//
// swagger:model AllocatorHealthStatus
type AllocatorHealthStatus struct {

	// Whether the allocator is connected
	// Required: true
	Connected *bool `json:"connected"`

	// Checks used to determine if an allocator is healthy or not
	HealthChecks *AllocatorHealthChecks `json:"health_checks,omitempty"`

	// Whether the allocator is healthy, meaning it is either connected or has no instances
	// Required: true
	Healthy *bool `json:"healthy"`

	// Whether the allocator is in maintenance mode (meaning that new workload won't be assigned to it)
	// Required: true
	MaintenanceMode *bool `json:"maintenance_mode"`

	// Timestamp when allocator last entered or exited maintenance mode
	// Format: date-time
	MaintenanceModeTimestamp strfmt.DateTime `json:"maintenance_mode_timestamp,omitempty"`
}

// Validate validates this allocator health status
func (m *AllocatorHealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceModeTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorHealthStatus) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("connected", "body", m.Connected); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorHealthStatus) validateHealthChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthChecks) { // not required
		return nil
	}

	if m.HealthChecks != nil {
		if err := m.HealthChecks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health_checks")
			}
			return err
		}
	}

	return nil
}

func (m *AllocatorHealthStatus) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorHealthStatus) validateMaintenanceMode(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_mode", "body", m.MaintenanceMode); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorHealthStatus) validateMaintenanceModeTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceModeTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("maintenance_mode_timestamp", "body", "date-time", m.MaintenanceModeTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this allocator health status based on the context it is used
func (m *AllocatorHealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealthChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorHealthStatus) contextValidateHealthChecks(ctx context.Context, formats strfmt.Registry) error {

	if m.HealthChecks != nil {
		if err := m.HealthChecks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health_checks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocatorHealthStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatorHealthStatus) UnmarshalBinary(b []byte) error {
	var res AllocatorHealthStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
