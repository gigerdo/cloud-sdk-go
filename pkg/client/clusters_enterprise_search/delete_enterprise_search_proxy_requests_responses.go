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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteEnterpriseSearchProxyRequestsReader is a Reader for the DeleteEnterpriseSearchProxyRequests structure.
type DeleteEnterpriseSearchProxyRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnterpriseSearchProxyRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnterpriseSearchProxyRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteEnterpriseSearchProxyRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEnterpriseSearchProxyRequestsOK creates a DeleteEnterpriseSearchProxyRequestsOK with default headers values
func NewDeleteEnterpriseSearchProxyRequestsOK() *DeleteEnterpriseSearchProxyRequestsOK {
	return &DeleteEnterpriseSearchProxyRequestsOK{}
}

/* DeleteEnterpriseSearchProxyRequestsOK describes a response with status code 200, with default header values.

The request has been processed successfully through the proxy
*/
type DeleteEnterpriseSearchProxyRequestsOK struct {
}

func (o *DeleteEnterpriseSearchProxyRequestsOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/enterprise_search/{cluster_id}/proxy/{enterprise_search_path}][%d] deleteEnterpriseSearchProxyRequestsOK ", 200)
}

func (o *DeleteEnterpriseSearchProxyRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnterpriseSearchProxyRequestsNotFound creates a DeleteEnterpriseSearchProxyRequestsNotFound with default headers values
func NewDeleteEnterpriseSearchProxyRequestsNotFound() *DeleteEnterpriseSearchProxyRequestsNotFound {
	return &DeleteEnterpriseSearchProxyRequestsNotFound{}
}

/* DeleteEnterpriseSearchProxyRequestsNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type DeleteEnterpriseSearchProxyRequestsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteEnterpriseSearchProxyRequestsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/enterprise_search/{cluster_id}/proxy/{enterprise_search_path}][%d] deleteEnterpriseSearchProxyRequestsNotFound  %+v", 404, o.Payload)
}
func (o *DeleteEnterpriseSearchProxyRequestsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteEnterpriseSearchProxyRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}