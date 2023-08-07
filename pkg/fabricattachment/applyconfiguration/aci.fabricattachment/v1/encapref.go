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

// EncapRefApplyConfiguration represents an declarative configuration of the EncapRef type for use
// with apply.
type EncapRefApplyConfiguration struct {
	NadVlanMapRef *string `json:"nadVlanMap,omitempty"`
	Key           *string `json:"key,omitempty"`
}

// EncapRefApplyConfiguration constructs an declarative configuration of the EncapRef type for use with
// apply.
func EncapRef() *EncapRefApplyConfiguration {
	return &EncapRefApplyConfiguration{}
}

// WithNadVlanMapRef sets the NadVlanMapRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NadVlanMapRef field is set to the value of the last call.
func (b *EncapRefApplyConfiguration) WithNadVlanMapRef(value string) *EncapRefApplyConfiguration {
	b.NadVlanMapRef = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *EncapRefApplyConfiguration) WithKey(value string) *EncapRefApplyConfiguration {
	b.Key = &value
	return b
}
