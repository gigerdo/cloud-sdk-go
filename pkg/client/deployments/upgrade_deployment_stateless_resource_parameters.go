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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUpgradeDeploymentStatelessResourceParams creates a new UpgradeDeploymentStatelessResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeDeploymentStatelessResourceParams() *UpgradeDeploymentStatelessResourceParams {
	return &UpgradeDeploymentStatelessResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeDeploymentStatelessResourceParamsWithTimeout creates a new UpgradeDeploymentStatelessResourceParams object
// with the ability to set a timeout on a request.
func NewUpgradeDeploymentStatelessResourceParamsWithTimeout(timeout time.Duration) *UpgradeDeploymentStatelessResourceParams {
	return &UpgradeDeploymentStatelessResourceParams{
		timeout: timeout,
	}
}

// NewUpgradeDeploymentStatelessResourceParamsWithContext creates a new UpgradeDeploymentStatelessResourceParams object
// with the ability to set a context for a request.
func NewUpgradeDeploymentStatelessResourceParamsWithContext(ctx context.Context) *UpgradeDeploymentStatelessResourceParams {
	return &UpgradeDeploymentStatelessResourceParams{
		Context: ctx,
	}
}

// NewUpgradeDeploymentStatelessResourceParamsWithHTTPClient creates a new UpgradeDeploymentStatelessResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeDeploymentStatelessResourceParamsWithHTTPClient(client *http.Client) *UpgradeDeploymentStatelessResourceParams {
	return &UpgradeDeploymentStatelessResourceParams{
		HTTPClient: client,
	}
}

/* UpgradeDeploymentStatelessResourceParams contains all the parameters to send to the API endpoint
   for the upgrade deployment stateless resource operation.

   Typically these are written to a http.Request.
*/
type UpgradeDeploymentStatelessResourceParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* StatelessResourceKind.

	   The kind of stateless resource
	*/
	StatelessResourceKind string

	/* ValidateOnly.

	   When `true`, returns the update version without performing the upgrade
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade deployment stateless resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeDeploymentStatelessResourceParams) WithDefaults() *UpgradeDeploymentStatelessResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade deployment stateless resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeDeploymentStatelessResourceParams) SetDefaults() {
	var (
		validateOnlyDefault = bool(false)
	)

	val := UpgradeDeploymentStatelessResourceParams{
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithTimeout(timeout time.Duration) *UpgradeDeploymentStatelessResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithContext(ctx context.Context) *UpgradeDeploymentStatelessResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithHTTPClient(client *http.Client) *UpgradeDeploymentStatelessResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithDeploymentID(deploymentID string) *UpgradeDeploymentStatelessResourceParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithRefID(refID string) *UpgradeDeploymentStatelessResourceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithStatelessResourceKind adds the statelessResourceKind to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithStatelessResourceKind(statelessResourceKind string) *UpgradeDeploymentStatelessResourceParams {
	o.SetStatelessResourceKind(statelessResourceKind)
	return o
}

// SetStatelessResourceKind adds the statelessResourceKind to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetStatelessResourceKind(statelessResourceKind string) {
	o.StatelessResourceKind = statelessResourceKind
}

// WithValidateOnly adds the validateOnly to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) WithValidateOnly(validateOnly *bool) *UpgradeDeploymentStatelessResourceParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the upgrade deployment stateless resource params
func (o *UpgradeDeploymentStatelessResourceParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeDeploymentStatelessResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	// path param stateless_resource_kind
	if err := r.SetPathParam("stateless_resource_kind", o.StatelessResourceKind); err != nil {
		return err
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
