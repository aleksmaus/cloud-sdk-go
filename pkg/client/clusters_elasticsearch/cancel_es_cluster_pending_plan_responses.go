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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// CancelEsClusterPendingPlanReader is a Reader for the CancelEsClusterPendingPlan structure.
type CancelEsClusterPendingPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelEsClusterPendingPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelEsClusterPendingPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewCancelEsClusterPendingPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewCancelEsClusterPendingPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewCancelEsClusterPendingPlanRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelEsClusterPendingPlanOK creates a CancelEsClusterPendingPlanOK with default headers values
func NewCancelEsClusterPendingPlanOK() *CancelEsClusterPendingPlanOK {
	return &CancelEsClusterPendingPlanOK{}
}

/*CancelEsClusterPendingPlanOK handles this case with default header values.

The pending plan has been successfully cancelled
*/
type CancelEsClusterPendingPlanOK struct {
	Payload models.EmptyResponse
}

func (o *CancelEsClusterPendingPlanOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/elasticsearch/{cluster_id}/plan/pending][%d] cancelEsClusterPendingPlanOK  %+v", 200, o.Payload)
}

func (o *CancelEsClusterPendingPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEsClusterPendingPlanNotFound creates a CancelEsClusterPendingPlanNotFound with default headers values
func NewCancelEsClusterPendingPlanNotFound() *CancelEsClusterPendingPlanNotFound {
	return &CancelEsClusterPendingPlanNotFound{}
}

/*CancelEsClusterPendingPlanNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type CancelEsClusterPendingPlanNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CancelEsClusterPendingPlanNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/elasticsearch/{cluster_id}/plan/pending][%d] cancelEsClusterPendingPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelEsClusterPendingPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEsClusterPendingPlanPreconditionFailed creates a CancelEsClusterPendingPlanPreconditionFailed with default headers values
func NewCancelEsClusterPendingPlanPreconditionFailed() *CancelEsClusterPendingPlanPreconditionFailed {
	return &CancelEsClusterPendingPlanPreconditionFailed{}
}

/*CancelEsClusterPendingPlanPreconditionFailed handles this case with default header values.

There is not currently applied plan - eg the cluster has not finished provisioning, or the provisioning failed. (code: `clusters.cluster_plan_state_error`)
*/
type CancelEsClusterPendingPlanPreconditionFailed struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CancelEsClusterPendingPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /clusters/elasticsearch/{cluster_id}/plan/pending][%d] cancelEsClusterPendingPlanPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CancelEsClusterPendingPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEsClusterPendingPlanRetryWith creates a CancelEsClusterPendingPlanRetryWith with default headers values
func NewCancelEsClusterPendingPlanRetryWith() *CancelEsClusterPendingPlanRetryWith {
	return &CancelEsClusterPendingPlanRetryWith{}
}

/*CancelEsClusterPendingPlanRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CancelEsClusterPendingPlanRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CancelEsClusterPendingPlanRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /clusters/elasticsearch/{cluster_id}/plan/pending][%d] cancelEsClusterPendingPlanRetryWith  %+v", 449, o.Payload)
}

func (o *CancelEsClusterPendingPlanRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}