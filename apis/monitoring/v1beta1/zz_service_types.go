/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BasicServiceObservation struct {

	// Labels that specify the resource that emits the monitoring data
	// which is used for SLO reporting of this Service.
	ServiceLabels map[string]*string `json:"serviceLabels,omitempty" tf:"service_labels,omitempty"`

	// The type of service that this basic service defines, e.g.
	// APP_ENGINE service type
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`
}

type BasicServiceParameters struct {

	// Labels that specify the resource that emits the monitoring data
	// which is used for SLO reporting of this Service.
	// +kubebuilder:validation:Optional
	ServiceLabels map[string]*string `json:"serviceLabels,omitempty" tf:"service_labels,omitempty"`

	// The type of service that this basic service defines, e.g.
	// APP_ENGINE service type
	// +kubebuilder:validation:Optional
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`
}

type ServiceObservation struct {

	// A well-known service type, defined by its service type and service labels.
	// Valid values of service types and services labels are described at
	// https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
	// Structure is documented below.
	BasicService []BasicServiceObservation `json:"basicService,omitempty" tf:"basic_service,omitempty"`

	// Name used for UI elements listing this Service.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/services/{{service_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry []ServiceTelemetryObservation `json:"telemetry,omitempty" tf:"telemetry,omitempty"`

	// Labels which have been used to annotate the service. Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type ServiceParameters struct {

	// A well-known service type, defined by its service type and service labels.
	// Valid values of service types and services labels are described at
	// https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BasicService []BasicServiceParameters `json:"basicService,omitempty" tf:"basic_service,omitempty"`

	// Name used for UI elements listing this Service.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Labels which have been used to annotate the service. Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	// +kubebuilder:validation:Optional
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type ServiceTelemetryObservation struct {

	// The full name of the resource that defines this service.
	// Formatted as described in
	// https://cloud.google.com/apis/design/resource_names.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

type ServiceTelemetryParameters struct {
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. A Service is a discrete, autonomous, and network-accessible unit, designed to solve an individual concern (Wikipedia).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
