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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteOrganizationMembershipsReader is a Reader for the DeleteOrganizationMemberships structure.
type DeleteOrganizationMembershipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationMembershipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationMembershipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteOrganizationMembershipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationMembershipsOK creates a DeleteOrganizationMembershipsOK with default headers values
func NewDeleteOrganizationMembershipsOK() *DeleteOrganizationMembershipsOK {
	return &DeleteOrganizationMembershipsOK{}
}

/*
DeleteOrganizationMembershipsOK describes a response with status code 200, with default header values.

Organization membership deleted successfully
*/
type DeleteOrganizationMembershipsOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete organization memberships o k response has a 2xx status code
func (o *DeleteOrganizationMembershipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization memberships o k response has a 3xx status code
func (o *DeleteOrganizationMembershipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization memberships o k response has a 4xx status code
func (o *DeleteOrganizationMembershipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization memberships o k response has a 5xx status code
func (o *DeleteOrganizationMembershipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization memberships o k response a status code equal to that given
func (o *DeleteOrganizationMembershipsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete organization memberships o k response
func (o *DeleteOrganizationMembershipsOK) Code() int {
	return 200
}

func (o *DeleteOrganizationMembershipsOK) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/members/{user_ids}][%d] deleteOrganizationMembershipsOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationMembershipsOK) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/members/{user_ids}][%d] deleteOrganizationMembershipsOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationMembershipsOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteOrganizationMembershipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMembershipsNotFound creates a DeleteOrganizationMembershipsNotFound with default headers values
func NewDeleteOrganizationMembershipsNotFound() *DeleteOrganizationMembershipsNotFound {
	return &DeleteOrganizationMembershipsNotFound{}
}

/*
	DeleteOrganizationMembershipsNotFound describes a response with status code 404, with default header values.

	* User not found. (code: `user.not_found`)

* Organization not found. (code: `organization.not_found`)
* Organization membership not found. (code: `organization.membership_not_found`)
*/
type DeleteOrganizationMembershipsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete organization memberships not found response has a 2xx status code
func (o *DeleteOrganizationMembershipsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization memberships not found response has a 3xx status code
func (o *DeleteOrganizationMembershipsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization memberships not found response has a 4xx status code
func (o *DeleteOrganizationMembershipsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization memberships not found response has a 5xx status code
func (o *DeleteOrganizationMembershipsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization memberships not found response a status code equal to that given
func (o *DeleteOrganizationMembershipsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete organization memberships not found response
func (o *DeleteOrganizationMembershipsNotFound) Code() int {
	return 404
}

func (o *DeleteOrganizationMembershipsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/members/{user_ids}][%d] deleteOrganizationMembershipsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationMembershipsNotFound) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/members/{user_ids}][%d] deleteOrganizationMembershipsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationMembershipsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteOrganizationMembershipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}