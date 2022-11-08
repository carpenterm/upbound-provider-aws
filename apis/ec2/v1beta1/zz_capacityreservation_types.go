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

type CapacityReservationObservation struct {

	// The ARN of the Capacity Reservation.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Capacity Reservation ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS account that owns the Capacity Reservation.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CapacityReservationParameters struct {

	// The Availability Zone in which to create the Capacity Reservation.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// Indicates the way in which the Capacity Reservation ends. Specify either unlimited or limited.
	// +kubebuilder:validation:Optional
	EndDateType *string `json:"endDateType,omitempty" tf:"end_date_type,omitempty"`

	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	// +kubebuilder:validation:Optional
	EphemeralStorage *bool `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// The number of instances for which to reserve capacity.
	// +kubebuilder:validation:Required
	InstanceCount *float64 `json:"instanceCount" tf:"instance_count,omitempty"`

	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either open or targeted.
	// +kubebuilder:validation:Optional
	InstanceMatchCriteria *string `json:"instanceMatchCriteria,omitempty" tf:"instance_match_criteria,omitempty"`

	// The type of operating system for which to reserve capacity. Valid options are Linux/UNIX, Red Hat Enterprise Linux, SUSE Linux, Windows, Windows with SQL Server, Windows with SQL Server Enterprise, Windows with SQL Server Standard or Windows with SQL Server Web.
	// +kubebuilder:validation:Required
	InstancePlatform *string `json:"instancePlatform" tf:"instance_platform,omitempty"`

	// The instance type for which to reserve capacity.
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates the tenancy of the Capacity Reservation. Specify either default or dedicated.
	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

// CapacityReservationSpec defines the desired state of CapacityReservation
type CapacityReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CapacityReservationParameters `json:"forProvider"`
}

// CapacityReservationStatus defines the observed state of CapacityReservation.
type CapacityReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CapacityReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityReservation is the Schema for the CapacityReservations API. Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CapacityReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CapacityReservationSpec   `json:"spec"`
	Status            CapacityReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityReservationList contains a list of CapacityReservations
type CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CapacityReservation `json:"items"`
}

// Repository type metadata.
var (
	CapacityReservation_Kind             = "CapacityReservation"
	CapacityReservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CapacityReservation_Kind}.String()
	CapacityReservation_KindAPIVersion   = CapacityReservation_Kind + "." + CRDGroupVersion.String()
	CapacityReservation_GroupVersionKind = CRDGroupVersion.WithKind(CapacityReservation_Kind)
)

func init() {
	SchemeBuilder.Register(&CapacityReservation{}, &CapacityReservationList{})
}
