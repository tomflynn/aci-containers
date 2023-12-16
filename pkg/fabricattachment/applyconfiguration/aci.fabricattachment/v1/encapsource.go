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

import (
	acifabricattachmentv1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/apis/aci.fabricattachment/v1"
)

// EncapSourceApplyConfiguration represents an declarative configuration of the EncapSource type for use
// with apply.
type EncapSourceApplyConfiguration struct {
	VlanList *string                          `json:"vlanList,omitempty"`
	EncapRef *EncapRefApplyConfiguration      `json:"encapRef,omitempty"`
	Mode     *acifabricattachmentv1.EncapMode `json:"mode,omitempty"`
}

// EncapSourceApplyConfiguration constructs an declarative configuration of the EncapSource type for use with
// apply.
func EncapSource() *EncapSourceApplyConfiguration {
	return &EncapSourceApplyConfiguration{}
}

// WithVlanList sets the VlanList field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VlanList field is set to the value of the last call.
func (b *EncapSourceApplyConfiguration) WithVlanList(value string) *EncapSourceApplyConfiguration {
	b.VlanList = &value
	return b
}

// WithEncapRef sets the EncapRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EncapRef field is set to the value of the last call.
func (b *EncapSourceApplyConfiguration) WithEncapRef(value *EncapRefApplyConfiguration) *EncapSourceApplyConfiguration {
	b.EncapRef = value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *EncapSourceApplyConfiguration) WithMode(value acifabricattachmentv1.EncapMode) *EncapSourceApplyConfiguration {
	b.Mode = &value
	return b
}
