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

type EventHandlersObservation struct {

	// The name of the event to handle.
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// (Output)
	// The unique identifier of this event handler.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	TriggerFulfillment []TriggerFulfillmentObservation `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type EventHandlersParameters struct {

	// The name of the event to handle.
	// +kubebuilder:validation:Optional
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []TriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type FlowObservation struct {

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the flow.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A flow's event handlers serve two purposes:
	// They are responsible for handling events (e.g. no match, webhook errors) in the flow.
	// They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
	// Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	// Structure is documented below.
	EventHandlers []EventHandlersObservation `json:"eventHandlers,omitempty" tf:"event_handlers,omitempty"`

	// an identifier for the resource with format {{parent}}/flows/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The language of the following fields in flow:
	// Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases
	// Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The unique identifier of the flow.
	// Format: projects//locations//agents//flows/.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NLU related settings of the flow.
	// Structure is documented below.
	NluSettings []NluSettingsObservation `json:"nluSettings,omitempty" tf:"nlu_settings,omitempty"`

	// The agent to create a flow for.
	// Format: projects//locations//agents/.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// A flow's transition route group serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
	// Format:projects//locations//agents//flows//transitionRouteGroups/.
	TransitionRouteGroups []*string `json:"transitionRouteGroups,omitempty" tf:"transition_route_groups,omitempty"`

	// A flow's transition routes serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
	// TransitionRoutes are evalauted in the following order:
	// TransitionRoutes with intent specified.
	// TransitionRoutes with only condition specified.
	// TransitionRoutes with intent specified are inherited by pages in the flow.
	// Structure is documented below.
	TransitionRoutes []TransitionRoutesObservation `json:"transitionRoutes,omitempty" tf:"transition_routes,omitempty"`
}

type FlowParameters struct {

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the flow.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A flow's event handlers serve two purposes:
	// They are responsible for handling events (e.g. no match, webhook errors) in the flow.
	// They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.
	// Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventHandlers []EventHandlersParameters `json:"eventHandlers,omitempty" tf:"event_handlers,omitempty"`

	// The language of the following fields in flow:
	// Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases
	// Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// NLU related settings of the flow.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NluSettings []NluSettingsParameters `json:"nluSettings,omitempty" tf:"nlu_settings,omitempty"`

	// The agent to create a flow for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// A flow's transition route group serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.
	// Format:projects//locations//agents//flows//transitionRouteGroups/.
	// +kubebuilder:validation:Optional
	TransitionRouteGroups []*string `json:"transitionRouteGroups,omitempty" tf:"transition_route_groups,omitempty"`

	// A flow's transition routes serve two purposes:
	// They are responsible for matching the user's first utterances in the flow.
	// They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.
	// TransitionRoutes are evalauted in the following order:
	// TransitionRoutes with intent specified.
	// TransitionRoutes with only condition specified.
	// TransitionRoutes with intent specified are inherited by pages in the flow.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TransitionRoutes []TransitionRoutesParameters `json:"transitionRoutes,omitempty" tf:"transition_routes,omitempty"`
}

type MessagesObservation struct {

	// A collection of text responses.
	Text []TextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type MessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []TextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type MessagesTextObservation struct {

	// (Output)
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`

	// A collection of text responses.
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type MessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type NluSettingsObservation struct {

	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
	// If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty" tf:"classification_threshold,omitempty"`

	// Indicates NLU model training mode.
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty" tf:"model_training_mode,omitempty"`

	// Indicates the type of NLU model.
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`
}

type NluSettingsParameters struct {

	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
	// If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	// +kubebuilder:validation:Optional
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty" tf:"classification_threshold,omitempty"`

	// Indicates NLU model training mode.
	// +kubebuilder:validation:Optional
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty" tf:"model_training_mode,omitempty"`

	// Indicates the type of NLU model.
	// +kubebuilder:validation:Optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`
}

type TextObservation struct {

	// (Output)
	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`

	// A collection of text responses.
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type TextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type TransitionRoutesObservation struct {

	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects//locations//agents//intents/. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// (Output)
	// The unique identifier of this transition route.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.
	// Structure is documented below.
	TriggerFulfillment []TransitionRoutesTriggerFulfillmentObservation `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type TransitionRoutesParameters struct {

	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects//locations//agents//intents/. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []TransitionRoutesTriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type TransitionRoutesTriggerFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	Messages []TriggerFulfillmentMessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type TransitionRoutesTriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []TriggerFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type TriggerFulfillmentMessagesObservation struct {

	// A collection of text responses.
	Text []MessagesTextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type TriggerFulfillmentMessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []MessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type TriggerFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	Messages []MessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type TriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []MessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

// FlowSpec defines the desired state of Flow
type FlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlowParameters `json:"forProvider"`
}

// FlowStatus defines the observed state of Flow.
type FlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Flow is the Schema for the Flows API. Flows represents the conversation flows when you build your chatbot agent.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Flow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   FlowSpec   `json:"spec"`
	Status FlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowList contains a list of Flows
type FlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Flow `json:"items"`
}

// Repository type metadata.
var (
	Flow_Kind             = "Flow"
	Flow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Flow_Kind}.String()
	Flow_KindAPIVersion   = Flow_Kind + "." + CRDGroupVersion.String()
	Flow_GroupVersionKind = CRDGroupVersion.WithKind(Flow_Kind)
)

func init() {
	SchemeBuilder.Register(&Flow{}, &FlowList{})
}
