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

package clusters_enterprise_search

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

// NewPutEnterpriseSearchProxyRequestsParams creates a new PutEnterpriseSearchProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEnterpriseSearchProxyRequestsParams() *PutEnterpriseSearchProxyRequestsParams {
	return &PutEnterpriseSearchProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEnterpriseSearchProxyRequestsParamsWithTimeout creates a new PutEnterpriseSearchProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewPutEnterpriseSearchProxyRequestsParamsWithTimeout(timeout time.Duration) *PutEnterpriseSearchProxyRequestsParams {
	return &PutEnterpriseSearchProxyRequestsParams{
		timeout: timeout,
	}
}

// NewPutEnterpriseSearchProxyRequestsParamsWithContext creates a new PutEnterpriseSearchProxyRequestsParams object
// with the ability to set a context for a request.
func NewPutEnterpriseSearchProxyRequestsParamsWithContext(ctx context.Context) *PutEnterpriseSearchProxyRequestsParams {
	return &PutEnterpriseSearchProxyRequestsParams{
		Context: ctx,
	}
}

// NewPutEnterpriseSearchProxyRequestsParamsWithHTTPClient creates a new PutEnterpriseSearchProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEnterpriseSearchProxyRequestsParamsWithHTTPClient(client *http.Client) *PutEnterpriseSearchProxyRequestsParams {
	return &PutEnterpriseSearchProxyRequestsParams{
		HTTPClient: client,
	}
}

/* PutEnterpriseSearchProxyRequestsParams contains all the parameters to send to the API endpoint
   for the put enterprise search proxy requests operation.

   Typically these are written to a http.Request.
*/
type PutEnterpriseSearchProxyRequestsParams struct {

	/* XManagementRequest.

	   When set to `true`, includes the X-Management-Request header value.
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload that is used to proxy the EnterpriseSearch deployment.
	*/
	Body string

	/* ClusterID.

	   The EnterpriseSearch deployment identifier
	*/
	ClusterID string

	/* EnterpriseSearchPath.

	   The URL part to proxy to the EnterpriseSearch cluster. Example: /api/ent/v1/internal/read_only_mode or /api/ent/v1/internal/health
	*/
	EnterpriseSearchPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEnterpriseSearchProxyRequestsParams) WithDefaults() *PutEnterpriseSearchProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEnterpriseSearchProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithTimeout(timeout time.Duration) *PutEnterpriseSearchProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithContext(ctx context.Context) *PutEnterpriseSearchProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithHTTPClient(client *http.Client) *PutEnterpriseSearchProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *PutEnterpriseSearchProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithBody(body string) *PutEnterpriseSearchProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithClusterID(clusterID string) *PutEnterpriseSearchProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEnterpriseSearchPath adds the enterpriseSearchPath to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) WithEnterpriseSearchPath(enterpriseSearchPath string) *PutEnterpriseSearchProxyRequestsParams {
	o.SetEnterpriseSearchPath(enterpriseSearchPath)
	return o
}

// SetEnterpriseSearchPath adds the enterpriseSearchPath to the put enterprise search proxy requests params
func (o *PutEnterpriseSearchProxyRequestsParams) SetEnterpriseSearchPath(enterpriseSearchPath string) {
	o.EnterpriseSearchPath = enterpriseSearchPath
}

// WriteToRequest writes these params to a swagger request
func (o *PutEnterpriseSearchProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param enterprise_search_path
	if err := r.SetPathParam("enterprise_search_path", o.EnterpriseSearchPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}