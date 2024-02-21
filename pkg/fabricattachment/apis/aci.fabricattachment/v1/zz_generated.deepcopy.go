//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AciNodeLinkAdjacency) DeepCopyInto(out *AciNodeLinkAdjacency) {
	*out = *in
	if in.FabricLink != nil {
		in, out := &in.FabricLink, &out.FabricLink
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]PodAttachment, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AciNodeLinkAdjacency.
func (in *AciNodeLinkAdjacency) DeepCopy() *AciNodeLinkAdjacency {
	if in == nil {
		return nil
	}
	out := new(AciNodeLinkAdjacency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BridgeDomain) DeepCopyInto(out *BridgeDomain) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Vrf = in.Vrf
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BridgeDomain.
func (in *BridgeDomain) DeepCopy() *BridgeDomain {
	if in == nil {
		return nil
	}
	out := new(BridgeDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contracts) DeepCopyInto(out *Contracts) {
	*out = *in
	if in.Consumer != nil {
		in, out := &in.Consumer, &out.Consumer
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contracts.
func (in *Contracts) DeepCopy() *Contracts {
	if in == nil {
		return nil
	}
	out := new(Contracts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncapRef) DeepCopyInto(out *EncapRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncapRef.
func (in *EncapRef) DeepCopy() *EncapRef {
	if in == nil {
		return nil
	}
	out := new(EncapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncapSource) DeepCopyInto(out *EncapSource) {
	*out = *in
	out.EncapRef = in.EncapRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncapSource.
func (in *EncapSource) DeepCopy() *EncapSource {
	if in == nil {
		return nil
	}
	out := new(EncapSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Epg) DeepCopyInto(out *Epg) {
	*out = *in
	in.Contracts.DeepCopyInto(&out.Contracts)
	in.BD.DeepCopyInto(&out.BD)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Epg.
func (in *Epg) DeepCopy() *Epg {
	if in == nil {
		return nil
	}
	out := new(Epg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricVlanPool) DeepCopyInto(out *FabricVlanPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricVlanPool.
func (in *FabricVlanPool) DeepCopy() *FabricVlanPool {
	if in == nil {
		return nil
	}
	out := new(FabricVlanPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricVlanPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricVlanPoolList) DeepCopyInto(out *FabricVlanPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricVlanPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricVlanPoolList.
func (in *FabricVlanPoolList) DeepCopy() *FabricVlanPoolList {
	if in == nil {
		return nil
	}
	out := new(FabricVlanPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricVlanPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricVlanPoolSpec) DeepCopyInto(out *FabricVlanPoolSpec) {
	*out = *in
	if in.Vlans != nil {
		in, out := &in.Vlans, &out.Vlans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricVlanPoolSpec.
func (in *FabricVlanPoolSpec) DeepCopy() *FabricVlanPoolSpec {
	if in == nil {
		return nil
	}
	out := new(FabricVlanPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricVlanPoolStatus) DeepCopyInto(out *FabricVlanPoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricVlanPoolStatus.
func (in *FabricVlanPoolStatus) DeepCopy() *FabricVlanPoolStatus {
	if in == nil {
		return nil
	}
	out := new(FabricVlanPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NADVlanRef) DeepCopyInto(out *NADVlanRef) {
	*out = *in
	if in.Aeps != nil {
		in, out := &in.Aeps, &out.Aeps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NADVlanRef.
func (in *NADVlanRef) DeepCopy() *NADVlanRef {
	if in == nil {
		return nil
	}
	out := new(NADVlanRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NadVlanMap) DeepCopyInto(out *NadVlanMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NadVlanMap.
func (in *NadVlanMap) DeepCopy() *NadVlanMap {
	if in == nil {
		return nil
	}
	out := new(NadVlanMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NadVlanMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NadVlanMapList) DeepCopyInto(out *NadVlanMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NadVlanMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NadVlanMapList.
func (in *NadVlanMapList) DeepCopy() *NadVlanMapList {
	if in == nil {
		return nil
	}
	out := new(NadVlanMapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NadVlanMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NadVlanMapSpec) DeepCopyInto(out *NadVlanMapSpec) {
	*out = *in
	if in.NadVlanMapping != nil {
		in, out := &in.NadVlanMapping, &out.NadVlanMapping
		*out = make(map[string][]VlanSpec, len(*in))
		for key, val := range *in {
			var outVal []VlanSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]VlanSpec, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NadVlanMapSpec.
func (in *NadVlanMapSpec) DeepCopy() *NadVlanMapSpec {
	if in == nil {
		return nil
	}
	out := new(NadVlanMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NadVlanMapStatus) DeepCopyInto(out *NadVlanMapStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NadVlanMapStatus.
func (in *NadVlanMapStatus) DeepCopy() *NadVlanMapStatus {
	if in == nil {
		return nil
	}
	out := new(NadVlanMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFabricConfiguration) DeepCopyInto(out *NetworkFabricConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFabricConfiguration.
func (in *NetworkFabricConfiguration) DeepCopy() *NetworkFabricConfiguration {
	if in == nil {
		return nil
	}
	out := new(NetworkFabricConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFabricConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFabricConfigurationList) DeepCopyInto(out *NetworkFabricConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkFabricConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFabricConfigurationList.
func (in *NetworkFabricConfigurationList) DeepCopy() *NetworkFabricConfigurationList {
	if in == nil {
		return nil
	}
	out := new(NetworkFabricConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFabricConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFabricConfigurationSpec) DeepCopyInto(out *NetworkFabricConfigurationSpec) {
	*out = *in
	if in.VlanRefs != nil {
		in, out := &in.VlanRefs, &out.VlanRefs
		*out = make([]VlanRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NADVlanRefs != nil {
		in, out := &in.NADVlanRefs, &out.NADVlanRefs
		*out = make([]NADVlanRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFabricConfigurationSpec.
func (in *NetworkFabricConfigurationSpec) DeepCopy() *NetworkFabricConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkFabricConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFabricConfigurationStatus) DeepCopyInto(out *NetworkFabricConfigurationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFabricConfigurationStatus.
func (in *NetworkFabricConfigurationStatus) DeepCopy() *NetworkFabricConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkFabricConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFabricNetworkAttachment) DeepCopyInto(out *NodeFabricNetworkAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFabricNetworkAttachment.
func (in *NodeFabricNetworkAttachment) DeepCopy() *NodeFabricNetworkAttachment {
	if in == nil {
		return nil
	}
	out := new(NodeFabricNetworkAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFabricNetworkAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFabricNetworkAttachmentList) DeepCopyInto(out *NodeFabricNetworkAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFabricNetworkAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFabricNetworkAttachmentList.
func (in *NodeFabricNetworkAttachmentList) DeepCopy() *NodeFabricNetworkAttachmentList {
	if in == nil {
		return nil
	}
	out := new(NodeFabricNetworkAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFabricNetworkAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFabricNetworkAttachmentSpec) DeepCopyInto(out *NodeFabricNetworkAttachmentSpec) {
	*out = *in
	out.NetworkRef = in.NetworkRef
	out.EncapVlan = in.EncapVlan
	if in.AciTopology != nil {
		in, out := &in.AciTopology, &out.AciTopology
		*out = make(map[string]AciNodeLinkAdjacency, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFabricNetworkAttachmentSpec.
func (in *NodeFabricNetworkAttachmentSpec) DeepCopy() *NodeFabricNetworkAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFabricNetworkAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFabricNetworkAttachmentStatus) DeepCopyInto(out *NodeFabricNetworkAttachmentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFabricNetworkAttachmentStatus.
func (in *NodeFabricNetworkAttachmentStatus) DeepCopy() *NodeFabricNetworkAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(NodeFabricNetworkAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjRef) DeepCopyInto(out *ObjRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjRef.
func (in *ObjRef) DeepCopy() *ObjRef {
	if in == nil {
		return nil
	}
	out := new(ObjRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAttachment) DeepCopyInto(out *PodAttachment) {
	*out = *in
	out.PodRef = in.PodRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAttachment.
func (in *PodAttachment) DeepCopy() *PodAttachment {
	if in == nil {
		return nil
	}
	out := new(PodAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRF) DeepCopyInto(out *VRF) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRF.
func (in *VRF) DeepCopy() *VRF {
	if in == nil {
		return nil
	}
	out := new(VRF)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanRef) DeepCopyInto(out *VlanRef) {
	*out = *in
	if in.Aeps != nil {
		in, out := &in.Aeps, &out.Aeps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Epg.DeepCopyInto(&out.Epg)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanRef.
func (in *VlanRef) DeepCopy() *VlanRef {
	if in == nil {
		return nil
	}
	out := new(VlanRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpec) DeepCopyInto(out *VlanSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpec.
func (in *VlanSpec) DeepCopy() *VlanSpec {
	if in == nil {
		return nil
	}
	out := new(VlanSpec)
	in.DeepCopyInto(out)
	return out
}
