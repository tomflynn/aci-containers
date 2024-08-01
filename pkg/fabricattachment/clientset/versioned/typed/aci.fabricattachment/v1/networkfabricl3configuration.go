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

// NetworkFabricL3ConfigurationsGetter has a method to return a NetworkFabricL3ConfigurationInterface.
// A group's client should implement this interface.
type NetworkFabricL3ConfigurationsGetter interface {
	NetworkFabricL3Configurations() NetworkFabricL3ConfigurationInterface
}

// NetworkFabricL3ConfigurationInterface has methods to work with NetworkFabricL3Configuration resources.
type NetworkFabricL3ConfigurationInterface interface {
	Create(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.CreateOptions) (*v1.NetworkFabricL3Configuration, error)
	Update(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.UpdateOptions) (*v1.NetworkFabricL3Configuration, error)
	UpdateStatus(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.UpdateOptions) (*v1.NetworkFabricL3Configuration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkFabricL3Configuration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkFabricL3ConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkFabricL3Configuration, err error)
	Apply(ctx context.Context, networkFabricL3Configuration *acifabricattachmentv1.NetworkFabricL3ConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkFabricL3Configuration, err error)
	ApplyStatus(ctx context.Context, networkFabricL3Configuration *acifabricattachmentv1.NetworkFabricL3ConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkFabricL3Configuration, err error)
	NetworkFabricL3ConfigurationExpansion
}

// networkFabricL3Configurations implements NetworkFabricL3ConfigurationInterface
type networkFabricL3Configurations struct {
	client rest.Interface
}

// newNetworkFabricL3Configurations returns a NetworkFabricL3Configurations
func newNetworkFabricL3Configurations(c *AciV1Client) *networkFabricL3Configurations {
	return &networkFabricL3Configurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the networkFabricL3Configuration, and returns the corresponding networkFabricL3Configuration object, and an error if there is any.
func (c *networkFabricL3Configurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Get().
		Resource("networkfabricl3configurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkFabricL3Configurations that match those selectors.
func (c *networkFabricL3Configurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkFabricL3ConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkFabricL3ConfigurationList{}
	err = c.client.Get().
		Resource("networkfabricl3configurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkFabricL3Configurations.
func (c *networkFabricL3Configurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networkfabricl3configurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkFabricL3Configuration and creates it.  Returns the server's representation of the networkFabricL3Configuration, and an error, if there is any.
func (c *networkFabricL3Configurations) Create(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.CreateOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Post().
		Resource("networkfabricl3configurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkFabricL3Configuration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkFabricL3Configuration and updates it. Returns the server's representation of the networkFabricL3Configuration, and an error, if there is any.
func (c *networkFabricL3Configurations) Update(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.UpdateOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Put().
		Resource("networkfabricl3configurations").
		Name(networkFabricL3Configuration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkFabricL3Configuration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkFabricL3Configurations) UpdateStatus(ctx context.Context, networkFabricL3Configuration *v1.NetworkFabricL3Configuration, opts metav1.UpdateOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Put().
		Resource("networkfabricl3configurations").
		Name(networkFabricL3Configuration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkFabricL3Configuration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkFabricL3Configuration and deletes it. Returns an error if one occurs.
func (c *networkFabricL3Configurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networkfabricl3configurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkFabricL3Configurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networkfabricl3configurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkFabricL3Configuration.
func (c *networkFabricL3Configurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkFabricL3Configuration, err error) {
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Patch(pt).
		Resource("networkfabricl3configurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkFabricL3Configuration.
func (c *networkFabricL3Configurations) Apply(ctx context.Context, networkFabricL3Configuration *acifabricattachmentv1.NetworkFabricL3ConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	if networkFabricL3Configuration == nil {
		return nil, fmt.Errorf("networkFabricL3Configuration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(networkFabricL3Configuration)
	if err != nil {
		return nil, err
	}
	name := networkFabricL3Configuration.Name
	if name == nil {
		return nil, fmt.Errorf("networkFabricL3Configuration.Name must be provided to Apply")
	}
	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("networkfabricl3configurations").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *networkFabricL3Configurations) ApplyStatus(ctx context.Context, networkFabricL3Configuration *acifabricattachmentv1.NetworkFabricL3ConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkFabricL3Configuration, err error) {
	if networkFabricL3Configuration == nil {
		return nil, fmt.Errorf("networkFabricL3Configuration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(networkFabricL3Configuration)
	if err != nil {
		return nil, err
	}

	name := networkFabricL3Configuration.Name
	if name == nil {
		return nil, fmt.Errorf("networkFabricL3Configuration.Name must be provided to Apply")
	}

	result = &v1.NetworkFabricL3Configuration{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("networkfabricl3configurations").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}