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

// BFDRefApplyConfiguration represents an declarative configuration of the BFDRef type for use
// with apply.
type BFDRefApplyConfiguration struct {
	Enabled   *bool                     `json:"enabled,omitempty"`
	BFDIfPRef *ObjRefApplyConfiguration `json:"bfdIfPRef,omitempty"`
}

// BFDRefApplyConfiguration constructs an declarative configuration of the BFDRef type for use with
// apply.
func BFDRef() *BFDRefApplyConfiguration {
	return &BFDRefApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *BFDRefApplyConfiguration) WithEnabled(value bool) *BFDRefApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithBFDIfPRef sets the BFDIfPRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BFDIfPRef field is set to the value of the last call.
func (b *BFDRefApplyConfiguration) WithBFDIfPRef(value *ObjRefApplyConfiguration) *BFDRefApplyConfiguration {
	b.BFDIfPRef = value
	return b
}
