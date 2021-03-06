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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// PutKibProxyRequestsReader is a Reader for the PutKibProxyRequests structure.
type PutKibProxyRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutKibProxyRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutKibProxyRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutKibProxyRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutKibProxyRequestsOK creates a PutKibProxyRequestsOK with default headers values
func NewPutKibProxyRequestsOK() *PutKibProxyRequestsOK {
	return &PutKibProxyRequestsOK{}
}

/*PutKibProxyRequestsOK handles this case with default header values.

The request has been processed successfully through the proxy
*/
type PutKibProxyRequestsOK struct {
}

func (o *PutKibProxyRequestsOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/kibana/{cluster_id}/proxy/{kibana_path}][%d] putKibProxyRequestsOK ", 200)
}

func (o *PutKibProxyRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutKibProxyRequestsNotFound creates a PutKibProxyRequestsNotFound with default headers values
func NewPutKibProxyRequestsNotFound() *PutKibProxyRequestsNotFound {
	return &PutKibProxyRequestsNotFound{}
}

/*PutKibProxyRequestsNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type PutKibProxyRequestsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *PutKibProxyRequestsNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/kibana/{cluster_id}/proxy/{kibana_path}][%d] putKibProxyRequestsNotFound  %+v", 404, o.Payload)
}

func (o *PutKibProxyRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
