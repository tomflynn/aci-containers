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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/noironetworks/aci-containers/pkg/gbpcrd/apis/acipolicy/v1"
	scheme "github.com/noironetworks/aci-containers/pkg/gbpcrd/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PodIFsGetter has a method to return a PodIFInterface.
// A group's client should implement this interface.
type PodIFsGetter interface {
	PodIFs(namespace string) PodIFInterface
}

// PodIFInterface has methods to work with PodIF resources.
type PodIFInterface interface {
	Create(ctx context.Context, podIF *v1.PodIF, opts metav1.CreateOptions) (*v1.PodIF, error)
	Update(ctx context.Context, podIF *v1.PodIF, opts metav1.UpdateOptions) (*v1.PodIF, error)
	UpdateStatus(ctx context.Context, podIF *v1.PodIF, opts metav1.UpdateOptions) (*v1.PodIF, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PodIF, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PodIFList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PodIF, err error)
	PodIFExpansion
}

// podIFs implements PodIFInterface
type podIFs struct {
	client rest.Interface
	ns     string
}

// newPodIFs returns a PodIFs
func newPodIFs(c *AciV1Client, namespace string) *podIFs {
	return &podIFs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podIF, and returns the corresponding podIF object, and an error if there is any.
func (c *podIFs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PodIF, err error) {
	result = &v1.PodIF{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podifs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodIFs that match those selectors.
func (c *podIFs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PodIFList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PodIFList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podIFs.
func (c *podIFs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podIF and creates it.  Returns the server's representation of the podIF, and an error, if there is any.
func (c *podIFs) Create(ctx context.Context, podIF *v1.PodIF, opts metav1.CreateOptions) (result *v1.PodIF, err error) {
	result = &v1.PodIF{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podIF).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podIF and updates it. Returns the server's representation of the podIF, and an error, if there is any.
func (c *podIFs) Update(ctx context.Context, podIF *v1.PodIF, opts metav1.UpdateOptions) (result *v1.PodIF, err error) {
	result = &v1.PodIF{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podifs").
		Name(podIF.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podIF).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *podIFs) UpdateStatus(ctx context.Context, podIF *v1.PodIF, opts metav1.UpdateOptions) (result *v1.PodIF, err error) {
	result = &v1.PodIF{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podifs").
		Name(podIF.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podIF).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podIF and deletes it. Returns an error if one occurs.
func (c *podIFs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podifs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podIFs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podifs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podIF.
func (c *podIFs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PodIF, err error) {
	result = &v1.PodIF{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podifs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
