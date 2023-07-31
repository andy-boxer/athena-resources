package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Lid is a specification for a Lid resource
type Lid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LidSpec   `json:"spec"`
	Status LidStatus `json:"status"`
}

// LidSpec is the spec for a Lid resource
type LidSpec struct {
	ImageUrl       string `json:"ImageUrl"`
	DeploymentKind string `json:"DeploymentKind"`
}

// LidStatus is the status for a Lid resource
type LidStatus struct {
	AvailableReplicas int32  `json:"availableReplicas"`
	Phase             string `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LidList is a list of Lid resources
type LidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Lid `json:"items"`
}
