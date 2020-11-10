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
	"github.com/sirupsen/logrus"

	"github.com/noironetworks/aci-containers/pkg/apicapi"
	erspanpolicy "github.com/noironetworks/aci-containers/pkg/erspanpolicy/apis/aci.erspan/v1"
	erspanclientset "github.com/noironetworks/aci-containers/pkg/erspanpolicy/clientset/versioned"
	podIfpolicy "github.com/noironetworks/aci-containers/pkg/gbpcrd/apis/acipolicy/v1"
	podIfclientset "github.com/noironetworksaci-containers/pkg/gbpcrd/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

const (
	erspanCRDName = "erspanpolicies.aci.erspan"
	pofIfCRDName  = "podifpolicies.aci.podif"
)

func ErspanPolicyLogger(log *logrus.Logger, erspan *erspanpolicy.ErspanPolicy) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"namespace": erspan.ObjectMeta.Namespace,
		"name":      erspan.ObjectMeta.Name,
		"spec":      erspan.Spec,
	})
}

func erspanInit(cont *AciController, stopCh <-chan struct{}) {
	cont.log.Debug("Initializing erspan client")
	restconfig := cont.env.RESTConfig()
	erspanClient, err := erspanclientset.NewForConfig(restconfig)
	if err != nil {
		cont.log.Errorf("Failed to intialize erspan client")
		return
	}
	cont.initErspanInformerFromClient(erspanClient)
	cont.erspanInformer.Run(stopCh)
}

func podIfInit(cont *AciController, stopCh <-chan struct{}) {
	cont.log.Debug("Initializing podIf client")
	restconfig := cont.env.RESTConfig()
	podIfClient, err := podIfclientset.NewForConfig(restconfig)
	if err != nil {
		cont.log.Errorf("Failed to intialize podIf client")
		return
	}
	cont.initPodIfInformerFromClient(podIfClient)
	cont.podIfInformer.Run(stopCh)
}

func (cont *AciController) initErspanInformerFromClient(
	erspanClient *erspanclientset.Clientset) {
	cont.initErspanInformerBase(
		cache.NewListWatchFromClient(
			erspanClient.AciV1().RESTClient(), "erspanpolicies",
			metav1.NamespaceAll, fields.Everything()))
}

func (cont *AciController) initPodIfInformerFromClient(
	podIfClient *podIfclientset.Clientset) {
	cont.initPodIfInformerBase(
		cache.NewListWatchFromClient(
			podIfClient.AciV1().RESTClient(), "podIfpolicies",
			metav1.NamespaceAll, fields.Everything()))
}

func (cont *AciController) initErspanInformerBase(listWatch *cache.ListWatch) {
	cont.erspanIndexer, cont.erspanInformer = cache.NewIndexerInformer(
		listWatch,
		&erspanpolicy.ErspanPolicy{}, 0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				cont.erspanPolicyUpdated(obj)
			},
			UpdateFunc: func(_, obj interface{}) {
				cont.erspanPolicyUpdated(obj)
			},
			DeleteFunc: func(obj interface{}) {
				cont.erspanPolicyDelete(obj)
			},
		},
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	cont.log.Debug("Initializing Erspan Policy Informers")
}

func (cont *AciController) initPodIfInformerBase(listWatch *cache.ListWatch) {
	cont.podIfIndexer, cont.podIfInformer = cache.NewIndexerInformer(
		listWatch, &aciv1.PodIF{}, 0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				cont.podIFAdded(obj)
			},
			UpdateFunc: func(oldobj interface{}, newobj interface{}) {
				cont.podIFAdded(newobj)
			},
			DeleteFunc: func(obj interface{}) {
				cont.podIFDeleted(obj)
			},
		},
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	cont.log.Debug("Initializing podif Informers")
}

func (cont *AciController) podIFAdded(obj interface{}) {
	podif, ok := obj.(*podIfpolicy.PodIF)
	if !ok {
		kw.log.Errorf("podIFAdded: Bad object type")
		return
	}

	cont.log.Infof("podIFAdded - %s", podif.ObjectMeta.Name)

}

func (cont *AciController) podIFDeleted(obj interface{}) {
	podif, ok := obj.(*podIfpolicy.PodIF)
	if !ok {
		kw.log.Errorf("podIFDeleted: Bad object type")
		return
	}

	cont.log.Infof("podIFDeleted - %s", podif.ObjectMeta.Name)
}

func (cont *AciController) erspanPolicyUpdated(obj interface{}) {
	erspanPolicy := obj.(*erspanpolicy.ErspanPolicy)
	key, err := cache.MetaNamespaceKeyFunc(erspanPolicy)
	if err != nil {
		ErspanPolicyLogger(cont.log, erspanPolicy).
			Error("Could not create key:" + err.Error())
		return
	}
	cont.queueErspanUpdateByKey(key)

}

func (cont *AciController) queueErspanUpdateByKey(key string) {
	cont.erspanQueue.Add(key)
}

func (cont *AciController) erspanPolicyDelete(obj interface{}) {
	span, isSpan := obj.(*erspanpolicy.ErspanPolicy)
	if !isSpan {
		deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			ErspanPolicyLogger(cont.log, span).
				Error("Received unexpected object: ", obj)
			return
		}
		span, ok = deletedState.Obj.(*erspanpolicy.ErspanPolicy)
		if !ok {
			ErspanPolicyLogger(cont.log, span).
				Error("DeletedFinalStateUnknown contained non-erspan object: ", deletedState.Obj)
			return
		}
	}
	spankey, err := cache.MetaNamespaceKeyFunc(span)
	if err != nil {
		ErspanPolicyLogger(cont.log, span).
			Error("Could not create erspan key: ", err)
		return
	}
	cont.apicConn.ClearApicObjects(cont.aciNameForKey("span", spankey))

}

type EndPointData struct {
	MacAddr string
	EPG     string
}

func (cont *AciController) handleErspanPolUpdate(obj interface{}) bool {
	span, ok := obj.(*erspanpolicy.ErspanPolicy)
	if !ok {
		cont.log.Error("handleErspanPolUpdate: Bad object type")
		return false
	}
	podif, ok := obj.(*aciv1.PodIF)
	if !ok {
		kw.log.Errorf("podIFAdded: Bad object type")
		return
	}
	logger := ErspanPolicyLogger(cont.log, span)
	key, err := cache.MetaNamespaceKeyFunc(span)
	if err != nil {
		logger.Error("Could not create erspan policy key: ", err)
		return false
	}
	labelKey := cont.aciNameForKey("span", key)
	cont.log.Debug("create erspanpolicy")

	var podIftoEp = map[string]*EndPointData{}
	podIftoEp["PodName"] = &EndPointData{MacAddr: podif.Status.MacAddr, EPG: podif.Status.EPG}
	PodName := podIftoEp["PodName"]
	mac := PodName.MacAddr
	epg := PodName.EPG

	// Generate source policies
	srcGrp := apicapi.NewSpanVSrcGrp(labelKey)
	srcName := labelKey + "_Src"
	apicSlice := apicapi.ApicSlice{srcGrp}
	srcGrp.SetAttr("adminSt", span.Spec.Source.AdminState)
	src := apicapi.NewSpanVSrc(srcGrp.GetDn(), srcName)
	srcGrp.AddChild(src)
	src.SetAttr("dir", span.Spec.Source.Direction)

	//epg := cont.podIFAdded.PodName.EPG

	fvCEpDn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s/cep-%s", cont.config.AciPolicyTenant, cont.config.AciAppProfile, epg, mac)
	srcCEp := apicapi.NewSpanRsSrcToVPort(src.GetDn(), fvCEpDn)
	src.AddChild(srcCEp)

	// Generate destination policies
	destGrp := apicapi.NewSpanVDestGrp(labelKey)
	destName := labelKey + "_Dest"
	dest := apicapi.NewSpanVDest(destGrp.GetDn(), destName)
	destGrp.AddChild(dest)
	destSummary := apicapi.NewSpanVEpgSummary(dest.GetDn())
	dest.AddChild(destSummary)
	destSummary.SetAttr("dstIp", span.Spec.Dest.DstIp)
	destSummary.SetAttr("flowId", strconv.Itoa(span.Spec.Dest.FlowId))
	apicSlice = append(apicSlice, destGrp)

	// Set tag
	lbl := apicapi.NewSpanSpanLbl(srcGrpDn, labelKey)
	lbl.SetAttr("tag", span.Spec.Dest.Tag)

	//Enable erspan policy on all discovered vpc channels
	nMap := make(map[string]string)

	for device := range cont.nodeOpflexDevice {
		fabricPath, ok := cont.fabricPathForNode(device)
		if !ok {
			continue
		}
		nMap[device] = fabricPath
	}

	var paths []string
	fabPath := nMap[device]
	for device := range nMap {
		paths = append(paths, fabPath.fabricPath)
	}
	var vpcs []string

	for _, channel := range vpcs {

		vpc := apicapi.NewInfraAccBndlGrp(channel)
		infraRsSpanVSrcGrp := apicapi.NewInfraRsSpanVSrcGrp(vpc.GetDn(), labelKey)
		vpc.AddChild(infraRsSpanVSrcGrp)
		infraRsSpanVDestGrp := apicapi.NewInfraRsSpanVDestGrp(vpc.GetDn(), labelKey)
		vpc.AddChild(infraRsSpanVDestGrp)
	}

	cont.log.Info("creating erspan session", apicSlice)
	cont.apicConn.WriteApicObjects(labelKey, apicSlice)

	return false

}
