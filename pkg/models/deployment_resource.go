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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentResource Data for a deployment resource
// swagger:model DeploymentResource
type DeploymentResource struct {

	// An encoded string that provides other Elastic services with the necessary information to connect to this Elasticsearch and Kibana
	CloudID string `json:"cloud_id,omitempty"`

	// Credentials for logging into a created resource. Only provided on initial create and absent otherwise.
	Credentials *ClusterCredentials `json:"credentials,omitempty"`

	// The Elasticsearch cluster that this resource depends on.
	ElasticsearchClusterRefID string `json:"elasticsearch_cluster_ref_id,omitempty"`

	// A system-unique id for the created resource
	// Required: true
	ID *string `json:"id"`

	// The kind of resource
	// Required: true
	Kind *string `json:"kind"`

	// A locally-unique friendly alias for this Elasticsearch cluster
	// Required: true
	RefID *string `json:"ref_id"`

	// Identifier of the region in which this resource runs.
	// Required: true
	Region *string `json:"region"`

	// Secret token for using a created resource. Only provided on initial create and absent otherwise.
	SecretToken string `json:"secret_token,omitempty"`
}

// Validate validates this deployment resource
func (m *DeploymentResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentResource) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentResource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateRefID(formats strfmt.Registry) error {

	if err := validate.Required("ref_id", "body", m.RefID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentResource) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentResource) UnmarshalBinary(b []byte) error {
	var res DeploymentResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}