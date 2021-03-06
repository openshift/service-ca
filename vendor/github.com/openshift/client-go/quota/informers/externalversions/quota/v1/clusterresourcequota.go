// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	quota_v1 "github.com/openshift/api/quota/v1"
	versioned "github.com/openshift/client-go/quota/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/quota/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/quota/listers/quota/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterResourceQuotaInformer provides access to a shared informer and lister for
// ClusterResourceQuotas.
type ClusterResourceQuotaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterResourceQuotaLister
}

type clusterResourceQuotaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterResourceQuotaInformer constructs a new informer for ClusterResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterResourceQuotaInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterResourceQuotaInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterResourceQuotaInformer constructs a new informer for ClusterResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterResourceQuotaInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuotaV1().ClusterResourceQuotas().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuotaV1().ClusterResourceQuotas().Watch(options)
			},
		},
		&quota_v1.ClusterResourceQuota{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterResourceQuotaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterResourceQuotaInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterResourceQuotaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quota_v1.ClusterResourceQuota{}, f.defaultInformer)
}

func (f *clusterResourceQuotaInformer) Lister() v1.ClusterResourceQuotaLister {
	return v1.NewClusterResourceQuotaLister(f.Informer().GetIndexer())
}
