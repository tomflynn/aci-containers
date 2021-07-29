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
	v1 "github.com/noironetworks/aci-containers/pkg/snatpolicy/apis/aci.snat/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SnatPolicyLister helps list SnatPolicies.
// All objects returned here must be treated as read-only.
type SnatPolicyLister interface {
	// List lists all SnatPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SnatPolicy, err error)
	// Get retrieves the SnatPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SnatPolicy, error)
	SnatPolicyListerExpansion
}

// snatPolicyLister implements the SnatPolicyLister interface.
type snatPolicyLister struct {
	indexer cache.Indexer
}

// NewSnatPolicyLister returns a new SnatPolicyLister.
func NewSnatPolicyLister(indexer cache.Indexer) SnatPolicyLister {
	return &snatPolicyLister{indexer: indexer}
}

// List lists all SnatPolicies in the indexer.
func (s *snatPolicyLister) List(selector labels.Selector) (ret []*v1.SnatPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SnatPolicy))
	})
	return ret, err
}

// Get retrieves the SnatPolicy from the index for a given name.
func (s *snatPolicyLister) Get(name string) (*v1.SnatPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("snatpolicy"), name)
	}
	return obj.(*v1.SnatPolicy), nil
}
