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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UnclaimTrafficFilterLinkIDReader is a Reader for the UnclaimTrafficFilterLinkID structure.
type UnclaimTrafficFilterLinkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnclaimTrafficFilterLinkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnclaimTrafficFilterLinkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnclaimTrafficFilterLinkIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnclaimTrafficFilterLinkIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnclaimTrafficFilterLinkIDOK creates a UnclaimTrafficFilterLinkIDOK with default headers values
func NewUnclaimTrafficFilterLinkIDOK() *UnclaimTrafficFilterLinkIDOK {
	return &UnclaimTrafficFilterLinkIDOK{}
}

/*
UnclaimTrafficFilterLinkIDOK describes a response with status code 200, with default header values.

The claimed link id was successfully deleted.
*/
type UnclaimTrafficFilterLinkIDOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this unclaim traffic filter link Id o k response has a 2xx status code
func (o *UnclaimTrafficFilterLinkIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unclaim traffic filter link Id o k response has a 3xx status code
func (o *UnclaimTrafficFilterLinkIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unclaim traffic filter link Id o k response has a 4xx status code
func (o *UnclaimTrafficFilterLinkIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unclaim traffic filter link Id o k response has a 5xx status code
func (o *UnclaimTrafficFilterLinkIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unclaim traffic filter link Id o k response a status code equal to that given
func (o *UnclaimTrafficFilterLinkIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unclaim traffic filter link Id o k response
func (o *UnclaimTrafficFilterLinkIDOK) Code() int {
	return 200
}

func (o *UnclaimTrafficFilterLinkIDOK) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdOK  %+v", 200, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDOK) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdOK  %+v", 200, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *UnclaimTrafficFilterLinkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnclaimTrafficFilterLinkIDBadRequest creates a UnclaimTrafficFilterLinkIDBadRequest with default headers values
func NewUnclaimTrafficFilterLinkIDBadRequest() *UnclaimTrafficFilterLinkIDBadRequest {
	return &UnclaimTrafficFilterLinkIDBadRequest{}
}

/*
UnclaimTrafficFilterLinkIDBadRequest describes a response with status code 400, with default header values.

Error validating the request. (code: `traffic_filter_claimed_link_id.invalid_input`)
*/
type UnclaimTrafficFilterLinkIDBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this unclaim traffic filter link Id bad request response has a 2xx status code
func (o *UnclaimTrafficFilterLinkIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unclaim traffic filter link Id bad request response has a 3xx status code
func (o *UnclaimTrafficFilterLinkIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unclaim traffic filter link Id bad request response has a 4xx status code
func (o *UnclaimTrafficFilterLinkIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unclaim traffic filter link Id bad request response has a 5xx status code
func (o *UnclaimTrafficFilterLinkIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unclaim traffic filter link Id bad request response a status code equal to that given
func (o *UnclaimTrafficFilterLinkIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the unclaim traffic filter link Id bad request response
func (o *UnclaimTrafficFilterLinkIDBadRequest) Code() int {
	return 400
}

func (o *UnclaimTrafficFilterLinkIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdBadRequest  %+v", 400, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDBadRequest) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdBadRequest  %+v", 400, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UnclaimTrafficFilterLinkIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUnclaimTrafficFilterLinkIDInternalServerError creates a UnclaimTrafficFilterLinkIDInternalServerError with default headers values
func NewUnclaimTrafficFilterLinkIDInternalServerError() *UnclaimTrafficFilterLinkIDInternalServerError {
	return &UnclaimTrafficFilterLinkIDInternalServerError{}
}

/*
UnclaimTrafficFilterLinkIDInternalServerError describes a response with status code 500, with default header values.

Error deleting the traffic filter claimed link id. (code: `traffic_filter_claimed_link_id.request_execution_failed`)
*/
type UnclaimTrafficFilterLinkIDInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this unclaim traffic filter link Id internal server error response has a 2xx status code
func (o *UnclaimTrafficFilterLinkIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unclaim traffic filter link Id internal server error response has a 3xx status code
func (o *UnclaimTrafficFilterLinkIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unclaim traffic filter link Id internal server error response has a 4xx status code
func (o *UnclaimTrafficFilterLinkIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this unclaim traffic filter link Id internal server error response has a 5xx status code
func (o *UnclaimTrafficFilterLinkIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this unclaim traffic filter link Id internal server error response a status code equal to that given
func (o *UnclaimTrafficFilterLinkIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the unclaim traffic filter link Id internal server error response
func (o *UnclaimTrafficFilterLinkIDInternalServerError) Code() int {
	return 500
}

func (o *UnclaimTrafficFilterLinkIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/link-ids/_unclaim][%d] unclaimTrafficFilterLinkIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UnclaimTrafficFilterLinkIDInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UnclaimTrafficFilterLinkIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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