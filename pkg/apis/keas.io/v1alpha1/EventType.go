package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="URI",type="string",JSONPath=".spec.schemaUri",description="The schema uri that's mapped to the dataschema property of cloudevents"
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="The description of the event type and/or schema"
type EventType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec EventTypeSpecification `json:"spec"`
}

type EventTypeSpecification struct {
	Schema string `json:"schema"`
	// +kubebuilder:validation:Format=uri
	SchemaUri string `json:"schemaUri"`
	// +kubebuilder:validation:Pattern=`^[\w .,&\-'"]*$`
	Description string `json:"description,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EventTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EventType `json:"items"`
}
