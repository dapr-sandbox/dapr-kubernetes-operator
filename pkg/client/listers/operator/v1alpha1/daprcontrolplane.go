/*
Copyright 2023.

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

package v1alpha1

import (
	v1alpha1 "github.com/dapr-sandbox/dapr-kubernetes-operator/api/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DaprControlPlaneLister helps list DaprControlPlanes.
// All objects returned here must be treated as read-only.
type DaprControlPlaneLister interface {
	// List lists all DaprControlPlanes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DaprControlPlane, err error)
	// DaprControlPlanes returns an object that can list and get DaprControlPlanes.
	DaprControlPlanes(namespace string) DaprControlPlaneNamespaceLister
	DaprControlPlaneListerExpansion
}

// daprControlPlaneLister implements the DaprControlPlaneLister interface.
type daprControlPlaneLister struct {
	indexer cache.Indexer
}

// NewDaprControlPlaneLister returns a new DaprControlPlaneLister.
func NewDaprControlPlaneLister(indexer cache.Indexer) DaprControlPlaneLister {
	return &daprControlPlaneLister{indexer: indexer}
}

// List lists all DaprControlPlanes in the indexer.
func (s *daprControlPlaneLister) List(selector labels.Selector) (ret []*v1alpha1.DaprControlPlane, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DaprControlPlane))
	})
	return ret, err
}

// DaprControlPlanes returns an object that can list and get DaprControlPlanes.
func (s *daprControlPlaneLister) DaprControlPlanes(namespace string) DaprControlPlaneNamespaceLister {
	return daprControlPlaneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DaprControlPlaneNamespaceLister helps list and get DaprControlPlanes.
// All objects returned here must be treated as read-only.
type DaprControlPlaneNamespaceLister interface {
	// List lists all DaprControlPlanes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DaprControlPlane, err error)
	// Get retrieves the DaprControlPlane from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DaprControlPlane, error)
	DaprControlPlaneNamespaceListerExpansion
}

// daprControlPlaneNamespaceLister implements the DaprControlPlaneNamespaceLister
// interface.
type daprControlPlaneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DaprControlPlanes in the indexer for a given namespace.
func (s daprControlPlaneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DaprControlPlane, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DaprControlPlane))
	})
	return ret, err
}

// Get retrieves the DaprControlPlane from the indexer for a given namespace and name.
func (s daprControlPlaneNamespaceLister) Get(name string) (*v1alpha1.DaprControlPlane, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("daprcontrolplane"), name)
	}
	return obj.(*v1alpha1.DaprControlPlane), nil
}