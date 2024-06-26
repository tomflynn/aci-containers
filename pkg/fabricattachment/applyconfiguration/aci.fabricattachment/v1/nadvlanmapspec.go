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

// NadVlanMapSpecApplyConfiguration represents an declarative configuration of the NadVlanMapSpec type for use
// with apply.
type NadVlanMapSpecApplyConfiguration struct {
	NadVlanMapping map[string][]v1.VlanSpec `json:"nadVlanMapping,omitempty"`
}

// NadVlanMapSpecApplyConfiguration constructs an declarative configuration of the NadVlanMapSpec type for use with
// apply.
func NadVlanMapSpec() *NadVlanMapSpecApplyConfiguration {
	return &NadVlanMapSpecApplyConfiguration{}
}

// WithNadVlanMapping puts the entries into the NadVlanMapping field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NadVlanMapping field,
// overwriting an existing map entries in NadVlanMapping field with the same key.
func (b *NadVlanMapSpecApplyConfiguration) WithNadVlanMapping(entries map[string][]v1.VlanSpec) *NadVlanMapSpecApplyConfiguration {
	if b.NadVlanMapping == nil && len(entries) > 0 {
		b.NadVlanMapping = make(map[string][]v1.VlanSpec, len(entries))
	}
	for k, v := range entries {
		b.NadVlanMapping[k] = v
	}
	return b
}
