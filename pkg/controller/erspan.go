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

	erspanpolicy "github.com/noironetworks/aci-containers/pkg/erspanpolicy/apis/aci.erspan/v1"
	erspanclientset "github.com/noironetworks/aci-containers/pkg/erspanpolicy/clientset/versioned"
	podIfclientset "github.com/noironetworksaci-containers/pkg/gbpcrd/clientset/versioned/"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	podIfpolicy "github.com/noironetworks/aci-containers/pkg/gbpcrd/apis/acipolicy/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"github.com/noironetworks/aci-containers/pkg/apicapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	"github.com/noironetworks/aci-containers/pkg/apicapi"
)

const (
	erspanCRDName = "erspanpolicies.aci.erspan"
	pofIfCRDName = "podifpolicies.aci.podif"
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
		listWatch, &erspanpolicy.ErspanPolicy{}, 0,
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

type EndPointData struct {
	MacAddr   string   `json:"macaddr,omitempty"`
	EPG       string   `json:"epg,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	PodName   string   `json:"podname,omitempty"`
}

func (cont *AciController) podIFAdded(obj interface{}) {
	podif, ok := obj.(*podIfpolicy.PodIF)
	if !ok {
		kw.log.Errorf("podIFAdded: Bad object type")
		return
	}

	cont.log.Infof("podIFAdded - %s", podif.ObjectMeta.Name)
	mapOfPodNametoEP := make(map[string]*EndPointData){
		PodName:   podif.Status.PodName,
		MacAddr:   podif.Status.MacAddr,
		EPG:       podif.Status.EPG,
		Namespace: podif.Status.PodNS,
	}
	
}

func (cont *AciController) podIFDeleted(obj interface{}) {
	podif, ok := obj.(*podIfpolicy.PodIF)
	if !ok {
		kw.log.Errorf("podIFDeleted: Bad object type")
		return
	}

	kw.log.Infof("podIFDeleted - %s", podif.ObjectMeta.Name)
	ep := cont.EndPoint{
		MacAddr:   podif.Status.MacAddr,
		EPG:       podif.Status.EPG,
		Namespace: podif.Status.PodNS,
		PodName:   podif.Status.PodName,
	}
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

func (cont *AciController) handleErspanPolUpdate(obj interface{}) bool {
	span, ok := obj.(*erspanpolicy.ErspanPolicy)
	if !ok {
		cont.log.Error("handleErspanPolUpdate: Bad object type")
		return false
	}
	logger := ErspanPolicyLogger(cont.log, span)
	key, err := cache.MetaNamespaceKeyFunc(span)
	if err != nil {
		logger.Error("Could not create erspan policy key: ", err)
		return false
	}
	labelKey := cont.aciNameForKey("span", key)
	cont.log.Debug("create erspanpolicy")
	
	// Generate source policies
	srcGrp := apicapi.NewSpanVSrcGrp(labelKey)
	srcGrpName := srcGrp.GetName()
	apicSlice := apicapi.ApicSlice{srcGrp}
	srcGrp.SetAttr("adminSt", span.Spec.Source.AdminState)
	src := apicapi.NewSpanVSrc(srcGrp.GetDn(), labelKey)
	srcGrp.AddChild(src)
	src.SetAttr("dir", span.Spec.Source.Direction)
	//fvCEpDn := 
	srcCEp := apicapi.NewSpanRsSrcToVPort(src.GetDn())
	src.AddChild(srcCEp)
	
	// Generate destination policies
	destGrp := apicapi.NewSpanVDestGrp(labelKey)
	destGrpName := destGrp.GetName()
	dest := apicapi.NewSpanVDest(destGrp.GetDn(), labelKey)
	destGrp.AddChild(dest)
	destSummary := apicapi.NewSpanVEpgSummary(dest.GetDn())
	dest.AddChild(destSummary)
	destSummary.SetAttr("dstIp", span.Spec.Dest.DstIp)
	destSummary.SetAttr("flowId", strconv.Itoa(span.Spec.Dest.FlowId))
	destSummary.SetAttr("flowId", span.Spec.Dest.FlowId)
	
	srcGrp.GetAttr["name"] = tnSpanVSrcGrpName
	destGrp.GetAttr["name"] = tnSpanVDestGrpName
	
	// Set tag
	lbl := apicapi.NewSpanSpanLbl(srcGrpDn, tnSpanVDestGrpName)
	lbl.SetAttr("tag", span.Spec.Dest.Tag)
	
	//Enable erspan policy on all discovered vpc channels
	vpc := apicapi.NewInfraAccBndlGrp()
	infraRsSpanVSrcGrp := apicapi.NewInfraRsSpanVSrcGrp(vpc.GetDn(), tnSpanVSrcGrpName)
	vpc.AddChild(infraRsSpanVSrcGrp)
	infraRsSpanVDestGrp := apicapi.NewInfraRsSpanVDestGrp(vpc.GetDn(), tnSpanVDestGrpName)
	vpc.AddChild(infraRsSpanVDestGrp)
	return false
	
	cont.apicConn.WriteApicObjects(labelKey, apicSlice)

}


