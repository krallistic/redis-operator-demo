/*
Created by codegen
*/
// This file was automatically generated by informer-gen

package v1

import (
	krallistic_com_v1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1"
	versioned "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned"
	internalinterfaces "github.com/krallistic/redis-operator-demo/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/krallistic/redis-operator-demo/pkg/client/listers/krallistic/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RedisdbInformer provides access to a shared informer and lister for
// Redisdbs.
type RedisdbInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RedisdbLister
}

type redisdbInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewRedisdbInformer constructs a new informer for Redisdb type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRedisdbInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.KrallisticV1().Redisdbs(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.KrallisticV1().Redisdbs(namespace).Watch(options)
			},
		},
		&krallistic_com_v1.Redisdb{},
		resyncPeriod,
		indexers,
	)
}

func defaultRedisdbInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewRedisdbInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *redisdbInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&krallistic_com_v1.Redisdb{}, defaultRedisdbInformer)
}

func (f *redisdbInformer) Lister() v1.RedisdbLister {
	return v1.NewRedisdbLister(f.Informer().GetIndexer())
}
