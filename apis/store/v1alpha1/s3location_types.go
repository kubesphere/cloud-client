/*
Copyright 2021 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type S3Credential struct {
	AccessKeyID string `json:"accessKeyID,omitempty"`

	SecretAccessKey string `json:"secretAccessKey,omitempty"`
}

type S3LocationSpec struct {
	Endpoint string `json:"endpoint,omitempty"`

	Credential S3Credential `json:"credential,omitempty"`

	Region string `json:"region,omitempty"`

	S3ForcePathStyle bool `json:"s3ForcePathStyle,omitempty"`

	DisableSSL bool `json:"disableSSL,omitempty"`

	Provider string `json:"provider,omitempty"`

	Bucket string `json:"bucket,omitempty"`

	Default bool `json:"default,omitempty"`
}

type S3LocationConditionType string

const (
	S3LocationConditionTypeAvailable S3LocationConditionType = "Available"
)

type S3LocationCondition struct {
	// Type of the condition
	Type S3LocationConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

type S3LocationStatus struct {
	// Represents the latest available observations of a s3Location's current state.
	Conditions []S3LocationCondition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// S3Location is the schema for the s3Locations API
type S3Location struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3LocationSpec   `json:"spec,omitempty"`
	Status S3LocationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type S3LocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Location `json:"items"`
}
