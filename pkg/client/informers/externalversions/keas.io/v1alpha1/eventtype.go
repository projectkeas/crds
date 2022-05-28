// Auto-generated for github.com/projectkeas/crds
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	keasiov1alpha1 "github.com/projectkeas/crds/pkg/apis/keas.io/v1alpha1"
	versioned "github.com/projectkeas/crds/pkg/client/clientset/versioned"
	internalinterfaces "github.com/projectkeas/crds/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/projectkeas/crds/pkg/client/listers/keas.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EventTypeInformer provides access to a shared informer and lister for
// EventTypes.
type EventTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EventTypeLister
}

type eventTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEventTypeInformer constructs a new informer for EventType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEventTypeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEventTypeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEventTypeInformer constructs a new informer for EventType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEventTypeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeasV1alpha1().EventTypes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeasV1alpha1().EventTypes(namespace).Watch(context.TODO(), options)
			},
		},
		&keasiov1alpha1.EventType{},
		resyncPeriod,
		indexers,
	)
}

func (f *eventTypeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEventTypeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *eventTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&keasiov1alpha1.EventType{}, f.defaultInformer)
}

func (f *eventTypeInformer) Lister() v1alpha1.EventTypeLister {
	return v1alpha1.NewEventTypeLister(f.Informer().GetIndexer())
}