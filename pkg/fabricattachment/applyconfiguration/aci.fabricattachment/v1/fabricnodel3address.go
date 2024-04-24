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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FabricNodeL3AddressApplyConfiguration represents an declarative configuration of the FabricNodeL3Address type for use
// with apply.
type FabricNodeL3AddressApplyConfiguration struct {
	PrimaryAddress *string                          `json:"primaryAddress,omitempty"`
	NodeRef        *FabricNodeRefApplyConfiguration `json:"nodeRef,omitempty"`
}

// FabricNodeL3AddressApplyConfiguration constructs an declarative configuration of the FabricNodeL3Address type for use with
// apply.
func FabricNodeL3Address() *FabricNodeL3AddressApplyConfiguration {
	return &FabricNodeL3AddressApplyConfiguration{}
}

// WithPrimaryAddress sets the PrimaryAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrimaryAddress field is set to the value of the last call.
func (b *FabricNodeL3AddressApplyConfiguration) WithPrimaryAddress(value string) *FabricNodeL3AddressApplyConfiguration {
	b.PrimaryAddress = &value
	return b
}

// WithNodeRef sets the NodeRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeRef field is set to the value of the last call.
func (b *FabricNodeL3AddressApplyConfiguration) WithNodeRef(value *FabricNodeRefApplyConfiguration) *FabricNodeL3AddressApplyConfiguration {
	b.NodeRef = value
	return b
}
