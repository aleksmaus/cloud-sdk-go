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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEsClusterCcsSettingsParams creates a new GetEsClusterCcsSettingsParams object
// with the default values initialized.
func NewGetEsClusterCcsSettingsParams() *GetEsClusterCcsSettingsParams {
	var ()
	return &GetEsClusterCcsSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEsClusterCcsSettingsParamsWithTimeout creates a new GetEsClusterCcsSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEsClusterCcsSettingsParamsWithTimeout(timeout time.Duration) *GetEsClusterCcsSettingsParams {
	var ()
	return &GetEsClusterCcsSettingsParams{

		timeout: timeout,
	}
}

// NewGetEsClusterCcsSettingsParamsWithContext creates a new GetEsClusterCcsSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEsClusterCcsSettingsParamsWithContext(ctx context.Context) *GetEsClusterCcsSettingsParams {
	var ()
	return &GetEsClusterCcsSettingsParams{

		Context: ctx,
	}
}

// NewGetEsClusterCcsSettingsParamsWithHTTPClient creates a new GetEsClusterCcsSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEsClusterCcsSettingsParamsWithHTTPClient(client *http.Client) *GetEsClusterCcsSettingsParams {
	var ()
	return &GetEsClusterCcsSettingsParams{
		HTTPClient: client,
	}
}

/*GetEsClusterCcsSettingsParams contains all the parameters to send to the API endpoint
for the get es cluster ccs settings operation typically these are written to a http.Request
*/
type GetEsClusterCcsSettingsParams struct {

	/*ClusterID
	  Elasticsearch cluster identifier

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) WithTimeout(timeout time.Duration) *GetEsClusterCcsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) WithContext(ctx context.Context) *GetEsClusterCcsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) WithHTTPClient(client *http.Client) *GetEsClusterCcsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) WithClusterID(clusterID string) *GetEsClusterCcsSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get es cluster ccs settings params
func (o *GetEsClusterCcsSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEsClusterCcsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}