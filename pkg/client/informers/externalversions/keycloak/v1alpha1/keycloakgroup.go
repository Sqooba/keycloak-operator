// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	keycloakv1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	versioned "github.com/keycloak/keycloak-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/keycloak/keycloak-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/client/listers/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KeycloakGroupInformer provides access to a shared informer and lister for
// KeycloakGroups.
type KeycloakGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KeycloakGroupLister
}

type keycloakGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKeycloakGroupInformer constructs a new informer for KeycloakGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKeycloakGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKeycloakGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKeycloakGroupInformer constructs a new informer for KeycloakGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKeycloakGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&keycloakv1alpha1.KeycloakGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *keycloakGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKeycloakGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *keycloakGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&keycloakv1alpha1.KeycloakGroup{}, f.defaultInformer)
}

func (f *keycloakGroupInformer) Lister() v1alpha1.KeycloakGroupLister {
	return v1alpha1.NewKeycloakGroupLister(f.Informer().GetIndexer())
}
