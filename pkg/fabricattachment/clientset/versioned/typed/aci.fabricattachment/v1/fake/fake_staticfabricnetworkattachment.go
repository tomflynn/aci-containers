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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/apis/aci.fabricattachment/v1"
	acifabricattachmentv1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/applyconfiguration/aci.fabricattachment/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStaticFabricNetworkAttachments implements StaticFabricNetworkAttachmentInterface
type FakeStaticFabricNetworkAttachments struct {
	Fake *FakeAciV1
	ns   string
}

var staticfabricnetworkattachmentsResource = v1.SchemeGroupVersion.WithResource("staticfabricnetworkattachments")

var staticfabricnetworkattachmentsKind = v1.SchemeGroupVersion.WithKind("StaticFabricNetworkAttachment")

// Get takes name of the staticFabricNetworkAttachment, and returns the corresponding staticFabricNetworkAttachment object, and an error if there is any.
func (c *FakeStaticFabricNetworkAttachments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(staticfabricnetworkattachmentsResource, c.ns, name), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// List takes label and field selectors, and returns the list of StaticFabricNetworkAttachments that match those selectors.
func (c *FakeStaticFabricNetworkAttachments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StaticFabricNetworkAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(staticfabricnetworkattachmentsResource, staticfabricnetworkattachmentsKind, c.ns, opts), &v1.StaticFabricNetworkAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.StaticFabricNetworkAttachmentList{ListMeta: obj.(*v1.StaticFabricNetworkAttachmentList).ListMeta}
	for _, item := range obj.(*v1.StaticFabricNetworkAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested staticFabricNetworkAttachments.
func (c *FakeStaticFabricNetworkAttachments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(staticfabricnetworkattachmentsResource, c.ns, opts))

}

// Create takes the representation of a staticFabricNetworkAttachment and creates it.  Returns the server's representation of the staticFabricNetworkAttachment, and an error, if there is any.
func (c *FakeStaticFabricNetworkAttachments) Create(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.CreateOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(staticfabricnetworkattachmentsResource, c.ns, staticFabricNetworkAttachment), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// Update takes the representation of a staticFabricNetworkAttachment and updates it. Returns the server's representation of the staticFabricNetworkAttachment, and an error, if there is any.
func (c *FakeStaticFabricNetworkAttachments) Update(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(staticfabricnetworkattachmentsResource, c.ns, staticFabricNetworkAttachment), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStaticFabricNetworkAttachments) UpdateStatus(ctx context.Context, staticFabricNetworkAttachment *v1.StaticFabricNetworkAttachment, opts metav1.UpdateOptions) (*v1.StaticFabricNetworkAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(staticfabricnetworkattachmentsResource, "status", c.ns, staticFabricNetworkAttachment), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// Delete takes name of the staticFabricNetworkAttachment and deletes it. Returns an error if one occurs.
func (c *FakeStaticFabricNetworkAttachments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(staticfabricnetworkattachmentsResource, c.ns, name, opts), &v1.StaticFabricNetworkAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStaticFabricNetworkAttachments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(staticfabricnetworkattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.StaticFabricNetworkAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched staticFabricNetworkAttachment.
func (c *FakeStaticFabricNetworkAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StaticFabricNetworkAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(staticfabricnetworkattachmentsResource, c.ns, name, pt, data, subresources...), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied staticFabricNetworkAttachment.
func (c *FakeStaticFabricNetworkAttachments) Apply(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	if staticFabricNetworkAttachment == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment provided to Apply must not be nil")
	}
	data, err := json.Marshal(staticFabricNetworkAttachment)
	if err != nil {
		return nil, err
	}
	name := staticFabricNetworkAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(staticfabricnetworkattachmentsResource, c.ns, *name, types.ApplyPatchType, data), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeStaticFabricNetworkAttachments) ApplyStatus(ctx context.Context, staticFabricNetworkAttachment *acifabricattachmentv1.StaticFabricNetworkAttachmentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.StaticFabricNetworkAttachment, err error) {
	if staticFabricNetworkAttachment == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment provided to Apply must not be nil")
	}
	data, err := json.Marshal(staticFabricNetworkAttachment)
	if err != nil {
		return nil, err
	}
	name := staticFabricNetworkAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("staticFabricNetworkAttachment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(staticfabricnetworkattachmentsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.StaticFabricNetworkAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StaticFabricNetworkAttachment), err
}