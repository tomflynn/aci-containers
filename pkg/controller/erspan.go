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
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/noironetworks/aci-containers/pkg/apicapi"
	erspanpolicy "github.com/noironetworks/aci-containers/pkg/erspanpolicy/apis/aci.erspan/v1"
	erspanclientset "github.com/noironetworks/aci-containers/pkg/erspanpolicy/clientset/versioned"
	podIfpolicy "github.com/noironetworks/aci-containers/pkg/gbpcrd/apis/acipolicy/v1"
	podIfclientset "github.com/noironetworks/aci-containers/pkg/gbpcrd/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

const (
	erspanCRDName = "erspanpolicies.aci.erspan"
	pofIfCRDName  = "podifs.aci.aw"
)

type PodSelector struct {
	Labels    map[string]string `json:"labels,omitempt"`
	Namespace string            `json:"namespace,omitempty"`
}

type ErspanPolicySpec struct {
	Selector PodSelector        `json:"selector,omitempty"`
	Source   ErspanPolicingType `json:"source,omitempty"`
	Dest     ErspanPolicingType `json:"destination,omitempty"`
}

type ErspanPolicingType struct {
	AdminState string `json:"admin_state"`
	Direction  string `json:"direction"`
	DestIp     string `json:"destIp"`
	FlowId     int    `json:"flowId"`
	Tag        string `json:"tag,omitempty"`
}

func ErspanPolicyLogger(log *logrus.Logger, erspan *erspanpolicy.ErspanPolicy) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"namespace": erspan.ObjectMeta.Namespace,
		"name":      erspan.ObjectMeta.Name,
		"spec":      erspan.Spec,
	})
}

func PodIfPolicyLogger(log *logrus.Logger, podIf *podIfpolicy.PodIF) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"namespace": podIf.ObjectMeta.Namespace,
		"name":      podIf.ObjectMeta.Name,
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
		listWatch, &podIfpolicy.PodIF{}, 0,
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
		cont.log.Errorf("podIFAdded: Bad object type")
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(podif)
	if err != nil {
		PodIfPolicyLogger(cont.log, podif).
			Error("Could not create key:" + err.Error())
		return
	}

	cont.log.Infof("podIFAdded - %s", podif.ObjectMeta.Name)
	cont.podIftoEp = make(map[string]*EndPointData)
	cont.podIftoEp[podif.Status.PodName] = &EndPointData{
		MacAddr:   podif.Status.MacAddr,
		EPG:       podif.Status.EPG,
		Namespace: podif.Status.PodNS}

	cont.queuePodIfUpdateByKey(key)
	fmt.Println("podIftoEp map podIFAdded:", cont.podIftoEp)

}

func (cont *AciController) podIFDeleted(obj interface{}) {
	podif, ok := obj.(*podIfpolicy.PodIF)
	if !ok {
		cont.log.Errorf("podIFDeleted: Bad object type")
		return
	}
	cont.log.Infof("podIFAdded - %s", podif.ObjectMeta.Name)
	cont.podIftoEp = make(map[string]*EndPointData)
	cont.podIftoEp[podif.Status.PodName] = &EndPointData{
		MacAddr:   podif.Status.MacAddr,
		EPG:       podif.Status.EPG,
		Namespace: podif.Status.PodNS}
	fmt.Println("podIftoEp map podIFDeleted:", cont.podIftoEp)
}

func (cont *AciController) erspanPolicyUpdated(obj interface{}) {
	erspanPolicy := obj.(*erspanpolicy.ErspanPolicy)
	spankey, err := cache.MetaNamespaceKeyFunc(erspanPolicy)
	if err != nil {
		ErspanPolicyLogger(cont.log, erspanPolicy).
			Error("Could not create key:" + err.Error())
		return
	}
	cont.queueErspanUpdateByKey(spankey)

}

func (cont *AciController) queueErspanUpdateByKey(key string) {
	cont.erspanQueue.Add(key)
}

func (cont *AciController) queuePodIfUpdateByKey(key string) {
	cont.podIfQueue.Add(key)
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

// func (cont *AciController) initErspanPolPodIndex() {
// cont.erspanPolPods = index.NewPodSelectorIndex(
// cont.log,
// cont.podIndexer, cont.namespaceIndexer, cont.erspanIndexer,
// cache.MetaNamespaceKeyFunc,
// func(obj interface{}) []index.PodSelector {
// span := obj.(*erspanpolicy.ErspanPolicy)
// ls := &metav1.LabelSelector{MatchLabels: span.Spec.Selector.Labels}
// return index.PodSelectorFromNsAndSelector(span.ObjectMeta.Namespace, ls)
// },
// )
// cont.erspanPolPods.SetPodUpdateCallback(func(podkey string) {
// podobj, exists, err := cont.podIndexer.GetByKey(podkey)
// if exists && err == nil {
// cont.queuePodUpdate(podobj.(*v1.Pod))
// }
// })

// func (cont *AciController) getMacFromPodName(podKeys []string, podName string) []string {
// cont.podIftoEp := make(map[string]*EndPointData)
// for _, podkey := range podKeys {
// podobj, exists, err := cont.podIndexer.GetByKey(podkey)
// if exists && err == nil {
// pod := podobj.(*v1.Pod)
// mac, err := cont.LookupMacByPodName(*pod, podName)
// if err != nil {
// continue
// }
// if _, ok := cont.podIftoEp[string(MacAddr: mac, EPG: epg)];

//Find mac addr by podName in podif database
// func LookupMacByPodName(pod v1.Pod, name string) {
// podif,_ := obj.(*podIfpolicy.PodIF)
// for _, p := range cont.erspanPolPods {
// if p.Name == podif.Status.PodName {
// return podif.Status.MacAddr, nil
// }
// }
// }

func (cont *AciController) writeErspanApic(obj interface{}) bool {
	span, ok := obj.(*erspanpolicy.ErspanPolicy)
	if !ok {
		cont.log.Error("writeErspanApic: Bad object type")
		return false
	}
	// podif, ok := obj.(*podIfpolicy.PodIF)
	// if !ok {
	// cont.log.Errorf("podIF: Bad object type")
	// return false
	// }
	erspanLogger := ErspanPolicyLogger(cont.log, span)
	spankey, err := cache.MetaNamespaceKeyFunc(span)
	if err != nil {
		erspanLogger.Error("Could not create erspan policy key: ", err)
		return false
	}
	labelKey := cont.aciNameForKey("span", spankey)
	cont.log.Debug("create erspanpolicy")

	//podKeys := cont.erspanPolPods.GetPodForObj(spanKey)
	//macs = cont.getMacFromPodName(podKeys, )

	//array of pods that are matched by podSelector
	// matchingPods := []string{}

	// cont.podIftoEp= make(map[string]*EndPointData)
	// for _, PodName := range matchingPods{
	// cont.podIftoEp[PodName] = &EndPointData{MacAddr: podif.Status.MacAddr, EPG: podif.Status.EPG}
	// }

	// for podkey := range matchingPods {
	// if _, ok := matchingPods[podkey]; ok {
	// continue
	// }
	// PodName:= podif.Status.PodName
	//PodName := cont.podIftoEp[podkey]
	// if PodName != nil {
	// podIftoEp[PodName] = &EndPointData{MacAddr: podif.Status.MacAddr, EPG: podif.Status.EPG}

	// }

	// var podIftoEp = map[string]*EndPointData{}
	// podIftoEp["PodName"] = &EndPointData{MacAddr: podif.Status.MacAddr, EPG: podif.Status.EPG}
	// PodName := podIftoEp["PodName"]
	// mac := PodName.MacAddr
	// epg := PodName.EPG

	// Generate source policies
	srcGrp := apicapi.NewSpanVSrcGrp(labelKey)
	srcName := labelKey + "_Src"
	apicSlice := apicapi.ApicSlice{srcGrp}
	srcGrp.SetAttr("adminSt", span.Spec.Source.AdminState)
	src := apicapi.NewSpanVSrc(srcGrp.GetDn(), srcName)
	srcGrp.AddChild(src)
	src.SetAttr("dir", span.Spec.Source.Direction)

	//epg := cont.podIFAdded.PodName.EPG

	//fvCEpDn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s/cep-%s", cont.config.AciPolicyTenant, cont.config.AciAppProfile, epg, mac)
	//srcCEp := apicapi.NewSpanRsSrcToVPort(src.GetDn(), fvCEpDn)
	//src.AddChild(srcCEp)

	// Generate destination policies
	destGrp := apicapi.NewSpanVDestGrp(labelKey)
	destName := labelKey + "_Dest"
	dest := apicapi.NewSpanVDest(destGrp.GetDn(), destName)
	destGrp.AddChild(dest)
	destSummary := apicapi.NewSpanVEpgSummary(dest.GetDn())
	dest.AddChild(destSummary)
	destSummary.SetAttr("dstIp", span.Spec.Dest.DestIp)
	destSummary.SetAttr("flowId", strconv.Itoa(span.Spec.Dest.FlowId))
	apicSlice = append(apicSlice, destGrp)

	// Set tag
	lbl := apicapi.NewSpanSpanLbl(srcGrp.GetDn(), labelKey)
	lbl.SetAttr("tag", span.Spec.Dest.Tag)

	//Enable erspan policy on all discovered vpc channels
	// nMap := make(map[string]string)

	// for device := range cont.nodeOpflexDevice {
	// fabricPath, ok := cont.fabricPathForNode(device)
	// if !ok {
	// continue
	// }
	// nMap[device] = fabricPath
	// }

	// var paths []string
	// fabPath := nMap[device]
	// for device := range nMap {
	// paths = append(paths, fabPath.fabricPath)
	// }
	// var vpcs []string

	// for _, channel := range vpcs {

	// vpc := apicapi.NewInfraAccBndlGrp(channel)
	// infraRsSpanVSrcGrp := apicapi.NewInfraRsSpanVSrcGrp(vpc.GetDn(), labelKey)
	// vpc.AddChild(infraRsSpanVSrcGrp)
	// infraRsSpanVDestGrp := apicapi.NewInfraRsSpanVDestGrp(vpc.GetDn(), labelKey)
	// vpc.AddChild(infraRsSpanVDestGrp)
	// }
	cont.log.Info("creating erspan session", apicSlice)
	cont.apicConn.WriteApicObjects(labelKey, apicSlice)

	return false

}
