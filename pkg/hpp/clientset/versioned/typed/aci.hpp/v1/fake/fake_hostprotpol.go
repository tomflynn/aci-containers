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

// FakeHostprotPols implements HostprotPolInterface
type FakeHostprotPols struct {
	Fake *FakeAciV1
	ns   string
}

var hostprotpolsResource = v1.SchemeGroupVersion.WithResource("hostprotpols")

var hostprotpolsKind = v1.SchemeGroupVersion.WithKind("HostprotPol")

// Get takes name of the hostprotPol, and returns the corresponding hostprotPol object, and an error if there is any.
func (c *FakeHostprotPols) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HostprotPol, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostprotpolsResource, c.ns, name), &v1.HostprotPol{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotPol), err
}

// List takes label and field selectors, and returns the list of HostprotPols that match those selectors.
func (c *FakeHostprotPols) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HostprotPolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostprotpolsResource, hostprotpolsKind, c.ns, opts), &v1.HostprotPolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.HostprotPolList{ListMeta: obj.(*v1.HostprotPolList).ListMeta}
	for _, item := range obj.(*v1.HostprotPolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostprotPols.
func (c *FakeHostprotPols) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostprotpolsResource, c.ns, opts))

}

// Create takes the representation of a hostprotPol and creates it.  Returns the server's representation of the hostprotPol, and an error, if there is any.
func (c *FakeHostprotPols) Create(ctx context.Context, hostprotPol *v1.HostprotPol, opts metav1.CreateOptions) (result *v1.HostprotPol, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostprotpolsResource, c.ns, hostprotPol), &v1.HostprotPol{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotPol), err
}

// Update takes the representation of a hostprotPol and updates it. Returns the server's representation of the hostprotPol, and an error, if there is any.
func (c *FakeHostprotPols) Update(ctx context.Context, hostprotPol *v1.HostprotPol, opts metav1.UpdateOptions) (result *v1.HostprotPol, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostprotpolsResource, c.ns, hostprotPol), &v1.HostprotPol{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotPol), err
}

// Delete takes name of the hostprotPol and deletes it. Returns an error if one occurs.
func (c *FakeHostprotPols) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hostprotpolsResource, c.ns, name, opts), &v1.HostprotPol{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostprotPols) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostprotpolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.HostprotPolList{})
	return err
}

// Patch applies the patch and returns the patched hostprotPol.
func (c *FakeHostprotPols) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HostprotPol, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostprotpolsResource, c.ns, name, pt, data, subresources...), &v1.HostprotPol{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotPol), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied hostprotPol.
func (c *FakeHostprotPols) Apply(ctx context.Context, hostprotPol *acihppv1.HostprotPolApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HostprotPol, err error) {
	if hostprotPol == nil {
		return nil, fmt.Errorf("hostprotPol provided to Apply must not be nil")
	}
	data, err := json.Marshal(hostprotPol)
	if err != nil {
		return nil, err
	}
	name := hostprotPol.Name
	if name == nil {
		return nil, fmt.Errorf("hostprotPol.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostprotpolsResource, c.ns, *name, types.ApplyPatchType, data), &v1.HostprotPol{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HostprotPol), err
}
