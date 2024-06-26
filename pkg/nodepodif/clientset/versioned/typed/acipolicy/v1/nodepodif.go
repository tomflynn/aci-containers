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

	v1 "github.com/noironetworks/aci-containers/pkg/nodepodif/apis/acipolicy/v1"
	scheme "github.com/noironetworks/aci-containers/pkg/nodepodif/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodePodIFsGetter has a method to return a NodePodIFInterface.
// A group's client should implement this interface.
type NodePodIFsGetter interface {
	NodePodIFs(namespace string) NodePodIFInterface
}

// NodePodIFInterface has methods to work with NodePodIF resources.
type NodePodIFInterface interface {
	Create(ctx context.Context, nodePodIF *v1.NodePodIF, opts metav1.CreateOptions) (*v1.NodePodIF, error)
	Update(ctx context.Context, nodePodIF *v1.NodePodIF, opts metav1.UpdateOptions) (*v1.NodePodIF, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NodePodIF, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NodePodIFList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodePodIF, err error)
	NodePodIFExpansion
}

// nodePodIFs implements NodePodIFInterface
type nodePodIFs struct {
	client rest.Interface
	ns     string
}

// newNodePodIFs returns a NodePodIFs
func newNodePodIFs(c *AciV1Client, namespace string) *nodePodIFs {
	return &nodePodIFs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodePodIF, and returns the corresponding nodePodIF object, and an error if there is any.
func (c *nodePodIFs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NodePodIF, err error) {
	result = &v1.NodePodIF{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepodifs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodePodIFs that match those selectors.
func (c *nodePodIFs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodePodIFList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NodePodIFList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepodifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodePodIFs.
func (c *nodePodIFs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodepodifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodePodIF and creates it.  Returns the server's representation of the nodePodIF, and an error, if there is any.
func (c *nodePodIFs) Create(ctx context.Context, nodePodIF *v1.NodePodIF, opts metav1.CreateOptions) (result *v1.NodePodIF, err error) {
	result = &v1.NodePodIF{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodepodifs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodePodIF).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodePodIF and updates it. Returns the server's representation of the nodePodIF, and an error, if there is any.
func (c *nodePodIFs) Update(ctx context.Context, nodePodIF *v1.NodePodIF, opts metav1.UpdateOptions) (result *v1.NodePodIF, err error) {
	result = &v1.NodePodIF{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodepodifs").
		Name(nodePodIF.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodePodIF).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodePodIF and deletes it. Returns an error if one occurs.
func (c *nodePodIFs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepodifs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodePodIFs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepodifs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodePodIF.
func (c *nodePodIFs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodePodIF, err error) {
	result = &v1.NodePodIF{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodepodifs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
