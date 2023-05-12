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

type ActionObservation struct {

	// A list of URLs of the IP resources used for this NAT rule.
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	SourceNATActiveIps []*string `json:"sourceNatActiveIps,omitempty" tf:"source_nat_active_ips,omitempty"`

	// A list of URLs of the IP resources to be drained.
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	SourceNATDrainIps []*string `json:"sourceNatDrainIps,omitempty" tf:"source_nat_drain_ips,omitempty"`
}

type ActionParameters struct {

	// A list of URLs of the IP resources used for this NAT rule.
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	// +kubebuilder:validation:Optional
	SourceNATActiveIps []*string `json:"sourceNatActiveIps,omitempty" tf:"source_nat_active_ips,omitempty"`

	// A list of URLs of the IP resources to be drained.
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	// +kubebuilder:validation:Optional
	SourceNATDrainIps []*string `json:"sourceNatDrainIps,omitempty" tf:"source_nat_drain_ips,omitempty"`
}

type RouterNATLogConfigObservation struct {

	// Indicates whether or not to export logs.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the desired filtering of logs on this NAT.
	// Possible values are: ERRORS_ONLY, TRANSLATIONS_ONLY, ALL.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type RouterNATLogConfigParameters struct {

	// Indicates whether or not to export logs.
	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// Specifies the desired filtering of logs on this NAT.
	// Possible values are: ERRORS_ONLY, TRANSLATIONS_ONLY, ALL.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`
}

type RouterNATObservation struct {

	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNATIps []*string `json:"drainNatIps,omitempty" tf:"drain_nat_ips,omitempty"`

	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty" tf:"enable_dynamic_port_allocation,omitempty"`

	// Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
	// see the official documentation.
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty" tf:"enable_endpoint_independent_mapping,omitempty"`

	// an identifier for the resource with format {{project}}/{{region}}/{{router}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec *float64 `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`

	// Configuration for logging on NAT
	// Structure is documented below.
	LogConfig []RouterNATLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	MaxPortsPerVM *float64 `json:"maxPortsPerVm,omitempty" tf:"max_ports_per_vm,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVM *float64 `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`

	// How external IPs should be allocated for this NAT. Valid values are
	// AUTO_ONLY for only allowing NAT IPs allocated by Google Cloud
	// Platform, or MANUAL_ONLY for only user-allocated NAT IP addresses.
	// Possible values are: MANUAL_ONLY, AUTO_ONLY.
	NATIPAllocateOption *string `json:"natIpAllocateOption,omitempty" tf:"nat_ip_allocate_option,omitempty"`

	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NATIps []*string `json:"natIps,omitempty" tf:"nat_ips,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and NAT reside.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the Cloud Router in which this NAT will be configured.
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// A list of rules associated with this NAT.
	// Structure is documented below.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// How NAT should be configured per Subnetwork.
	// If ALL_SUBNETWORKS_ALL_IP_RANGES, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	// Possible values are: ALL_SUBNETWORKS_ALL_IP_RANGES, ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, LIST_OF_SUBNETWORKS.
	SourceSubnetworkIPRangesToNAT *string `json:"sourceSubnetworkIpRangesToNat,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`

	// One or more subnetwork NAT configurations. Only used if
	// source_subnetwork_ip_ranges_to_nat is set to LIST_OF_SUBNETWORKS
	// Structure is documented below.
	Subnetwork []SubnetworkObservation `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TCPEstablishedIdleTimeoutSec *float64 `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	TCPTimeWaitTimeoutSec *float64 `json:"tcpTimeWaitTimeoutSec,omitempty" tf:"tcp_time_wait_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TCPTransitoryIdleTimeoutSec *float64 `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UDPIdleTimeoutSec *float64 `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type RouterNATParameters struct {

	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	// +kubebuilder:validation:Optional
	DrainNATIps []*string `json:"drainNatIps,omitempty" tf:"drain_nat_ips,omitempty"`

	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	// +kubebuilder:validation:Optional
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty" tf:"enable_dynamic_port_allocation,omitempty"`

	// Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
	// see the official documentation.
	// +kubebuilder:validation:Optional
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty" tf:"enable_endpoint_independent_mapping,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	IcmpIdleTimeoutSec *float64 `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`

	// Configuration for logging on NAT
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []RouterNATLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	// +kubebuilder:validation:Optional
	MaxPortsPerVM *float64 `json:"maxPortsPerVm,omitempty" tf:"max_ports_per_vm,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT.
	// +kubebuilder:validation:Optional
	MinPortsPerVM *float64 `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`

	// How external IPs should be allocated for this NAT. Valid values are
	// AUTO_ONLY for only allowing NAT IPs allocated by Google Cloud
	// Platform, or MANUAL_ONLY for only user-allocated NAT IP addresses.
	// Possible values are: MANUAL_ONLY, AUTO_ONLY.
	// +kubebuilder:validation:Optional
	NATIPAllocateOption *string `json:"natIpAllocateOption,omitempty" tf:"nat_ip_allocate_option,omitempty"`

	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	// +kubebuilder:validation:Optional
	NATIps []*string `json:"natIps,omitempty" tf:"nat_ips,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and NAT reside.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The name of the Cloud Router in which this NAT will be configured.
	// +crossplane:generate:reference:type=Router
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// A list of rules associated with this NAT.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// How NAT should be configured per Subnetwork.
	// If ALL_SUBNETWORKS_ALL_IP_RANGES, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	// Possible values are: ALL_SUBNETWORKS_ALL_IP_RANGES, ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, LIST_OF_SUBNETWORKS.
	// +kubebuilder:validation:Optional
	SourceSubnetworkIPRangesToNAT *string `json:"sourceSubnetworkIpRangesToNat,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`

	// One or more subnetwork NAT configurations. Only used if
	// source_subnetwork_ip_ranges_to_nat is set to LIST_OF_SUBNETWORKS
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Subnetwork []SubnetworkParameters `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	// +kubebuilder:validation:Optional
	TCPEstablishedIdleTimeoutSec *float64 `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	// +kubebuilder:validation:Optional
	TCPTimeWaitTimeoutSec *float64 `json:"tcpTimeWaitTimeoutSec,omitempty" tf:"tcp_time_wait_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	TCPTransitoryIdleTimeoutSec *float64 `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	UDPIdleTimeoutSec *float64 `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type RulesObservation struct {

	// The action to be enforced for traffic that matches this rule.
	// Structure is documented below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	// If it evaluates to true, the corresponding action is enforced.
	// The following examples are valid match expressions for public NAT:
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	// The following example is a valid match expression for private NAT:
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// An integer uniquely identifying a rule in the list.
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`
}

type RulesParameters struct {

	// The action to be enforced for traffic that matches this rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	// If it evaluates to true, the corresponding action is enforced.
	// The following examples are valid match expressions for public NAT:
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	// The following example is a valid match expression for private NAT:
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	// +kubebuilder:validation:Required
	Match *string `json:"match" tf:"match,omitempty"`

	// An integer uniquely identifying a rule in the list.
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	// +kubebuilder:validation:Required
	RuleNumber *float64 `json:"ruleNumber" tf:"rule_number,omitempty"`
}

type SubnetworkObservation struct {

	// Self-link of subnetwork to NAT
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of the secondary ranges of the subnetwork that are allowed
	// to use NAT. This can be populated only if
	// LIST_OF_SECONDARY_IP_RANGES is one of the values in
	// sourceIpRangesToNat
	SecondaryIPRangeNames []*string `json:"secondaryIpRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`

	// List of options for which source IPs in the subnetwork
	// should have NAT enabled. Supported values include:
	// ALL_IP_RANGES, LIST_OF_SECONDARY_IP_RANGES,
	// PRIMARY_IP_RANGE.
	SourceIPRangesToNAT []*string `json:"sourceIpRangesToNat,omitempty" tf:"source_ip_ranges_to_nat,omitempty"`
}

type SubnetworkParameters struct {

	// Self-link of subnetwork to NAT
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Subnetwork to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// List of the secondary ranges of the subnetwork that are allowed
	// to use NAT. This can be populated only if
	// LIST_OF_SECONDARY_IP_RANGES is one of the values in
	// sourceIpRangesToNat
	// +kubebuilder:validation:Optional
	SecondaryIPRangeNames []*string `json:"secondaryIpRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`

	// List of options for which source IPs in the subnetwork
	// should have NAT enabled. Supported values include:
	// ALL_IP_RANGES, LIST_OF_SECONDARY_IP_RANGES,
	// PRIMARY_IP_RANGE.
	// +kubebuilder:validation:Required
	SourceIPRangesToNAT []*string `json:"sourceIpRangesToNat" tf:"source_ip_ranges_to_nat,omitempty"`
}

// RouterNATSpec defines the desired state of RouterNAT
type RouterNATSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterNATParameters `json:"forProvider"`
}

// RouterNATStatus defines the observed state of RouterNAT.
type RouterNATStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterNATObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterNAT is the Schema for the RouterNATs API. A NAT service created in a router.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.natIpAllocateOption)",message="natIpAllocateOption is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sourceSubnetworkIpRangesToNat)",message="sourceSubnetworkIpRangesToNat is a required parameter"
	Spec   RouterNATSpec   `json:"spec"`
	Status RouterNATStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterNATList contains a list of RouterNATs
type RouterNATList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterNAT `json:"items"`
}

// Repository type metadata.
var (
	RouterNAT_Kind             = "RouterNAT"
	RouterNAT_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterNAT_Kind}.String()
	RouterNAT_KindAPIVersion   = RouterNAT_Kind + "." + CRDGroupVersion.String()
	RouterNAT_GroupVersionKind = CRDGroupVersion.WithKind(RouterNAT_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterNAT{}, &RouterNATList{})
}
