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

package roleapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// CreateParams is consumed by Create.
type CreateParams struct {
	*api.API

	Role *models.RoleAggregateCreateData
}

// Validate ensures the parameters are usable
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("role create")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Role == nil {
		merr = merr.Append(errors.New("role definition cannot be empty"))
	}

	return merr.ErrorOrNil()
}

// Create creates a new role.
func Create(params CreateParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.PlatformInfrastructure.CreateBlueprinterRole(
			platform_infrastructure.NewCreateBlueprinterRoleParams().
				WithBody(params.Role),
			params.AuthWriter,
		),
	)
}