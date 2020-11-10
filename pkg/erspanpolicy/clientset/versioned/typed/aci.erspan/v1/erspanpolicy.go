/***
Copyright 2019 Cisco Systems Inc. All rights reserved.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/noironetworks/aci-containers/pkg/erspanpolicy/apis/aci.erspan/v1"
	scheme "github.com/noironetworks/aci-containers/pkg/erspanpolicy/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ErspanPoliciesGetter has a method to return a ErspanPolicyInterface.
// A group's client should implement this interface.
type ErspanPoliciesGetter interface {
	ErspanPolicies() ErspanPolicyInterface
}

// ErspanPolicyInterface has methods to work with ErspanPolicy resources.
type ErspanPolicyInterface interface {
	Create(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.CreateOptions) (*v1.ErspanPolicy, error)
	Update(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.UpdateOptions) (*v1.ErspanPolicy, error)
	UpdateStatus(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.UpdateOptions) (*v1.ErspanPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ErspanPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ErspanPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ErspanPolicy, err error)
	ErspanPolicyExpansion
}

// erspanPolicies implements ErspanPolicyInterface
type erspanPolicies struct {
	client rest.Interface
}

// newErspanPolicies returns a ErspanPolicies
func newErspanPolicies(c *AciV1Client) *erspanPolicies {
	return &erspanPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the erspanPolicy, and returns the corresponding erspanPolicy object, and an error if there is any.
func (c *erspanPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ErspanPolicy, err error) {
	result = &v1.ErspanPolicy{}
	err = c.client.Get().
		Resource("erspanpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ErspanPolicies that match those selectors.
func (c *erspanPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ErspanPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ErspanPolicyList{}
	err = c.client.Get().
		Resource("erspanpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested erspanPolicies.
func (c *erspanPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("erspanpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a erspanPolicy and creates it.  Returns the server's representation of the erspanPolicy, and an error, if there is any.
func (c *erspanPolicies) Create(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.CreateOptions) (result *v1.ErspanPolicy, err error) {
	result = &v1.ErspanPolicy{}
	err = c.client.Post().
		Resource("erspanpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(erspanPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a erspanPolicy and updates it. Returns the server's representation of the erspanPolicy, and an error, if there is any.
func (c *erspanPolicies) Update(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.UpdateOptions) (result *v1.ErspanPolicy, err error) {
	result = &v1.ErspanPolicy{}
	err = c.client.Put().
		Resource("erspanpolicies").
		Name(erspanPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(erspanPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *erspanPolicies) UpdateStatus(ctx context.Context, erspanPolicy *v1.ErspanPolicy, opts metav1.UpdateOptions) (result *v1.ErspanPolicy, err error) {
	result = &v1.ErspanPolicy{}
	err = c.client.Put().
		Resource("erspanpolicies").
		Name(erspanPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(erspanPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the erspanPolicy and deletes it. Returns an error if one occurs.
func (c *erspanPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("erspanpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *erspanPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("erspanpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched erspanPolicy.
func (c *erspanPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ErspanPolicy, err error) {
	result = &v1.ErspanPolicy{}
	err = c.client.Patch(pt).
		Resource("erspanpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}