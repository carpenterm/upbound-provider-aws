/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BusPolicyObservation struct {

	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the default event bus.
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// The name of the EventBridge event bus.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The text of the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type BusPolicyParameters struct {

	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the default event bus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Bus
	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus in cloudwatchevents to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRef *v1.Reference `json:"eventBusNameRef,omitempty" tf:"-"`

	// Selector for a Bus in cloudwatchevents to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// The text of the policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BusPolicySpec defines the desired state of BusPolicy
type BusPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BusPolicyParameters `json:"forProvider"`
}

// BusPolicyStatus defines the observed state of BusPolicy.
type BusPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BusPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BusPolicy is the Schema for the BusPolicys API. Provides a resource to create an EventBridge policy to support cross-account events.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BusPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy)",message="policy is a required parameter"
	Spec   BusPolicySpec   `json:"spec"`
	Status BusPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BusPolicyList contains a list of BusPolicys
type BusPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BusPolicy `json:"items"`
}

// Repository type metadata.
var (
	BusPolicy_Kind             = "BusPolicy"
	BusPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BusPolicy_Kind}.String()
	BusPolicy_KindAPIVersion   = BusPolicy_Kind + "." + CRDGroupVersion.String()
	BusPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BusPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BusPolicy{}, &BusPolicyList{})
}
