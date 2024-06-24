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

package applyconfiguration

import (
	v1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/apis/aci.fabricattachment/v1"
	acifabricattachmentv1 "github.com/noironetworks/aci-containers/pkg/fabricattachment/applyconfiguration/aci.fabricattachment/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=aci.fabricattachment, Version=v1
	case v1.SchemeGroupVersion.WithKind("AciNodeLinkAdjacency"):
		return &acifabricattachmentv1.AciNodeLinkAdjacencyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BGPPeerPolicy"):
		return &acifabricattachmentv1.BGPPeerPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BGPPeerPrefixPolicy"):
		return &acifabricattachmentv1.BGPPeerPrefixPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BGPPeerPrefixPolicyStatus"):
		return &acifabricattachmentv1.BGPPeerPrefixPolicyStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BridgeDomain"):
		return &acifabricattachmentv1.BridgeDomainApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConnectedL3Network"):
		return &acifabricattachmentv1.ConnectedL3NetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConnectedL3NetworkStatus"):
		return &acifabricattachmentv1.ConnectedL3NetworkStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Contracts"):
		return &acifabricattachmentv1.ContractsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EncapRef"):
		return &acifabricattachmentv1.EncapRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EncapSource"):
		return &acifabricattachmentv1.EncapSourceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Epg"):
		return &acifabricattachmentv1.EpgApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3Network"):
		return &acifabricattachmentv1.FabricL3NetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3Out"):
		return &acifabricattachmentv1.FabricL3OutApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3OutNextHop"):
		return &acifabricattachmentv1.FabricL3OutNextHopApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3OutNode"):
		return &acifabricattachmentv1.FabricL3OutNodeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3OutRtrNode"):
		return &acifabricattachmentv1.FabricL3OutRtrNodeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3OutStaticRoute"):
		return &acifabricattachmentv1.FabricL3OutStaticRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3OutStatus"):
		return &acifabricattachmentv1.FabricL3OutStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3Peers"):
		return &acifabricattachmentv1.FabricL3PeersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricL3Subnet"):
		return &acifabricattachmentv1.FabricL3SubnetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricNodeRef"):
		return &acifabricattachmentv1.FabricNodeRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricPodRef"):
		return &acifabricattachmentv1.FabricPodRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricTenantConfiguration"):
		return &acifabricattachmentv1.FabricTenantConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricTenantConfigurationStatus"):
		return &acifabricattachmentv1.FabricTenantConfigurationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricVlanPool"):
		return &acifabricattachmentv1.FabricVlanPoolApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricVlanPoolSpec"):
		return &acifabricattachmentv1.FabricVlanPoolSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricVlanPoolStatus"):
		return &acifabricattachmentv1.FabricVlanPoolStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricVrfConfiguration"):
		return &acifabricattachmentv1.FabricVrfConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FabricVrfConfigurationStatus"):
		return &acifabricattachmentv1.FabricVrfConfigurationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NADFabricL3Peer"):
		return &acifabricattachmentv1.NADFabricL3PeerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NadVlanMap"):
		return &acifabricattachmentv1.NadVlanMapApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NadVlanMapSpec"):
		return &acifabricattachmentv1.NadVlanMapSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NadVlanMapStatus"):
		return &acifabricattachmentv1.NadVlanMapStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NADVlanRef"):
		return &acifabricattachmentv1.NADVlanRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricConfiguration"):
		return &acifabricattachmentv1.NetworkFabricConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricConfigurationSpec"):
		return &acifabricattachmentv1.NetworkFabricConfigurationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricConfigurationStatus"):
		return &acifabricattachmentv1.NetworkFabricConfigurationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricL3ConfigSpec"):
		return &acifabricattachmentv1.NetworkFabricL3ConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricL3ConfigStatus"):
		return &acifabricattachmentv1.NetworkFabricL3ConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricL3Configuration"):
		return &acifabricattachmentv1.NetworkFabricL3ConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkFabricL3PeeringInfo"):
		return &acifabricattachmentv1.NetworkFabricL3PeeringInfoApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricL3Peer"):
		return &acifabricattachmentv1.NodeFabricL3PeerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricL3Peers"):
		return &acifabricattachmentv1.NodeFabricL3PeersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricL3PeersStatus"):
		return &acifabricattachmentv1.NodeFabricL3PeersStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricNetworkAttachment"):
		return &acifabricattachmentv1.NodeFabricNetworkAttachmentApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricNetworkAttachmentSpec"):
		return &acifabricattachmentv1.NodeFabricNetworkAttachmentSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeFabricNetworkAttachmentStatus"):
		return &acifabricattachmentv1.NodeFabricNetworkAttachmentStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjRef"):
		return &acifabricattachmentv1.ObjRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodAttachment"):
		return &acifabricattachmentv1.PodAttachmentApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyPrefix"):
		return &acifabricattachmentv1.PolicyPrefixApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyPrefixGroup"):
		return &acifabricattachmentv1.PolicyPrefixGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrimaryNetwork"):
		return &acifabricattachmentv1.PrimaryNetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VlanRef"):
		return &acifabricattachmentv1.VlanRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VlanSpec"):
		return &acifabricattachmentv1.VlanSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VRF"):
		return &acifabricattachmentv1.VRFApplyConfiguration{}

	}
	return nil
}
