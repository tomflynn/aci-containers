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

	v1 "github.com/noironetworks/aci-containers/pkg/hpp/apis/aci.hpp/v1"
	acihppv1 "github.com/noironetworks/aci-containers/pkg/hpp/applyconfiguration/aci.hpp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostprotRemoteIpContainers implements HostprotRemoteIpContainerInterface
type FakeHostprotRemoteIpContainers struct {
	Fake *FakeAciV1
	ns   string
}

var hostprotremoteipcontainersResource = v1.SchemeGroupVersion.WithResource("hostprotremoteipcontainers")

var hostprotremoteipcontainersKind = v1.SchemeGroupVersion.WithKind("HostprotRemoteIpContainer")

// Get takes name of the hostprotRemoteIpContainer, and returns the corresponding hostprotRemoteIpContainer object, and an error if there is any.
func (c *FakeHostprotRemoteIpContainers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HostprotRemoteIpContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostprotremoteipcontainersResource, c.ns, name), &v1.HostprotRemoteIpContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotRemoteIpContainer), err
}

// List takes label and field selectors, and returns the list of HostprotRemoteIpContainers that match those selectors.
func (c *FakeHostprotRemoteIpContainers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HostprotRemoteIpContainerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostprotremoteipcontainersResource, hostprotremoteipcontainersKind, c.ns, opts), &v1.HostprotRemoteIpContainerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.HostprotRemoteIpContainerList{ListMeta: obj.(*v1.HostprotRemoteIpContainerList).ListMeta}
	for _, item := range obj.(*v1.HostprotRemoteIpContainerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostprotRemoteIpContainers.
func (c *FakeHostprotRemoteIpContainers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostprotremoteipcontainersResource, c.ns, opts))

}

// Create takes the representation of a hostprotRemoteIpContainer and creates it.  Returns the server's representation of the hostprotRemoteIpContainer, and an error, if there is any.
func (c *FakeHostprotRemoteIpContainers) Create(ctx context.Context, hostprotRemoteIpContainer *v1.HostprotRemoteIpContainer, opts metav1.CreateOptions) (result *v1.HostprotRemoteIpContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostprotremoteipcontainersResource, c.ns, hostprotRemoteIpContainer), &v1.HostprotRemoteIpContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotRemoteIpContainer), err
}

// Update takes the representation of a hostprotRemoteIpContainer and updates it. Returns the server's representation of the hostprotRemoteIpContainer, and an error, if there is any.
func (c *FakeHostprotRemoteIpContainers) Update(ctx context.Context, hostprotRemoteIpContainer *v1.HostprotRemoteIpContainer, opts metav1.UpdateOptions) (result *v1.HostprotRemoteIpContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostprotremoteipcontainersResource, c.ns, hostprotRemoteIpContainer), &v1.HostprotRemoteIpContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotRemoteIpContainer), err
}

// Delete takes name of the hostprotRemoteIpContainer and deletes it. Returns an error if one occurs.
func (c *FakeHostprotRemoteIpContainers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hostprotremoteipcontainersResource, c.ns, name, opts), &v1.HostprotRemoteIpContainer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostprotRemoteIpContainers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostprotremoteipcontainersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.HostprotRemoteIpContainerList{})
	return err
}

// Patch applies the patch and returns the patched hostprotRemoteIpContainer.
func (c *FakeHostprotRemoteIpContainers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HostprotRemoteIpContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostprotremoteipcontainersResource, c.ns, name, pt, data, subresources...), &v1.HostprotRemoteIpContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotRemoteIpContainer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied hostprotRemoteIpContainer.
func (c *FakeHostprotRemoteIpContainers) Apply(ctx context.Context, hostprotRemoteIpContainer *acihppv1.HostprotRemoteIpContainerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HostprotRemoteIpContainer, err error) {
	if hostprotRemoteIpContainer == nil {
		return nil, fmt.Errorf("hostprotRemoteIpContainer provided to Apply must not be nil")
	}
	data, err := json.Marshal(hostprotRemoteIpContainer)
	if err != nil {
		return nil, err
	}
	name := hostprotRemoteIpContainer.Name
	if name == nil {
		return nil, fmt.Errorf("hostprotRemoteIpContainer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostprotremoteipcontainersResource, c.ns, *name, types.ApplyPatchType, data), &v1.HostprotRemoteIpContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotRemoteIpContainer), err
}
