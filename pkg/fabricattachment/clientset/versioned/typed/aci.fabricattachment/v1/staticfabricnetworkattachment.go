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
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/apis/aci.fabricattachment/v1"
	acifabricattachmentv1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/applyconfiguration/aci.fabricattachment/v1"
	scheme "github.com/noironetworks/aci-containers/pkg/fabricattachment/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StaticFabricNetworkAttachmentsGetter has a method to return a StaticFabricNetworkAttachmentInterface.
// A group's client should implement this interface.
type StaticFabricNetworkAttachmentsGetter interface {
	StaticFabricNetworkAttachments(namespace string) StaticFabricNetworkAttachmentInterface
}

// StaticFabricNetworkAttachmentInterface has methods to work with StaticFabricNetworkAttachment resources.
type StaticFabricNetworkAttachmentInterface interface {
	Create(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.CreateOptions) (*v1.StaticFabricNetworkAttachment, error)
	Update(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (*v1.StaticFabricNetworkAttachment, error)
	UpdateStatus(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (*v1.StaticFabricNetworkAttachment, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.StaticFabricNetworkAttachment, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.StaticFabricNetworkAttachmentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StaticFabricNetworkAttachment, err error)
	Apply(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error)
	ApplyStatus(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error)
	StaticFabricNetworkAttachmentExpansion
}

// staticFabricNetworkAttachments implements StaticFabricNetworkAttachmentInterface
type staticFabricNetworkAttachments struct {
	client rest.Interface
	ns     string
}

// newStaticFabricNetworkAttachments returns a StaticFabricNetworkAttachments
func newStaticFabricNetworkAttachments(c *AciV1Client, namespace string) *staticFabricNetworkAttachments {
	return &staticFabricNetworkAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the staticFabricNetworkAttachment, and returns the corresponding staticFabricNetworkAttachment object, and an error if there is any.
func (c *staticFabricNetworkAttachments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StaticFabricNetworkAttachments that match those selectors.
func (c *staticFabricNetworkAttachments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StaticFabricNetworkAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.StaticFabricNetworkAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested staticFabricNetworkAttachments.
func (c *staticFabricNetworkAttachments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a staticFabricNetworkAttachment and creates it.  Returns the server's representation of the staticFabricNetworkAttachment, and an error, if there is any.
func (c *staticFabricNetworkAttachments) Create(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.CreateOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(staticFabricNetworkAttachment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a staticFabricNetworkAttachment and updates it. Returns the server's representation of the staticFabricNetworkAttachment, and an error, if there is any.
func (c *staticFabricNetworkAttachments) Update(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(staticFabricNetworkAttachment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(staticFabricNetworkAttachment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *staticFabricNetworkAttachments) UpdateStatus(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(staticFabricNetworkAttachment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(staticFabricNetworkAttachment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the staticFabricNetworkAttachment and deletes it. Returns an error if one occurs.
func (c *staticFabricNetworkAttachments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *staticFabricNetworkAttachments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched staticFabricNetworkAttachment.
func (c *staticFabricNetworkAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StaticFabricNetworkAttachment, err error) {
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied staticFabricNetworkAttachment.
func (c *staticFabricNetworkAttachments) Apply(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	if staticFabricNetworkAttachment == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(staticFabricNetworkAttachment)
	if err != nil {
		return nil, err
	}
	name := staticFabricNetworkAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment.Name must be provided to Apply")
	}
	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *staticFabricNetworkAttachments) ApplyStatus(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	if staticFabricNetworkAttachment == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(staticFabricNetworkAttachment)
	if err != nil {
		return nil, err
	}

	name := staticFabricNetworkAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment.Name must be provided to Apply")
	}

	result = &v1.StaticFabricNetworkAttachment{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("staticfabricnetworkattachments").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}