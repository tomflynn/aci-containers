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

// NodeFabricL3PeersStatusApplyConfiguration represents an declarative configuration of the NodeFabricL3PeersStatus type for use
// with apply.
type NodeFabricL3PeersStatusApplyConfiguration struct {
	NADRefs     []NADFabricL3PeerApplyConfiguration            `json:"nadRefs,omitempty"`
	PeeringInfo []NetworkFabricL3PeeringInfoApplyConfiguration `json:"peeringInfo,omitempty"`
}

// NodeFabricL3PeersStatusApplyConfiguration constructs an declarative configuration of the NodeFabricL3PeersStatus type for use with
// apply.
func NodeFabricL3PeersStatus() *NodeFabricL3PeersStatusApplyConfiguration {
	return &NodeFabricL3PeersStatusApplyConfiguration{}
}

// WithNADRefs adds the given value to the NADRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NADRefs field.
func (b *NodeFabricL3PeersStatusApplyConfiguration) WithNADRefs(values ...*NADFabricL3PeerApplyConfiguration) *NodeFabricL3PeersStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNADRefs")
		}
		b.NADRefs = append(b.NADRefs, *values[i])
	}
	return b
}

// WithPeeringInfo adds the given value to the PeeringInfo field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PeeringInfo field.
func (b *NodeFabricL3PeersStatusApplyConfiguration) WithPeeringInfo(values ...*NetworkFabricL3PeeringInfoApplyConfiguration) *NodeFabricL3PeersStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeeringInfo")
		}
		b.PeeringInfo = append(b.PeeringInfo, *values[i])
	}
	return b
}
