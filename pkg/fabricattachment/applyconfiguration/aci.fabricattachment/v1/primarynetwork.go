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
	v1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/apis/aci.fabricattachment/v1"
)

// PrimaryNetworkApplyConfiguration represents an declarative configuration of the PrimaryNetwork type for use
// with apply.
type PrimaryNetworkApplyConfiguration struct {
	L3OutName           *string                          `json:"l3OutName,omitempty"`
	L3OutOnCommonTenant *bool                            `json:"l3OutOnCommonTenant,omitempty"`
	UseExistingL3Out    *bool                            `json:"useExistingL3Out,omitempty"`
	MaxNodes            *int                             `json:"maxNodes,omitempty"`
	Encap               *int                             `json:"encap,omitempty"`
	SviType             *v1.FabricSviType                `json:"sviType,omitempty"`
	PrimarySubnet       *string                          `json:"primarySubnet,omitempty"`
	BGPPeerPolicy       *BGPPeerPolicyApplyConfiguration `json:"bgpPeerPolicy,omitempty"`
}

// PrimaryNetworkApplyConfiguration constructs an declarative configuration of the PrimaryNetwork type for use with
// apply.
func PrimaryNetwork() *PrimaryNetworkApplyConfiguration {
	return &PrimaryNetworkApplyConfiguration{}
}

// WithL3OutName sets the L3OutName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the L3OutName field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithL3OutName(value string) *PrimaryNetworkApplyConfiguration {
	b.L3OutName = &value
	return b
}

// WithL3OutOnCommonTenant sets the L3OutOnCommonTenant field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the L3OutOnCommonTenant field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithL3OutOnCommonTenant(value bool) *PrimaryNetworkApplyConfiguration {
	b.L3OutOnCommonTenant = &value
	return b
}

// WithUseExistingL3Out sets the UseExistingL3Out field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseExistingL3Out field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithUseExistingL3Out(value bool) *PrimaryNetworkApplyConfiguration {
	b.UseExistingL3Out = &value
	return b
}

// WithMaxNodes sets the MaxNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxNodes field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithMaxNodes(value int) *PrimaryNetworkApplyConfiguration {
	b.MaxNodes = &value
	return b
}

// WithEncap sets the Encap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Encap field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithEncap(value int) *PrimaryNetworkApplyConfiguration {
	b.Encap = &value
	return b
}

// WithSviType sets the SviType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SviType field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithSviType(value v1.FabricSviType) *PrimaryNetworkApplyConfiguration {
	b.SviType = &value
	return b
}

// WithPrimarySubnet sets the PrimarySubnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrimarySubnet field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithPrimarySubnet(value string) *PrimaryNetworkApplyConfiguration {
	b.PrimarySubnet = &value
	return b
}

// WithBGPPeerPolicy sets the BGPPeerPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BGPPeerPolicy field is set to the value of the last call.
func (b *PrimaryNetworkApplyConfiguration) WithBGPPeerPolicy(value *BGPPeerPolicyApplyConfiguration) *PrimaryNetworkApplyConfiguration {
	b.BGPPeerPolicy = value
	return b
}