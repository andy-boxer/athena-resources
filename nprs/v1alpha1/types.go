package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Npr is a specification for a Npr resource
type Npr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NprSpec   `json:"spec"`
	Status NprStatus `json:"status"`
}

// NprSpec is the spec for a Npr resource
type NprSpec struct {
	Service     string `json:"Service"`
	ServiceHost string `json:"ServiceHost"`
	ServicePort int32  `json:"ServicePort"`
	ServiceTier string `json:"ServiceTier"`
}

// NprStatus is the status for a Npr resource
type NprStatus struct {
	Phase string `json:"Phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NprList is a list of Npr resources
type NprList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Npr `json:"items"`
}
