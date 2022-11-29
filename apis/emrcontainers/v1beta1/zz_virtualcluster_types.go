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

type ContainerProviderObservation struct {
}

type ContainerProviderParameters struct {

	// The name of the container provider that is running your EMR Containers cluster
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Cluster in eks to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Cluster in eks to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Nested list containing information about the configuration of the container provider
	// +kubebuilder:validation:Required
	Info []InfoParameters `json:"info" tf:"info,omitempty"`

	// The type of the container provider
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EksInfoObservation struct {
}

type EksInfoParameters struct {

	// The namespace where the EMR Containers cluster is running
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type InfoObservation struct {
}

type InfoParameters struct {

	// Nested list containing EKS-specific information about the cluster where the EMR Containers cluster is running
	// +kubebuilder:validation:Required
	EksInfo []EksInfoParameters `json:"eksInfo" tf:"eks_info,omitempty"`
}

type VirtualClusterObservation struct {

	// ARN of the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the container provider that is running your EMR Containers cluster
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualClusterParameters struct {

	// Configuration block for the container provider associated with your cluster.
	// +kubebuilder:validation:Required
	ContainerProvider []ContainerProviderParameters `json:"containerProvider" tf:"container_provider,omitempty"`

	// –  Name of the virtual cluster.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VirtualClusterSpec defines the desired state of VirtualCluster
type VirtualClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualClusterParameters `json:"forProvider"`
}

// VirtualClusterStatus defines the observed state of VirtualCluster.
type VirtualClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualCluster is the Schema for the VirtualClusters API. Manages an EMR Containers (EMR on EKS) Virtual Cluster
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualClusterSpec   `json:"spec"`
	Status            VirtualClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualClusterList contains a list of VirtualClusters
type VirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCluster `json:"items"`
}

// Repository type metadata.
var (
	VirtualCluster_Kind             = "VirtualCluster"
	VirtualCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualCluster_Kind}.String()
	VirtualCluster_KindAPIVersion   = VirtualCluster_Kind + "." + CRDGroupVersion.String()
	VirtualCluster_GroupVersionKind = CRDGroupVersion.WithKind(VirtualCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualCluster{}, &VirtualClusterList{})
}