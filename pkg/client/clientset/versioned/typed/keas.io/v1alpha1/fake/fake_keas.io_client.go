// Auto-generated for github.com/projectkeas/crds
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/projectkeas/crds/pkg/client/clientset/versioned/typed/keas.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKeasV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKeasV1alpha1) EventTypes(namespace string) v1alpha1.EventTypeInterface {
	return &FakeEventTypes{c, namespace}
}

func (c *FakeKeasV1alpha1) IngestionPolicies(namespace string) v1alpha1.IngestionPolicyInterface {
	return &FakeIngestionPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKeasV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
