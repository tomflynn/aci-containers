// Copyright 2020 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	//"net"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"

	erspanpolicy "github.com/noironetworks/aci-containers/pkg/erspanpolicy/apis/aci.erspan/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/noironetworks/aci-containers/pkg/apicapi"
	//"github.com/noironetworks/aci-containers/pkg/ipam"
	tu "github.com/noironetworks/aci-containers/pkg/testutil"
	//"github.com/stretchr/testify/assert"
	//"k8s.io/client-go/tools/cache"
)

func staticErspanReqKey() string {
	return "kube_span_static"
}

type spanTestAugment struct {
	endpoints []*v1.Endpoints
	services  []*v1.Service
}

type spanTest struct {
	erspanPol *erspanpolicy.ErspanPolicy
	aciObj    apicapi.ApicObject
	augment   *spanTestAugment
	desc      string
}

func addErspanServices(cont *testAciController, augment *spanTestAugment) {
	if augment == nil {
		return
	}
	for _, s := range augment.services {
		cont.fakeServiceSource.Add(s)
	}
	for _, e := range augment.endpoints {
		cont.fakeEndpointsSource.Add(e)
	}
}

func makeSpan(source apicapi.ApicSlice, dest apicapi.ApicSlice, srcGrpName string, srcName string,
	destGrpName string, destName string, adminState string, direction string, cep string,
	dstIp string, flowId string) apicapi.ApicObject {

	srcGrp1 := apicapi.NewSpanVSrcGrp(srcGrpName)
	srcGrp1.SetAttr("adminState", adminState)
	src1 := apicapi.NewSpanVSrc(srcGrp1.GetDn(), srcName)
	srcGrp1.AddChild(src1)
	src1.SetAttr("direction", direction)
	//srcCEp1 := apicapi.NewSpanRsSrcToVPort(src1.GetDn(), cep)
	//src.AddChild(srcCEp)

	destGrp1 := apicapi.NewSpanVDestGrp(destGrpName)
	dest1 := apicapi.NewSpanVDest(destGrp1.GetDn(), destName)
	destGrp1.AddChild(dest1)
	destSummary1 := apicapi.NewSpanVEpgSummary(dest1.GetDn())
	dest1.AddChild(destSummary1)
	destSummary1.SetAttr("dstIp", dstIp)
	destSummary1.SetAttr("flowId", flowId)

	return srcGrp1
}

func checkSpan(t *testing.T, st *spanTest, category string, cont *testAciController) {
	tu.WaitFor(t, category+"/"+st.desc, 500*time.Millisecond,
		func(last bool) (bool, error) {
			slice := apicapi.ApicSlice{st.aciObj}
			key := cont.aciNameForKey("span",
				st.erspanPol.Namespace+"_"+st.erspanPol.Name)
			apicapi.PrepareApicSlice(slice, "kube", key)

			if !tu.WaitEqual(t, last, slice,
				cont.apicConn.GetDesiredState(key), st.desc, key) {
				return false, nil
			}
			return true, nil
		})
}

func checkDeleteSpan(t *testing.T, st spanTest, cont *testAciController) {
	tu.WaitFor(t, "delete", 500*time.Millisecond,
		func(last bool) (bool, error) {
			//qr := qrTests[0]
			key := cont.aciNameForKey("span",
				st.erspanPol.Namespace+"_"+st.erspanPol.Name)
			if !tu.WaitEqual(t, last, 0,
				len(cont.apicConn.GetDesiredState(key)), "delete") {
				return false, nil
			}
			return true, nil
		})
}

func testerspanpolicy(name string, namespace string, source erspanpolicy.ErspanPolicingType,
	dest erspanpolicy.ErspanPolicingType, labels map[string]string) *erspanpolicy.ErspanPolicy {
	policy := &erspanpolicy.ErspanPolicy{
		Spec: erspanpolicy.ErspanPolicySpec{
			Source: source, Dest: dest,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: erspanpolicy.ErspanPolicyStatus{
			State: erspanpolicy.Ready,
		},
	}
	return policy
}
