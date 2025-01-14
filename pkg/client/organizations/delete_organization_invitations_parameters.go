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

package organizations

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
)

// NewDeleteOrganizationInvitationsParams creates a new DeleteOrganizationInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationInvitationsParams() *DeleteOrganizationInvitationsParams {
	return &DeleteOrganizationInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationInvitationsParamsWithTimeout creates a new DeleteOrganizationInvitationsParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationInvitationsParamsWithTimeout(timeout time.Duration) *DeleteOrganizationInvitationsParams {
	return &DeleteOrganizationInvitationsParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationInvitationsParamsWithContext creates a new DeleteOrganizationInvitationsParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationInvitationsParamsWithContext(ctx context.Context) *DeleteOrganizationInvitationsParams {
	return &DeleteOrganizationInvitationsParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationInvitationsParamsWithHTTPClient creates a new DeleteOrganizationInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationInvitationsParamsWithHTTPClient(client *http.Client) *DeleteOrganizationInvitationsParams {
	return &DeleteOrganizationInvitationsParams{
		HTTPClient: client,
	}
}

/* DeleteOrganizationInvitationsParams contains all the parameters to send to the API endpoint
   for the delete organization invitations operation.

   Typically these are written to a http.Request.
*/
type DeleteOrganizationInvitationsParams struct {

	/* InvitationTokens.

	   CSV list of Invitation tokens
	*/
	InvitationTokens string

	/* OrganizationID.

	   Identifier for the Organization
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInvitationsParams) WithDefaults() *DeleteOrganizationInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) WithTimeout(timeout time.Duration) *DeleteOrganizationInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) WithContext(ctx context.Context) *DeleteOrganizationInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) WithHTTPClient(client *http.Client) *DeleteOrganizationInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationTokens adds the invitationTokens to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) WithInvitationTokens(invitationTokens string) *DeleteOrganizationInvitationsParams {
	o.SetInvitationTokens(invitationTokens)
	return o
}

// SetInvitationTokens adds the invitationTokens to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) SetInvitationTokens(invitationTokens string) {
	o.InvitationTokens = invitationTokens
}

// WithOrganizationID adds the organizationID to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) WithOrganizationID(organizationID string) *DeleteOrganizationInvitationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization invitations params
func (o *DeleteOrganizationInvitationsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation_tokens
	if err := r.SetPathParam("invitation_tokens", o.InvitationTokens); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
