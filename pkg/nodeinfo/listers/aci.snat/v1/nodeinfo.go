/***
Copyright 2021 Cisco Systems Inc. All rights reserved.

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

package v1

import (
	v1 "github.com/noironetworks/aci-containers/pkg/nodeinfo/apis/aci.snat/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeInfoLister helps list NodeInfos.
// All objects returned here must be treated as read-only.
type NodeInfoLister interface {
	// List lists all NodeInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NodeInfo, err error)
	// NodeInfos returns an object that can list and get NodeInfos.
	NodeInfos(namespace string) NodeInfoNamespaceLister
	NodeInfoListerExpansion
}

// nodeInfoLister implements the NodeInfoLister interface.
type nodeInfoLister struct {
	indexer cache.Indexer
}

// NewNodeInfoLister returns a new NodeInfoLister.
func NewNodeInfoLister(indexer cache.Indexer) NodeInfoLister {
	return &nodeInfoLister{indexer: indexer}
}

// List lists all NodeInfos in the indexer.
func (s *nodeInfoLister) List(selector labels.Selector) (ret []*v1.NodeInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeInfo))
	})
	return ret, err
}

// NodeInfos returns an object that can list and get NodeInfos.
func (s *nodeInfoLister) NodeInfos(namespace string) NodeInfoNamespaceLister {
	return nodeInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeInfoNamespaceLister helps list and get NodeInfos.
// All objects returned here must be treated as read-only.
type NodeInfoNamespaceLister interface {
	// List lists all NodeInfos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NodeInfo, err error)
	// Get retrieves the NodeInfo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NodeInfo, error)
	NodeInfoNamespaceListerExpansion
}

// nodeInfoNamespaceLister implements the NodeInfoNamespaceLister
// interface.
type nodeInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeInfos in the indexer for a given namespace.
func (s nodeInfoNamespaceLister) List(selector labels.Selector) (ret []*v1.NodeInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeInfo))
	})
	return ret, err
}

// Get retrieves the NodeInfo from the indexer for a given namespace and name.
func (s nodeInfoNamespaceLister) Get(name string) (*v1.NodeInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodeinfo"), name)
	}
	return obj.(*v1.NodeInfo), nil
}
