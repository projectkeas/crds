package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Event Type",type="string",JSONPath=".spec.type",description="The name of the type that will be accepted by the ingestion engine"
// +kubebuilder:printcolumn:name="Owner",type="string",JSONPath=".spec.owner",description="The owner of the event type definition"
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="The description of the event type and/or schema"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version",description="The version of the event type that's being stored"
type EventType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec EventTypeSpecification `json:"spec"`
}

type EventTypeSpecification struct {
	// +kubebuilder:validation:Pattern=`^([0-9]{1,4}){1}\.([0-9]{1,4}){1}\.([0-9]{1,4}){1}$`
	Version string `json:"version"`
	Schema  string `json:"schema"`
	// +kubebuilder:validation:Pattern=`^[A-z\-]{3,63}$`
	Owner string `json:"owner"`
	// +kubebuilder:validation:Pattern=`^[A-z]{3,63}$`
	EventType string   `json:"type"`
	Sources   []string `json:"sources,omitempty"`
	SubTypes  []string `json:"subTypes,omitempty"`
	// +kubebuilder:validation:Pattern=`^[\w .,&\-'"]*$`
	Description string `json:"description,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EventTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EventType `json:"items"`
}
