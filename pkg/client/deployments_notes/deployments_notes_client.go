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

package deployments_notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployments notes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments notes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeploymentNote(params *CreateDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeploymentNoteCreated, error)

	DeleteDeploymentNote(params *DeleteDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeploymentNoteOK, error)

	GetDeploymentNote(params *GetDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentNoteOK, error)

	GetDeploymentNotes(params *GetDeploymentNotesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentNotesOK, error)

	UpdateDeploymentNote(params *UpdateDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeploymentNoteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDeploymentNote creates deployment note

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version. Use [Create a Comment](#Comments) instead.

Create note for the running deployment.
*/
func (a *Client) CreateDeploymentNote(params *CreateDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeploymentNoteCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentNoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-deployment-note",
		Method:             "POST",
		PathPattern:        "/deployments/{deployment_id}/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentNoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeploymentNoteCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-deployment-note: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDeploymentNote deletes deployment note

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version. Use [Delete Comment](#Comments) instead.

Delete note for the running deployment.
*/
func (a *Client) DeleteDeploymentNote(params *DeleteDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeploymentNoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentNoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-deployment-note",
		Method:             "DELETE",
		PathPattern:        "/deployments/{deployment_id}/notes/{note_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeploymentNoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeploymentNoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-deployment-note: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentNote gets a deployment note

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version. Use [Get Comment](#Comments) instead.

Gets a deployment note.
*/
func (a *Client) GetDeploymentNote(params *GetDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentNoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentNoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-deployment-note",
		Method:             "GET",
		PathPattern:        "/deployments/{deployment_id}/notes/{note_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentNoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentNoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-deployment-note: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentNotes gets deployment notes

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version. Use [List Comments](#Comments) instead.

Get deployment notes.
*/
func (a *Client) GetDeploymentNotes(params *GetDeploymentNotesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentNotesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentNotesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-deployment-notes",
		Method:             "GET",
		PathPattern:        "/deployments/{deployment_id}/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentNotesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentNotesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-deployment-notes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDeploymentNote updates deployment note

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version. Use [Update Comment](#Comments) instead.

Update note for the running deployment.
*/
func (a *Client) UpdateDeploymentNote(params *UpdateDeploymentNoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeploymentNoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentNoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-deployment-note",
		Method:             "PUT",
		PathPattern:        "/deployments/{deployment_id}/notes/{note_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeploymentNoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeploymentNoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-deployment-note: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
