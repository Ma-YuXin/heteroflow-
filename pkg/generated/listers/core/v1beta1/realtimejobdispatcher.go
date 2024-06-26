/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "heterflow/api/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RealtimeJobDispatcherLister helps list RealtimeJobDispatchers.
// All objects returned here must be treated as read-only.
type RealtimeJobDispatcherLister interface {
	// List lists all RealtimeJobDispatchers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.RealtimeJobDispatcher, err error)
	// RealtimeJobDispatchers returns an object that can list and get RealtimeJobDispatchers.
	RealtimeJobDispatchers(namespace string) RealtimeJobDispatcherNamespaceLister
	RealtimeJobDispatcherListerExpansion
}

// realtimeJobDispatcherLister implements the RealtimeJobDispatcherLister interface.
type realtimeJobDispatcherLister struct {
	indexer cache.Indexer
}

// NewRealtimeJobDispatcherLister returns a new RealtimeJobDispatcherLister.
func NewRealtimeJobDispatcherLister(indexer cache.Indexer) RealtimeJobDispatcherLister {
	return &realtimeJobDispatcherLister{indexer: indexer}
}

// List lists all RealtimeJobDispatchers in the indexer.
func (s *realtimeJobDispatcherLister) List(selector labels.Selector) (ret []*v1beta1.RealtimeJobDispatcher, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RealtimeJobDispatcher))
	})
	return ret, err
}

// RealtimeJobDispatchers returns an object that can list and get RealtimeJobDispatchers.
func (s *realtimeJobDispatcherLister) RealtimeJobDispatchers(namespace string) RealtimeJobDispatcherNamespaceLister {
	return realtimeJobDispatcherNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RealtimeJobDispatcherNamespaceLister helps list and get RealtimeJobDispatchers.
// All objects returned here must be treated as read-only.
type RealtimeJobDispatcherNamespaceLister interface {
	// List lists all RealtimeJobDispatchers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.RealtimeJobDispatcher, err error)
	// Get retrieves the RealtimeJobDispatcher from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.RealtimeJobDispatcher, error)
	RealtimeJobDispatcherNamespaceListerExpansion
}

// realtimeJobDispatcherNamespaceLister implements the RealtimeJobDispatcherNamespaceLister
// interface.
type realtimeJobDispatcherNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RealtimeJobDispatchers in the indexer for a given namespace.
func (s realtimeJobDispatcherNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.RealtimeJobDispatcher, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RealtimeJobDispatcher))
	})
	return ret, err
}

// Get retrieves the RealtimeJobDispatcher from the indexer for a given namespace and name.
func (s realtimeJobDispatcherNamespaceLister) Get(name string) (*v1beta1.RealtimeJobDispatcher, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("realtimejobdispatcher"), name)
	}
	return obj.(*v1beta1.RealtimeJobDispatcher), nil
}
