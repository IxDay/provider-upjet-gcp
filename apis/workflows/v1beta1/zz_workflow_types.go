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

type WorkflowObservation struct {

	// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/workflows/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the workflow.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The revision of the workflow. A new one is generated if the service account or source contents is changed.
	RevisionID *string `json:"revisionId,omitempty" tf:"revision_id,omitempty"`

	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the unique_id of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Workflow code to be executed. The size limit is 32KB.
	SourceContents *string `json:"sourceContents,omitempty" tf:"source_contents,omitempty"`

	// State of the workflow deployment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type WorkflowParameters struct {

	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	// +kubebuilder:validation:Optional
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the workflow.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the unique_id of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccount.
	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccount.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// Workflow code to be executed. The size limit is 32KB.
	// +kubebuilder:validation:Optional
	SourceContents *string `json:"sourceContents,omitempty" tf:"source_contents,omitempty"`
}

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkflowParameters `json:"forProvider"`
}

// WorkflowStatus defines the observed state of Workflow.
type WorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workflow is the Schema for the Workflows API. Workflow program to be executed by Workflows.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkflowSpec   `json:"spec"`
	Status            WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

// Repository type metadata.
var (
	Workflow_Kind             = "Workflow"
	Workflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workflow_Kind}.String()
	Workflow_KindAPIVersion   = Workflow_Kind + "." + CRDGroupVersion.String()
	Workflow_GroupVersionKind = CRDGroupVersion.WithKind(Workflow_Kind)
)

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
