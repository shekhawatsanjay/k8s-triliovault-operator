package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrilioVaultSpec defines the desired state of TrilioVault
type TrilioVaultSpec struct {

	// Scope for the application which will be installed in the cluster
	// NamespaceScope or ClusterScope
	ApplicationScope string `json:"applicationScope,omitempty"`

	// Namespace in which the application will be installed
	ApplicationNamespace string `json:"applicationNamespace,omitempty"`

	// Labels specifies the labels to attach to pods the operator creates for the
	// application.
	Labels map[string]string `json:"labels,omitempty"`

	// Additional annotaions to be applied to pods during installation
	Annotations map[string]string `json:"annotations,omitempty"`

	// Resources is the resource requirements for the containers.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`

	// NodeSelector specifies a map of key-value pairs. For the pod to be eligible
	// to run on a node, the node must have each of the indicated key-value pairs as
	// labels.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// The scheduling constraints on application pods.
	Affinity *v1.Affinity `json:"affinity,omitempty"`
}

// TrilioVaultStatus defines the observed state of TrilioVault
type TrilioVaultStatus struct {
}

// +kubebuilder:object:root=true

// TrilioVault is the Schema for the triliovaults API
type TrilioVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrilioVaultSpec   `json:"spec,omitempty"`
	Status TrilioVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrilioVaultList contains a list of TrilioVault
type TrilioVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrilioVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrilioVault{}, &TrilioVaultList{})
}
