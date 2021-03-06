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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewMoveKibanaClusterInstancesParams creates a new MoveKibanaClusterInstancesParams object
// with the default values initialized.
func NewMoveKibanaClusterInstancesParams() *MoveKibanaClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveKibanaClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		ValidateOnly:  &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveKibanaClusterInstancesParamsWithTimeout creates a new MoveKibanaClusterInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveKibanaClusterInstancesParamsWithTimeout(timeout time.Duration) *MoveKibanaClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveKibanaClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		ValidateOnly:  &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewMoveKibanaClusterInstancesParamsWithContext creates a new MoveKibanaClusterInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveKibanaClusterInstancesParamsWithContext(ctx context.Context) *MoveKibanaClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveKibanaClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		ValidateOnly:  &validateOnlyDefault,

		Context: ctx,
	}
}

// NewMoveKibanaClusterInstancesParamsWithHTTPClient creates a new MoveKibanaClusterInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveKibanaClusterInstancesParamsWithHTTPClient(client *http.Client) *MoveKibanaClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveKibanaClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		ValidateOnly:  &validateOnlyDefault,
		HTTPClient:    client,
	}
}

/*MoveKibanaClusterInstancesParams contains all the parameters to send to the API endpoint
for the move kibana cluster instances operation typically these are written to a http.Request
*/
type MoveKibanaClusterInstancesParams struct {

	/*Body
	  Overrides defaults for the move, including setting the configuration of instances specified in the path

	*/
	Body *models.TransientKibanaPlanConfiguration
	/*ClusterID
	  The Kibana deployment identifier.

	*/
	ClusterID string
	/*ForceUpdate
	  When `true`, cancels and overwrites the pending plans, or treats the instance as an error.

	*/
	ForceUpdate *bool
	/*IgnoreMissing
	  When `true` and the instance does not exist, proceeds to the next instance, or treats the instance as an error.

	*/
	IgnoreMissing *bool
	/*InstanceIds
	  A comma-separated list of instance identifiers.

	*/
	InstanceIds []string
	/*ValidateOnly
	  When `true`, validates the move request, then returns the calculated plan without applying the plan.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithTimeout(timeout time.Duration) *MoveKibanaClusterInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithContext(ctx context.Context) *MoveKibanaClusterInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithHTTPClient(client *http.Client) *MoveKibanaClusterInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithBody(body *models.TransientKibanaPlanConfiguration) *MoveKibanaClusterInstancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetBody(body *models.TransientKibanaPlanConfiguration) {
	o.Body = body
}

// WithClusterID adds the clusterID to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithClusterID(clusterID string) *MoveKibanaClusterInstancesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceUpdate adds the forceUpdate to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithForceUpdate(forceUpdate *bool) *MoveKibanaClusterInstancesParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithIgnoreMissing adds the ignoreMissing to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithIgnoreMissing(ignoreMissing *bool) *MoveKibanaClusterInstancesParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithInstanceIds(instanceIds []string) *MoveKibanaClusterInstancesParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WithValidateOnly adds the validateOnly to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) WithValidateOnly(validateOnly *bool) *MoveKibanaClusterInstancesParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move kibana cluster instances params
func (o *MoveKibanaClusterInstancesParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveKibanaClusterInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool
		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {
			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
		}

	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	valuesInstanceIds := o.InstanceIds

	joinedInstanceIds := swag.JoinByFormat(valuesInstanceIds, "csv")
	// path array param instance_ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedInstanceIds) > 0 {
		if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
			return err
		}
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
