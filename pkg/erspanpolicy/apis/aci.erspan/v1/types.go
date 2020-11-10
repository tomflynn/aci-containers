package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PolicyState string

const (
	Ready  PolicyState = "Ready"
	Failed PolicyState = "Failed"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ErspanPolicySpec defines the desired state of ErspanPolicy
type ErspanPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Selector PodSelector        `json:"selector,omitempty"`
	Source   ErspanPolicingType `json:"source,omitempty"`
	Dest     ErspanPolicingType `json:"destination,omitempty"`
}

//ErspanPolicingType contains all the attrbutes of erspan source and destination sessions.
type ErspanPolicingType struct {
	// +kubebuilder:validation:Enum=start,stop
	AdminState string `json:"admin_state"`
	// +kubebuilder:validation:Enum=in,out,both
	Direction  string `json:"direction"`
	DestIp     string `json:"destIp"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1023
	FlowId     int    `json:"flowId"`
	Tag        string `json:"tag,omitempty"`
}

// ErspanPolicyStatus defines the observed state of ErspanPolicy
type ErspanPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file7
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	State PolicyState `json:"state"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ErspanPolicy is the Schema for the erspanpolicies API
type ErspanPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ErspanPolicySpec   `json:"spec,omitempty"`
	Status ErspanPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ErspanPolicyList contains a list of ErspanPolicy
type ErspanPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ErspanPolicy `json:"items"`
}

type PodSelector struct {
	Labels    map[string]string `json:"labels,omitempt"`
	Namespace string            `json:"namespace,omitempty"`
}
