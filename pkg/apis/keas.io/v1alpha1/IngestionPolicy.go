package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
