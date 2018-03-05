/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package internalversion

import (
	federation "github.com/marun/fnord/pkg/apis/federation"
	internalclientset "github.com/marun/fnord/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/marun/fnord/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/marun/fnord/pkg/client/listers_generated/federation/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedReplicaSetPlacementInformer provides access to a shared informer and lister for
// FederatedReplicaSetPlacements.
type FederatedReplicaSetPlacementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FederatedReplicaSetPlacementLister
}

type federatedReplicaSetPlacementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedReplicaSetPlacementInformer constructs a new informer for FederatedReplicaSetPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedReplicaSetPlacementInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedReplicaSetPlacementInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedReplicaSetPlacementInformer constructs a new informer for FederatedReplicaSetPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedReplicaSetPlacementInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedReplicaSetPlacements(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedReplicaSetPlacements(namespace).Watch(options)
			},
		},
		&federation.FederatedReplicaSetPlacement{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedReplicaSetPlacementInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedReplicaSetPlacementInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedReplicaSetPlacementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation.FederatedReplicaSetPlacement{}, f.defaultInformer)
}

func (f *federatedReplicaSetPlacementInformer) Lister() internalversion.FederatedReplicaSetPlacementLister {
	return internalversion.NewFederatedReplicaSetPlacementLister(f.Informer().GetIndexer())
}
