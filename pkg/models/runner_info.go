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

// RunnerInfo Information about a specified runner.
//
// swagger:model RunnerInfo
type RunnerInfo struct {

	// Build Info of the container artifact
	BuildInfo *RunnerBuildInfo `json:"build_info,omitempty"`

	// Specifies if the runner is connected
	// Required: true
	Connected *bool `json:"connected"`

	// Containers for the runner
	// Required: true
	Containers []*RunnerContainerInfo `json:"containers"`

	// State of features of the runner
	Features map[string]bool `json:"features,omitempty"`

	// Checks used to determine if a runner is healthy or not
	HealthChecks *RunnerHealthChecks `json:"health_checks,omitempty"`

	// Specifies if the runner is healthy
	// Required: true
	Healthy *bool `json:"healthy"`

	// Host IP for the runner
	// Required: true
	HostIP *string `json:"host_ip"`

	// Public hostname for the runner
	// Required: true
	PublicHostname *string `json:"public_hostname"`

	// The region that this runner belongs to. Only populated in SaaS or federated ECE.
	Region string `json:"region,omitempty"`

	// Roles for the runner
	// Required: true
	Roles []*RunnerRoleInfo `json:"roles"`

	// The runner identifier
	// Required: true
	RunnerID *string `json:"runner_id"`

	// The runner name
	RunnerName string `json:"runner_name,omitempty"`

	// Meta data of the runner, like image ID or processor architecture
	Tags map[string]string `json:"tags,omitempty"`

	// Identifier of the zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this runner info
func (m *RunnerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunnerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunnerInfo) validateBuildInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildInfo) { // not required
		return nil
	}

	if m.BuildInfo != nil {
		if err := m.BuildInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_info")
			}
			return err
		}
	}

	return nil
}

func (m *RunnerInfo) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("connected", "body", m.Connected); err != nil {
		return err
	}

	return nil
}

func (m *RunnerInfo) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RunnerInfo) validateHealthChecks(formats strfmt.Registry) error {
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

func (m *RunnerInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *RunnerInfo) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *RunnerInfo) validatePublicHostname(formats strfmt.Registry) error {

	if err := validate.Required("public_hostname", "body", m.PublicHostname); err != nil {
		return err
	}

	return nil
}

func (m *RunnerInfo) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RunnerInfo) validateRunnerID(formats strfmt.Registry) error {

	if err := validate.Required("runner_id", "body", m.RunnerID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this runner info based on the context it is used
func (m *RunnerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunnerInfo) contextValidateBuildInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuildInfo != nil {
		if err := m.BuildInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_info")
			}
			return err
		}
	}

	return nil
}

func (m *RunnerInfo) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {
			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RunnerInfo) contextValidateHealthChecks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RunnerInfo) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {
			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunnerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunnerInfo) UnmarshalBinary(b []byte) error {
	var res RunnerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
