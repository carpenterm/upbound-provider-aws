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

type ReadinessCheckObservation struct {

	// ARN of the readiness_check
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName *string `json:"resourceSetName,omitempty" tf:"resource_set_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReadinessCheckParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name describing the resource set that will be monitored for readiness.
	// +kubebuilder:validation:Optional
	ResourceSetName *string `json:"resourceSetName,omitempty" tf:"resource_set_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ReadinessCheckSpec defines the desired state of ReadinessCheck
type ReadinessCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadinessCheckParameters `json:"forProvider"`
}

// ReadinessCheckStatus defines the observed state of ReadinessCheck.
type ReadinessCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadinessCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReadinessCheck is the Schema for the ReadinessChecks API. Provides an AWS Route 53 Recovery Readiness Readiness Check
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReadinessCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceSetName)",message="resourceSetName is a required parameter"
	Spec   ReadinessCheckSpec   `json:"spec"`
	Status ReadinessCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadinessCheckList contains a list of ReadinessChecks
type ReadinessCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadinessCheck `json:"items"`
}

// Repository type metadata.
var (
	ReadinessCheck_Kind             = "ReadinessCheck"
	ReadinessCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadinessCheck_Kind}.String()
	ReadinessCheck_KindAPIVersion   = ReadinessCheck_Kind + "." + CRDGroupVersion.String()
	ReadinessCheck_GroupVersionKind = CRDGroupVersion.WithKind(ReadinessCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadinessCheck{}, &ReadinessCheckList{})
}
