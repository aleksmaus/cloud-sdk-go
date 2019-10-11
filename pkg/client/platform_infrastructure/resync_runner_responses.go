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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// ResyncRunnerReader is a Reader for the ResyncRunner structure.
type ResyncRunnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncRunnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResyncRunnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 449:
		result := NewResyncRunnerRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewResyncRunnerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResyncRunnerOK creates a ResyncRunnerOK with default headers values
func NewResyncRunnerOK() *ResyncRunnerOK {
	return &ResyncRunnerOK{}
}

/*ResyncRunnerOK handles this case with default header values.

The runner resync operation executed successfully
*/
type ResyncRunnerOK struct {
	Payload models.EmptyResponse
}

func (o *ResyncRunnerOK) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/{runner_id}/_resync][%d] resyncRunnerOK  %+v", 200, o.Payload)
}

func (o *ResyncRunnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncRunnerRetryWith creates a ResyncRunnerRetryWith with default headers values
func NewResyncRunnerRetryWith() *ResyncRunnerRetryWith {
	return &ResyncRunnerRetryWith{}
}

/*ResyncRunnerRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncRunnerRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncRunnerRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/{runner_id}/_resync][%d] resyncRunnerRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncRunnerRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncRunnerInternalServerError creates a ResyncRunnerInternalServerError with default headers values
func NewResyncRunnerInternalServerError() *ResyncRunnerInternalServerError {
	return &ResyncRunnerInternalServerError{}
}

/*ResyncRunnerInternalServerError handles this case with default header values.

The runner resync operation failed for runner {runner_id}. (code: `runners.resync_failed`)
*/
type ResyncRunnerInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncRunnerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/{runner_id}/_resync][%d] resyncRunnerInternalServerError  %+v", 500, o.Payload)
}

func (o *ResyncRunnerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}