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

// BGPPeerPrefixPolicyStatusApplyConfiguration represents an declarative configuration of the BGPPeerPrefixPolicyStatus type for use
// with apply.
type BGPPeerPrefixPolicyStatusApplyConfiguration struct {
	BGPPeerPrefixPolicyApplyConfiguration `json:",inline"`
	Status                                *string `json:"status,omitempty"`
}

// BGPPeerPrefixPolicyStatusApplyConfiguration constructs an declarative configuration of the BGPPeerPrefixPolicyStatus type for use with
// apply.
func BGPPeerPrefixPolicyStatus() *BGPPeerPrefixPolicyStatusApplyConfiguration {
	return &BGPPeerPrefixPolicyStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BGPPeerPrefixPolicyStatusApplyConfiguration) WithName(value string) *BGPPeerPrefixPolicyStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithMaxPrefixes sets the MaxPrefixes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxPrefixes field is set to the value of the last call.
func (b *BGPPeerPrefixPolicyStatusApplyConfiguration) WithMaxPrefixes(value int) *BGPPeerPrefixPolicyStatusApplyConfiguration {
	b.MaxPrefixes = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *BGPPeerPrefixPolicyStatusApplyConfiguration) WithAction(value string) *BGPPeerPrefixPolicyStatusApplyConfiguration {
	b.Action = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *BGPPeerPrefixPolicyStatusApplyConfiguration) WithStatus(value string) *BGPPeerPrefixPolicyStatusApplyConfiguration {
	b.Status = &value
	return b
}
