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

// NewGetDeploymentEsResourceInfoParams creates a new GetDeploymentEsResourceInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentEsResourceInfoParams() *GetDeploymentEsResourceInfoParams {
	return &GetDeploymentEsResourceInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentEsResourceInfoParamsWithTimeout creates a new GetDeploymentEsResourceInfoParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentEsResourceInfoParamsWithTimeout(timeout time.Duration) *GetDeploymentEsResourceInfoParams {
	return &GetDeploymentEsResourceInfoParams{
		timeout: timeout,
	}
}

// NewGetDeploymentEsResourceInfoParamsWithContext creates a new GetDeploymentEsResourceInfoParams object
// with the ability to set a context for a request.
func NewGetDeploymentEsResourceInfoParamsWithContext(ctx context.Context) *GetDeploymentEsResourceInfoParams {
	return &GetDeploymentEsResourceInfoParams{
		Context: ctx,
	}
}

// NewGetDeploymentEsResourceInfoParamsWithHTTPClient creates a new GetDeploymentEsResourceInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentEsResourceInfoParamsWithHTTPClient(client *http.Client) *GetDeploymentEsResourceInfoParams {
	return &GetDeploymentEsResourceInfoParams{
		HTTPClient: client,
	}
}

/* GetDeploymentEsResourceInfoParams contains all the parameters to send to the API endpoint
   for the get deployment es resource info operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentEsResourceInfoParams struct {

	/* ConvertLegacyPlans.

	   If showing plans, whether to leave pre-2.0.0 plans in their legacy format (the default), or whether to update them to 2.0.x+ format (if 'true').
	*/
	ConvertLegacyPlans *bool

	/* DeploymentID.

	   Identifier for the Deployment
	*/
	DeploymentID string

	/* EnrichWithTemplate.

	   If showing plans, whether to enrich the plan by including the missing elements from the deployment template it is based on.

	   Default: true
	*/
	EnrichWithTemplate *bool

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ShowMetadata.

	   Whether to include the full cluster metadata in the response - can be large per cluster and also include credentials.
	*/
	ShowMetadata *bool

	/* ShowPlanDefaults.

	   If showing plans, whether to show values that are left at their default value (less readable but more informative).
	*/
	ShowPlanDefaults *bool

	/* ShowPlanHistory.

	   Whether to include with the current and pending plan information the plan history- can be very large per cluster.
	*/
	ShowPlanHistory *bool

	/* ShowPlanLogs.

	   Whether to include with the current and pending plan information the attempt log - can be very large per cluster.
	*/
	ShowPlanLogs *bool

	/* ShowPlans.

	   Whether to include the full current and pending plan information in the response - can be large per cluster.

	   Default: true
	*/
	ShowPlans *bool

	/* ShowSecurity.

	   Whether to include the Elasticsearch 2.x security information in the response - can be large per cluster and also include credentials.
	*/
	ShowSecurity *bool

	/* ShowSettings.

	   Whether to show cluster settings in the response.
	*/
	ShowSettings *bool

	/* ShowSystemAlerts.

	   Number of system alerts (such as forced restarts due to memory limits) to be included in the response - can be large per cluster. Negative numbers or 0 will not return field.
	*/
	ShowSystemAlerts *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment es resource info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentEsResourceInfoParams) WithDefaults() *GetDeploymentEsResourceInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment es resource info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentEsResourceInfoParams) SetDefaults() {
	var (
		convertLegacyPlansDefault = bool(false)

		enrichWithTemplateDefault = bool(true)

		showMetadataDefault = bool(false)

		showPlanDefaultsDefault = bool(false)

		showPlanHistoryDefault = bool(false)

		showPlanLogsDefault = bool(false)

		showPlansDefault = bool(true)

		showSecurityDefault = bool(false)

		showSettingsDefault = bool(false)

		showSystemAlertsDefault = int64(0)
	)

	val := GetDeploymentEsResourceInfoParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowMetadata:       &showMetadataDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
		ShowPlanHistory:    &showPlanHistoryDefault,
		ShowPlanLogs:       &showPlanLogsDefault,
		ShowPlans:          &showPlansDefault,
		ShowSecurity:       &showSecurityDefault,
		ShowSettings:       &showSettingsDefault,
		ShowSystemAlerts:   &showSystemAlertsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithTimeout(timeout time.Duration) *GetDeploymentEsResourceInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithContext(ctx context.Context) *GetDeploymentEsResourceInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithHTTPClient(client *http.Client) *GetDeploymentEsResourceInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConvertLegacyPlans adds the convertLegacyPlans to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithConvertLegacyPlans(convertLegacyPlans *bool) *GetDeploymentEsResourceInfoParams {
	o.SetConvertLegacyPlans(convertLegacyPlans)
	return o
}

// SetConvertLegacyPlans adds the convertLegacyPlans to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetConvertLegacyPlans(convertLegacyPlans *bool) {
	o.ConvertLegacyPlans = convertLegacyPlans
}

// WithDeploymentID adds the deploymentID to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithDeploymentID(deploymentID string) *GetDeploymentEsResourceInfoParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithEnrichWithTemplate adds the enrichWithTemplate to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithEnrichWithTemplate(enrichWithTemplate *bool) *GetDeploymentEsResourceInfoParams {
	o.SetEnrichWithTemplate(enrichWithTemplate)
	return o
}

// SetEnrichWithTemplate adds the enrichWithTemplate to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetEnrichWithTemplate(enrichWithTemplate *bool) {
	o.EnrichWithTemplate = enrichWithTemplate
}

// WithRefID adds the refID to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithRefID(refID string) *GetDeploymentEsResourceInfoParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithShowMetadata adds the showMetadata to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowMetadata(showMetadata *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowMetadata(showMetadata)
	return o
}

// SetShowMetadata adds the showMetadata to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowMetadata(showMetadata *bool) {
	o.ShowMetadata = showMetadata
}

// WithShowPlanDefaults adds the showPlanDefaults to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WithShowPlanHistory adds the showPlanHistory to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowPlanHistory(showPlanHistory *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowPlanHistory(showPlanHistory)
	return o
}

// SetShowPlanHistory adds the showPlanHistory to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowPlanHistory(showPlanHistory *bool) {
	o.ShowPlanHistory = showPlanHistory
}

// WithShowPlanLogs adds the showPlanLogs to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowPlanLogs(showPlanLogs *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowPlanLogs(showPlanLogs)
	return o
}

// SetShowPlanLogs adds the showPlanLogs to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowPlanLogs(showPlanLogs *bool) {
	o.ShowPlanLogs = showPlanLogs
}

// WithShowPlans adds the showPlans to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowPlans(showPlans *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowPlans(showPlans)
	return o
}

// SetShowPlans adds the showPlans to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowPlans(showPlans *bool) {
	o.ShowPlans = showPlans
}

// WithShowSecurity adds the showSecurity to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowSecurity(showSecurity *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowSecurity(showSecurity)
	return o
}

// SetShowSecurity adds the showSecurity to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowSecurity(showSecurity *bool) {
	o.ShowSecurity = showSecurity
}

// WithShowSettings adds the showSettings to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowSettings(showSettings *bool) *GetDeploymentEsResourceInfoParams {
	o.SetShowSettings(showSettings)
	return o
}

// SetShowSettings adds the showSettings to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowSettings(showSettings *bool) {
	o.ShowSettings = showSettings
}

// WithShowSystemAlerts adds the showSystemAlerts to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) WithShowSystemAlerts(showSystemAlerts *int64) *GetDeploymentEsResourceInfoParams {
	o.SetShowSystemAlerts(showSystemAlerts)
	return o
}

// SetShowSystemAlerts adds the showSystemAlerts to the get deployment es resource info params
func (o *GetDeploymentEsResourceInfoParams) SetShowSystemAlerts(showSystemAlerts *int64) {
	o.ShowSystemAlerts = showSystemAlerts
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentEsResourceInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConvertLegacyPlans != nil {

		// query param convert_legacy_plans
		var qrConvertLegacyPlans bool

		if o.ConvertLegacyPlans != nil {
			qrConvertLegacyPlans = *o.ConvertLegacyPlans
		}
		qConvertLegacyPlans := swag.FormatBool(qrConvertLegacyPlans)
		if qConvertLegacyPlans != "" {

			if err := r.SetQueryParam("convert_legacy_plans", qConvertLegacyPlans); err != nil {
				return err
			}
		}
	}

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	if o.EnrichWithTemplate != nil {

		// query param enrich_with_template
		var qrEnrichWithTemplate bool

		if o.EnrichWithTemplate != nil {
			qrEnrichWithTemplate = *o.EnrichWithTemplate
		}
		qEnrichWithTemplate := swag.FormatBool(qrEnrichWithTemplate)
		if qEnrichWithTemplate != "" {

			if err := r.SetQueryParam("enrich_with_template", qEnrichWithTemplate); err != nil {
				return err
			}
		}
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	if o.ShowMetadata != nil {

		// query param show_metadata
		var qrShowMetadata bool

		if o.ShowMetadata != nil {
			qrShowMetadata = *o.ShowMetadata
		}
		qShowMetadata := swag.FormatBool(qrShowMetadata)
		if qShowMetadata != "" {

			if err := r.SetQueryParam("show_metadata", qShowMetadata); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanDefaults != nil {

		// query param show_plan_defaults
		var qrShowPlanDefaults bool

		if o.ShowPlanDefaults != nil {
			qrShowPlanDefaults = *o.ShowPlanDefaults
		}
		qShowPlanDefaults := swag.FormatBool(qrShowPlanDefaults)
		if qShowPlanDefaults != "" {

			if err := r.SetQueryParam("show_plan_defaults", qShowPlanDefaults); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanHistory != nil {

		// query param show_plan_history
		var qrShowPlanHistory bool

		if o.ShowPlanHistory != nil {
			qrShowPlanHistory = *o.ShowPlanHistory
		}
		qShowPlanHistory := swag.FormatBool(qrShowPlanHistory)
		if qShowPlanHistory != "" {

			if err := r.SetQueryParam("show_plan_history", qShowPlanHistory); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanLogs != nil {

		// query param show_plan_logs
		var qrShowPlanLogs bool

		if o.ShowPlanLogs != nil {
			qrShowPlanLogs = *o.ShowPlanLogs
		}
		qShowPlanLogs := swag.FormatBool(qrShowPlanLogs)
		if qShowPlanLogs != "" {

			if err := r.SetQueryParam("show_plan_logs", qShowPlanLogs); err != nil {
				return err
			}
		}
	}

	if o.ShowPlans != nil {

		// query param show_plans
		var qrShowPlans bool

		if o.ShowPlans != nil {
			qrShowPlans = *o.ShowPlans
		}
		qShowPlans := swag.FormatBool(qrShowPlans)
		if qShowPlans != "" {

			if err := r.SetQueryParam("show_plans", qShowPlans); err != nil {
				return err
			}
		}
	}

	if o.ShowSecurity != nil {

		// query param show_security
		var qrShowSecurity bool

		if o.ShowSecurity != nil {
			qrShowSecurity = *o.ShowSecurity
		}
		qShowSecurity := swag.FormatBool(qrShowSecurity)
		if qShowSecurity != "" {

			if err := r.SetQueryParam("show_security", qShowSecurity); err != nil {
				return err
			}
		}
	}

	if o.ShowSettings != nil {

		// query param show_settings
		var qrShowSettings bool

		if o.ShowSettings != nil {
			qrShowSettings = *o.ShowSettings
		}
		qShowSettings := swag.FormatBool(qrShowSettings)
		if qShowSettings != "" {

			if err := r.SetQueryParam("show_settings", qShowSettings); err != nil {
				return err
			}
		}
	}

	if o.ShowSystemAlerts != nil {

		// query param show_system_alerts
		var qrShowSystemAlerts int64

		if o.ShowSystemAlerts != nil {
			qrShowSystemAlerts = *o.ShowSystemAlerts
		}
		qShowSystemAlerts := swag.FormatInt64(qrShowSystemAlerts)
		if qShowSystemAlerts != "" {

			if err := r.SetQueryParam("show_system_alerts", qShowSystemAlerts); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
