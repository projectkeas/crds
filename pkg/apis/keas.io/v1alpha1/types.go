package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This is all of our custom type definitions
// Types must be registered in register.go in the addKnownTypes method

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="The description of the ingestion policy"
// +kubebuilder:printcolumn:name="Default Allow",type="boolean",JSONPath=".spec.defaults.allow",description="The default value for allow"
// +kubebuilder:printcolumn:name="Default TTL",type="integer",JSONPath=".spec.defaults.ttl",description="The default value for ttl"
type IngestionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec IngestionPolicySpecification `json:"spec"`
}

type IngestionPolicySpecification struct {
	Policy      string                  `json:"policy"`
	Description string                  `json:"description,omitempty"`
	Defaults    IngestionPolicyDefaults `json:"defaults,omitempty"`
}

type IngestionPolicyDefaults struct {
	Allow bool `json:"allow,omitempty"`
	TTL   int  `json:"ttl,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type IngestionPolicyList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []IngestionPolicy `json:"items,omitempty"`
}

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
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []EventType `json:"items,omitempty"`
}
