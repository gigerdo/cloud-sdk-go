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

// NewEnableDeploymentResourceSlmParams creates a new EnableDeploymentResourceSlmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableDeploymentResourceSlmParams() *EnableDeploymentResourceSlmParams {
	return &EnableDeploymentResourceSlmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableDeploymentResourceSlmParamsWithTimeout creates a new EnableDeploymentResourceSlmParams object
// with the ability to set a timeout on a request.
func NewEnableDeploymentResourceSlmParamsWithTimeout(timeout time.Duration) *EnableDeploymentResourceSlmParams {
	return &EnableDeploymentResourceSlmParams{
		timeout: timeout,
	}
}

// NewEnableDeploymentResourceSlmParamsWithContext creates a new EnableDeploymentResourceSlmParams object
// with the ability to set a context for a request.
func NewEnableDeploymentResourceSlmParamsWithContext(ctx context.Context) *EnableDeploymentResourceSlmParams {
	return &EnableDeploymentResourceSlmParams{
		Context: ctx,
	}
}

// NewEnableDeploymentResourceSlmParamsWithHTTPClient creates a new EnableDeploymentResourceSlmParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableDeploymentResourceSlmParamsWithHTTPClient(client *http.Client) *EnableDeploymentResourceSlmParams {
	return &EnableDeploymentResourceSlmParams{
		HTTPClient: client,
	}
}

/* EnableDeploymentResourceSlmParams contains all the parameters to send to the API endpoint
   for the enable deployment resource slm operation.

   Typically these are written to a http.Request.
*/
type EnableDeploymentResourceSlmParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ValidateOnly.

	   When `true`, does not enable SLM but returns warnings if any applications may lose availability during SLM migration.
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable deployment resource slm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDeploymentResourceSlmParams) WithDefaults() *EnableDeploymentResourceSlmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable deployment resource slm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDeploymentResourceSlmParams) SetDefaults() {
	var (
		validateOnlyDefault = bool(false)
	)

	val := EnableDeploymentResourceSlmParams{
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithTimeout(timeout time.Duration) *EnableDeploymentResourceSlmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithContext(ctx context.Context) *EnableDeploymentResourceSlmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithHTTPClient(client *http.Client) *EnableDeploymentResourceSlmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithDeploymentID(deploymentID string) *EnableDeploymentResourceSlmParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithRefID(refID string) *EnableDeploymentResourceSlmParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithValidateOnly adds the validateOnly to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) WithValidateOnly(validateOnly *bool) *EnableDeploymentResourceSlmParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the enable deployment resource slm params
func (o *EnableDeploymentResourceSlmParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *EnableDeploymentResourceSlmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
