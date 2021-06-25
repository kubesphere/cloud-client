/*
Copyright 2021 The Kubesphere Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	storev1alpha1 "github.com/kubesphere/cloud-client/apis/store/v1alpha1"
	versioned "github.com/kubesphere/cloud-client/clientset/versioned"
	internalinterfaces "github.com/kubesphere/cloud-client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubesphere/cloud-client/listers/store/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// S3LocationInformer provides access to a shared informer and lister for
// S3Locations.
type S3LocationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.S3LocationLister
}

type s3LocationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewS3LocationInformer constructs a new informer for S3Location type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewS3LocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredS3LocationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredS3LocationInformer constructs a new informer for S3Location type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredS3LocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StoreV1alpha1().S3Locations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StoreV1alpha1().S3Locations(namespace).Watch(context.TODO(), options)
			},
		},
		&storev1alpha1.S3Location{},
		resyncPeriod,
		indexers,
	)
}

func (f *s3LocationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredS3LocationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *s3LocationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storev1alpha1.S3Location{}, f.defaultInformer)
}

func (f *s3LocationInformer) Lister() v1alpha1.S3LocationLister {
	return v1alpha1.NewS3LocationLister(f.Informer().GetIndexer())
}
