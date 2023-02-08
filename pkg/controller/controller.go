// Copyright 2017 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yl2chen/cidranger"
	"golang.org/x/time/rate"

	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/discovery/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/noironetworks/aci-containers/pkg/apicapi"
	"github.com/noironetworks/aci-containers/pkg/index"
	"github.com/noironetworks/aci-containers/pkg/ipam"
	istiov1 "github.com/noironetworks/aci-containers/pkg/istiocrd/apis/aci.istio/v1"
	"github.com/noironetworks/aci-containers/pkg/metadata"
	nodeinfo "github.com/noironetworks/aci-containers/pkg/nodeinfo/apis/aci.snat/v1"
	rdConfig "github.com/noironetworks/aci-containers/pkg/rdconfig/apis/aci.snat/v1"
	snatglobalinfo "github.com/noironetworks/aci-containers/pkg/snatglobalinfo/apis/aci.snat/v1"
	"github.com/noironetworks/aci-containers/pkg/util"
	"k8s.io/client-go/kubernetes"
)

type podUpdateFunc func(*v1.Pod) (*v1.Pod, error)
type nodeUpdateFunc func(*v1.Node) (*v1.Node, error)
type serviceUpdateFunc func(*v1.Service) (*v1.Service, error)

type AciController struct {
	log    *logrus.Logger
	config *ControllerConfig
	env    Environment

	defaultEg string
	defaultSg string

	unitTestMode bool

	podQueue          workqueue.RateLimitingInterface
	netPolQueue       workqueue.RateLimitingInterface
	qosQueue          workqueue.RateLimitingInterface
	serviceQueue      workqueue.RateLimitingInterface
	snatQueue         workqueue.RateLimitingInterface
	netflowQueue      workqueue.RateLimitingInterface
	erspanQueue       workqueue.RateLimitingInterface
	snatNodeInfoQueue workqueue.RateLimitingInterface
	rdConfigQueue     workqueue.RateLimitingInterface
	istioQueue        workqueue.RateLimitingInterface

	namespaceIndexer      cache.Indexer
	namespaceInformer     cache.Controller
	podIndexer            cache.Indexer
	podInformer           cache.Controller
	endpointsIndexer      cache.Indexer
	endpointsInformer     cache.Controller
	serviceIndexer        cache.Indexer
	serviceInformer       cache.Controller
	replicaSetIndexer     cache.Indexer
	replicaSetInformer    cache.Controller
	deploymentIndexer     cache.Indexer
	deploymentInformer    cache.Controller
	nodeIndexer           cache.Indexer
	nodeInformer          cache.Controller
	networkPolicyIndexer  cache.Indexer
	networkPolicyInformer cache.Controller
	snatIndexer           cache.Indexer
	snatInformer          cache.Controller
	snatNodeInfoIndexer   cache.Indexer
	snatNodeInformer      cache.Controller
	crdInformer           cache.Controller
	rdConfigInformer      cache.Controller
	rdConfigIndexer       cache.Indexer
	qosIndexer            cache.Indexer
	qosInformer           cache.Controller
	netflowIndexer        cache.Indexer
	netflowInformer       cache.Controller
	erspanIndexer         cache.Indexer
	erspanInformer        cache.Controller
	nodePodIfIndexer      cache.Indexer
	nodePodIfInformer     cache.Controller
	istioIndexer          cache.Indexer
	istioInformer         cache.Controller
	endpointSliceIndexer  cache.Indexer
	endpointSliceInformer cache.Controller
	snatCfgInformer       cache.Controller
	updatePod             podUpdateFunc
	updateNode            nodeUpdateFunc
	updateServiceStatus   serviceUpdateFunc

	indexMutex sync.Mutex

	configuredPodNetworkIps *netIps
	podNetworkIps           *netIps
	serviceIps              *ipam.IpCache
	staticServiceIps        *netIps
	nodeServiceIps          *netIps

	// index of pods matched by deployments
	depPods *index.PodSelectorIndex
	// index of pods matched by network policies
	netPolPods *index.PodSelectorIndex
	// index of pods matched by network policy ingress rules
	netPolIngressPods *index.PodSelectorIndex
	// index of pods matched by network policy egress rules
	netPolEgressPods *index.PodSelectorIndex
	// index of IP addresses contained in endpoints objects
	endpointsIpIndex cidranger.Ranger
	// index of service target ports
	targetPortIndex map[string]*portIndexEntry
	// index of ip blocks referenced by network policy egress rules
	netPolSubnetIndex cidranger.Ranger
	// index of pods matched by erspan policies
	erspanPolPods *index.PodSelectorIndex

	apicConn *apicapi.ApicConnection

	nodeServiceMetaCache map[string]*nodeServiceMeta
	nodeOpflexDevice     map[string]apicapi.ApicSlice
	nodePodNetCache      map[string]*nodePodNetMeta
	serviceMetaCache     map[string]*serviceMeta
	snatPolicyCache      map[string]*ContSnatPolicy
	delayedEpSlices      []*DelayedEpSlice
	snatServices         map[string]bool
	snatNodeInfoCache    map[string]*nodeinfo.NodeInfo
	rdConfigCache        map[string]*rdConfig.RdConfig
	rdConfigSubnetCache  map[string]*rdConfig.RdConfigSpec
	istioCache           map[string]*istiov1.AciIstioOperator
	podIftoEp            map[string]*EndPointData
	// Node Name and Policy Name
	snatGlobalInfoCache map[string]map[string]*snatglobalinfo.GlobalInfo
	nodeSyncEnabled     bool
	serviceSyncEnabled  bool
	snatSyncEnabled     bool
	tunnelGetter        *tunnelState
	syncQueue           workqueue.RateLimitingInterface
	syncProcessors      map[string]func() bool
	serviceEndPoints    ServiceEndPointType
	crdHandlers         map[string]func(*AciController, <-chan struct{})
	stopCh              <-chan struct{}
	//index of containerportname to ctrPortNameEntry
	ctrPortNameCache map[string]*ctrPortNameEntry
	// named networkPolicies
	nmPortNp map[string]bool
	// cache to look for Epg DNs which are bound to Vmm domain
	cachedEpgDns             []string
	vmmClusterFaultSupported bool
}

type DelayedEpSlice struct {
	ServiceKey  string
	OldEpSlice  *v1beta1.EndpointSlice
	NewEpSlice  *v1beta1.EndpointSlice
	DelayedTime time.Time
}

type nodeServiceMeta struct {
	serviceEp metadata.ServiceEndpoint
}

type nodePodNetMeta struct {
	nodePods            map[string]bool
	podNetIps           metadata.NetIps
	podNetIpsAnnotation string
}

type serviceMeta struct {
	requestedIp      net.IP
	ingressIps       []net.IP
	staticIngressIps []net.IP
}

type ipIndexEntry struct {
	ipNet net.IPNet
	keys  map[string]bool
}

type targetPort struct {
	proto v1.Protocol
	ports []int
}

type portIndexEntry struct {
	port              targetPort
	serviceKeys       map[string]bool
	networkPolicyKeys map[string]bool
}

type portRangeSnat struct {
	start int
	end   int
}

//EndPointData holds PodIF data in controller.
type EndPointData struct {
	MacAddr    string
	EPG        string
	Namespace  string
	AppProfile string
}

type ctrPortNameEntry struct {
	// Proto+port->pods
	ctrNmpToPods map[string]map[string]bool
}

type ServiceEndPointType interface {
	InitClientInformer(kubeClient *kubernetes.Clientset)
	Run(stopCh <-chan struct{})
	Wait(stopCh <-chan struct{})
	UpdateServicesForNode(nodename string)
	GetnodesMetadata(key string, service *v1.Service, nodeMap map[string]*metadata.ServiceEndpoint)
	SetServiceApicObject(aobj apicapi.ApicObject, service *v1.Service) bool
	SetNpServiceAugmentForService(servicekey string, service *v1.Service, prs *portRemoteSubnet,
		portAugments map[string]*portServiceAugment, subnetIndex cidranger.Ranger, logger *logrus.Entry)
}

type serviceEndpoint struct {
	cont *AciController
}
type serviceEndpointSlice struct {
	cont *AciController
}

func (sep *serviceEndpoint) InitClientInformer(kubeClient *kubernetes.Clientset) {
	sep.cont.initEndpointsInformerFromClient(kubeClient)
}

func (seps *serviceEndpointSlice) InitClientInformer(kubeClient *kubernetes.Clientset) {
	seps.cont.initEndpointSliceInformerFromClient(kubeClient)
}

func (sep *serviceEndpoint) Run(stopCh <-chan struct{}) {
	go sep.cont.endpointsInformer.Run(stopCh)
}

func (seps *serviceEndpointSlice) Run(stopCh <-chan struct{}) {
	go seps.cont.endpointSliceInformer.Run(stopCh)
}

func (sep *serviceEndpoint) Wait(stopCh <-chan struct{}) {
	cache.WaitForCacheSync(stopCh,
		sep.cont.endpointsInformer.HasSynced,
		sep.cont.serviceInformer.HasSynced)
}

func (seps *serviceEndpointSlice) Wait(stopCh <-chan struct{}) {
	seps.cont.log.Debug("Waiting for EndPointSlicecache sync")
	cache.WaitForCacheSync(stopCh,
		seps.cont.endpointSliceInformer.HasSynced,
		seps.cont.serviceInformer.HasSynced)
}

func (e *ipIndexEntry) Network() net.IPNet {
	return e.ipNet
}

func newNodePodNetMeta() *nodePodNetMeta {
	return &nodePodNetMeta{
		nodePods: make(map[string]bool),
	}
}

func createQueue(name string) workqueue.RateLimitingInterface {
	return workqueue.NewNamedRateLimitingQueue(
		workqueue.NewMaxOfRateLimiter(
			workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond,
				10*time.Second),
			&workqueue.BucketRateLimiter{
				Limiter: rate.NewLimiter(rate.Limit(10), int(100)),
			},
		),
		"delta")
}

func NewController(config *ControllerConfig, env Environment, log *logrus.Logger, unittestmode bool) *AciController {
	cont := &AciController{
		log:          log,
		config:       config,
		env:          env,
		defaultEg:    "",
		defaultSg:    "",
		unitTestMode: unittestmode,

		podQueue:          createQueue("pod"),
		netPolQueue:       createQueue("networkPolicy"),
		qosQueue:          createQueue("qos"),
		netflowQueue:      createQueue("netflow"),
		erspanQueue:       createQueue("erspan"),
		serviceQueue:      createQueue("service"),
		snatQueue:         createQueue("snat"),
		snatNodeInfoQueue: createQueue("snatnodeinfo"),
		rdConfigQueue:     createQueue("rdconfig"),
		istioQueue:        createQueue("istio"),
		syncQueue: workqueue.NewNamedRateLimitingQueue(
			&workqueue.BucketRateLimiter{
				Limiter: rate.NewLimiter(rate.Limit(10), int(100)),
			}, "sync"),

		configuredPodNetworkIps: newNetIps(),
		podNetworkIps:           newNetIps(),
		serviceIps:              ipam.NewIpCache(),
		staticServiceIps:        newNetIps(),
		nodeServiceIps:          newNetIps(),

		nodeOpflexDevice: make(map[string]apicapi.ApicSlice),

		nodeServiceMetaCache: make(map[string]*nodeServiceMeta),
		nodePodNetCache:      make(map[string]*nodePodNetMeta),
		serviceMetaCache:     make(map[string]*serviceMeta),
		snatPolicyCache:      make(map[string]*ContSnatPolicy),
		snatServices:         make(map[string]bool),
		snatNodeInfoCache:    make(map[string]*nodeinfo.NodeInfo),
		rdConfigCache:        make(map[string]*rdConfig.RdConfig),
		rdConfigSubnetCache:  make(map[string]*rdConfig.RdConfigSpec),
		podIftoEp:            make(map[string]*EndPointData),
		snatGlobalInfoCache:  make(map[string]map[string]*snatglobalinfo.GlobalInfo),
		istioCache:           make(map[string]*istiov1.AciIstioOperator),
		crdHandlers:          make(map[string]func(*AciController, <-chan struct{})),
		ctrPortNameCache:     make(map[string]*ctrPortNameEntry),
		nmPortNp:             make(map[string]bool),
	}
	cont.syncProcessors = map[string]func() bool{
		"snatGlobalInfo": cont.syncSnatGlobalInfo,
		"rdConfig":       cont.syncRdConfig,
		"istioCR":        cont.createIstioCR,
	}
	return cont
}

func (cont *AciController) Init() {
	if cont.config.LBType != lbTypeAci {
		err := apicapi.AddMetaDataChild("vmmInjectedNs", "vmmInjectedNwPol")
		if err != nil {
			panic(err.Error())
		}
	}

	egdata, err := json.Marshal(cont.config.DefaultEg)
	if err != nil {
		cont.log.Error("Could not serialize default endpoint group")
		panic(err.Error())
	}
	cont.defaultEg = string(egdata)

	sgdata, err := json.Marshal(cont.config.DefaultSg)
	if err != nil {
		cont.log.Error("Could not serialize default security groups")
		panic(err.Error())
	}
	cont.defaultSg = string(sgdata)

	cont.log.Debug("Initializing IPAM")
	cont.initIpam()
	// check if the cluster supports endpoint slices
	// if cluster doesn't have the support fallback to endpoints
	kubeClient := cont.env.(*K8sEnvironment).kubeClient
	if util.IsEndPointSlicesSupported(kubeClient) {
		cont.serviceEndPoints = &serviceEndpointSlice{}
		cont.serviceEndPoints.(*serviceEndpointSlice).cont = cont
		cont.log.Info("Initializing ServiceEndpointSlices")
	} else {
		cont.serviceEndPoints = &serviceEndpoint{}
		cont.serviceEndPoints.(*serviceEndpoint).cont = cont
		cont.log.Info("Initializing ServiceEndpoints")
	}

	err = cont.env.Init(cont)
	if err != nil {
		panic(err.Error())
	}
}

func (cont *AciController) processQueue(queue workqueue.RateLimitingInterface,
	store cache.Store, handler func(interface{}) bool,
	postDelHandler func() bool, stopCh <-chan struct{}) {
	go wait.Until(func() {
		for {
			key, quit := queue.Get()
			if quit {
				break
			}

			var requeue bool
			switch key := key.(type) {
			case chan struct{}:
				close(key)
			case string:
				obj, exists, err := store.GetByKey(key)
				if err != nil {
					cont.log.Debugf("Error fetching object with key %s from store: %v", key, err)
				}
				//Handle Add/Update/Delete
				if exists && handler != nil {
					requeue = handler(obj)
				}
				//Handle Post Delete
				if !exists && postDelHandler != nil {
					requeue = postDelHandler()
				}
			}
			if requeue {
				queue.AddRateLimited(key)
			} else {
				queue.Forget(key)
			}
			queue.Done(key)

		}
	}, time.Second, stopCh)
	<-stopCh
	queue.ShutDown()
}

func (cont *AciController) globalStaticObjs() apicapi.ApicSlice {
	return apicapi.ApicSlice{}
}

func (cont *AciController) aciNameForKey(ktype string, key string) string {
	return util.AciNameForKey(cont.config.AciPrefix, ktype, key)
}

func (cont *AciController) initStaticObjs() {
	cont.env.InitStaticAciObjects()
	cont.apicConn.WriteApicObjects(cont.config.AciPrefix+"_static",
		cont.globalStaticObjs())
}

func (cont *AciController) vmmDomainProvider() (vmmProv string) {
	vmmProv = "Kubernetes"
	if strings.ToLower(cont.config.AciVmmDomainType) == "openshift" {
		vmmProv = "OpenShift"
	}
	return
}

func (cont *AciController) Run(stopCh <-chan struct{}) {
	var err error
	var privKey []byte
	var apicCert []byte

	cont.config.AciVrfDn = "uni/tn-" + cont.config.AciVrfTenant + "/ctx-" + cont.config.AciVrf

	if cont.config.ApicPrivateKeyPath != "" {
		privKey, err = ioutil.ReadFile(cont.config.ApicPrivateKeyPath)
		if err != nil {
			panic(err)
		}
	}
	if cont.config.ApicCertPath != "" {
		apicCert, err = ioutil.ReadFile(cont.config.ApicCertPath)
		if err != nil {
			panic(err)
		}
	}
	// If not defined, default is 1800
	if cont.config.ApicRefreshTimer == "" {
		cont.config.ApicRefreshTimer = "1800"
	}
	refreshTimeout, err := strconv.Atoi(cont.config.ApicRefreshTimer)
	if err != nil {
		panic(err)
	}
	cont.log.Info("ApicRefreshTimer conf is set to: ", refreshTimeout)

	// Bailout if the refreshTimeout is more than 12Hours or less than 5Mins
	if refreshTimeout > (12*60*60) || refreshTimeout < (5*60) {
		cont.log.Info("ApicRefreshTimer can't be more than 12Hrs or less than 5Mins")
		panic(err)
	}

	// If RefreshTickerAdjustInterval is not defined, default to 150Sec.
	if cont.config.ApicRefreshTickerAdjust == "" {
		cont.config.ApicRefreshTickerAdjust = "150"
	}
	refreshTickerAdjust, err := strconv.Atoi(cont.config.ApicRefreshTickerAdjust)
	if err != nil {
		panic(err)
	}

	//If ApicSubscriptionDelay is not defined, default to 100ms
	if cont.config.ApicSubscriptionDelay == 0 {
		cont.config.ApicSubscriptionDelay = 100
	}
	cont.log.Info("ApicSubscriptionDelay conf is set to: ", cont.config.ApicSubscriptionDelay)

	// If OpflexDeviceDeleteTimeout is not defined, default to 1800s
	if cont.config.OpflexDeviceDeleteTimeout == 0 {
		cont.config.OpflexDeviceDeleteTimeout = 1800
	}

	// If SleepTimeSnatGlobalInfoSync is not defined, default to 60
	if cont.config.SleepTimeSnatGlobalInfoSync == 0 {
		cont.config.SleepTimeSnatGlobalInfoSync = 60
	}

	// If not defined, default to 32
	if cont.config.PodIpPoolChunkSize == 0 {
		cont.config.PodIpPoolChunkSize = 32
	}
	cont.log.Info("PodIpPoolChunkSize conf is set to: ", cont.config.PodIpPoolChunkSize)

	// If not valid, default to 5000-65000
	// other permissible values 1-65000
	defStart := 5000
	defEnd := 65000
	if cont.config.SnatDefaultPortRangeStart == 0 {
		cont.config.SnatDefaultPortRangeStart = defStart
	}
	if cont.config.SnatDefaultPortRangeEnd == 0 {
		cont.config.SnatDefaultPortRangeEnd = defEnd
	}
	if cont.config.SnatDefaultPortRangeStart < 0 || cont.config.SnatDefaultPortRangeEnd < 0 ||
		cont.config.SnatDefaultPortRangeStart > defEnd || cont.config.SnatDefaultPortRangeEnd > defEnd ||
		cont.config.SnatDefaultPortRangeStart > cont.config.SnatDefaultPortRangeEnd {
		cont.config.SnatDefaultPortRangeStart = defStart
		cont.config.SnatDefaultPortRangeEnd = defEnd
	}

	// Set default value for pbr programming delay if services list is not empty
	// and delay value is empty
	if cont.config.ServiceGraphEndpointAddDelay.Delay == 0 &&
		cont.config.ServiceGraphEndpointAddDelay.Services != nil &&
		len(cont.config.ServiceGraphEndpointAddDelay.Services) > 0 {
		cont.config.ServiceGraphEndpointAddDelay.Delay = 90
	}
	if cont.config.ServiceGraphEndpointAddDelay.Delay > 0 {
		cont.log.Info("ServiceGraphEndpointAddDelay set to: ", cont.config.ServiceGraphEndpointAddDelay.Delay)
	}

	// Set contract scope for snat svc graph to global by default
	if cont.config.SnatSvcContractScope == "" {
		cont.config.SnatSvcContractScope = "global"
	}
	if cont.config.MaxSvcGraphNodes == 0 {
		cont.config.MaxSvcGraphNodes = 32
	}
	cont.log.Info("Max number of nodes per svc graph is set to: ", cont.config.MaxSvcGraphNodes)

	cont.apicConn, err = apicapi.New(cont.log, cont.config.ApicHosts,
		cont.config.ApicUsername, cont.config.ApicPassword,
		privKey, apicCert, cont.config.AciPrefix,
		refreshTimeout, refreshTickerAdjust, cont.config.ApicSubscriptionDelay)
	if err != nil {
		panic(err)
	}

	if len(cont.config.ApicHosts) != 0 {
		cont.log.WithFields(logrus.Fields{
			"mod":  "APICAPI",
			"host": cont.apicConn.Apic[cont.apicConn.ApicIndex],
		}).Debug("Connecting to APIC to determine the Version")

		version, err := cont.apicConn.GetVersion()
		if err != nil {
			cont.log.Error("Could not get APIC version")
			panic(err)
		}
		cont.apicConn.CachedVersion = version
		apicapi.ApicVersion = version
		if version >= "4.2(4i)" {
			cont.apicConn.SnatPbrFltrChain = true
		} else {
			cont.apicConn.SnatPbrFltrChain = false
		}
		if version >= "5.2" {
			cont.vmmClusterFaultSupported = true
		}

	} else { // For unit-tests
		cont.apicConn.SnatPbrFltrChain = true
	}

	cont.log.Debug("SnatPbrFltrChain set to:", cont.apicConn.SnatPbrFltrChain)
	// Make sure Pod/NodeBDs are assoicated to same VRF.
	if len(cont.config.ApicHosts) != 0 && cont.config.AciPodBdDn != "" && cont.config.AciNodeBdDn != "" {
		var expectedVrfRelations []string
		expectedVrfRelations = append(expectedVrfRelations, cont.config.AciPodBdDn, cont.config.AciNodeBdDn)
		cont.log.Debug("expectedVrfRelations:", expectedVrfRelations)
		err = cont.apicConn.ValidateAciVrfAssociation(cont.config.AciVrfDn, expectedVrfRelations)
		if err != nil {
			cont.log.Error("Pod/NodeBDs and AciL3Out VRF association is incorrect")
			panic(err)
		}
	}

	if len(cont.config.ApicHosts) != 0 && cont.vmmClusterFaultSupported {
		//Clear fault instances when the controller starts
		cont.clearFaultInstances()
		//Subscribe for vmmEpPD for a given domain
		var tnTargetFilterEpg string
		tnTargetFilterEpg += fmt.Sprintf("uni/vmmp-%s/dom-%s", cont.vmmDomainProvider(), cont.config.AciVmmDomain)
		subnetTargetFilterEpg := fmt.Sprintf("and(wcard(vmmEpPD.dn,\"%s\"))", tnTargetFilterEpg)
		cont.apicConn.AddSubscriptionClass("vmmEpPD",
			[]string{"vmmEpPD"}, subnetTargetFilterEpg)
		cont.apicConn.SetSubscriptionHooks("vmmEpPD",
			func(obj apicapi.ApicObject) bool {
				cont.vmmEpPDChanged(obj)
				return true
			},
			func(dn string) {
				cont.vmmEpPDDeleted(dn)
			})

	}

	cont.initStaticObjs()
	err = cont.env.PrepareRun(stopCh)
	if err != nil {
		panic(err.Error())
	}

	cont.apicConn.FullSyncHook = func() {
		// put a channel into each work queue and wait on it to
		// checkpoint object syncing in response to new subscription
		// updates
		cont.log.Debug("Starting checkpoint")
		var chans []chan struct{}
		qs := make([]workqueue.RateLimitingInterface, 0)
		_, ok := cont.env.(*K8sEnvironment)
		if ok {
			qs = []workqueue.RateLimitingInterface{
				cont.podQueue, cont.netPolQueue, cont.qosQueue,
				cont.serviceQueue, cont.snatQueue, cont.netflowQueue,
				cont.snatNodeInfoQueue, cont.rdConfigQueue, cont.erspanQueue,
			}
		}
		for _, q := range qs {
			c := make(chan struct{})
			chans = append(chans, c)
			q.Add(c)
		}
		for _, c := range chans {
			<-c
		}
		cont.log.Debug("Checkpoint complete")
	}

	if len(cont.config.ApicHosts) != 0 {
		cont.BuildSubnetDnCache(cont.config.AciVrfDn, cont.config.AciVrfDn)
		cont.scheduleRdConfig()
	}

	cont.apicConn.AddSubscriptionDn("uni/tn-"+cont.config.AciPolicyTenant,
		[]string{"hostprotPol"})
	cont.apicConn.AddSubscriptionDn("uni/tn-"+cont.config.AciVrfTenant,
		[]string{"fvBD", "vnsLDevVip", "vnsAbsGraph", "vnsLDevCtx",
			"vzFilter", "vzBrCP", "l3extInstP", "vnsSvcRedirectPol",
			"vnsRedirectHealthGroup", "fvIPSLAMonitoringPol"})
	cont.apicConn.AddSubscriptionDn(fmt.Sprintf("uni/tn-%s/out-%s",
		cont.config.AciVrfTenant, cont.config.AciL3Out),
		[]string{"fvRsCons"})
	vmmDn := fmt.Sprintf("comp/prov-%s/ctrlr-[%s]-%s/injcont",
		cont.env.VmmPolicy(), cont.config.AciVmmDomain,
		cont.config.AciVmmController)
	// Before subscribing to vmm objects, add vmmInjectedLabel as a child after explicit APIC version check
	// Since it is not supported for APIC versions < "5.0"
	cont.addVmmInjectedLabel()
	cont.apicConn.AddSubscriptionDn(vmmDn,
		[]string{"vmmInjectedHost", "vmmInjectedNs"})

	var tnTargetFilter string
	if len(cont.config.AciVrfRelatedTenants) > 0 {
		for _, tn := range cont.config.AciVrfRelatedTenants {
			tnTargetFilter += fmt.Sprintf("tn-%s|", tn)
		}
	} else {
		tnTargetFilter += fmt.Sprintf("tn-%s|tn-%s",
			cont.config.AciPolicyTenant, cont.config.AciVrfTenant)
	}
	subnetTargetFilter := fmt.Sprintf("and(wcard(fvSubnet.dn,\"%s\"))",
		tnTargetFilter)
	cont.apicConn.AddSubscriptionClass("fvSubnet",
		[]string{"fvSubnet"}, subnetTargetFilter)

	cont.apicConn.SetSubscriptionHooks("fvSubnet",
		func(obj apicapi.ApicObject) bool {
			cont.SubnetChanged(obj, cont.config.AciVrfDn)
			return true
		},
		func(dn string) {
			cont.SubnetDeleted(dn)
		})

	cont.apicConn.AddSubscriptionClass("opflexODev",
		[]string{"opflexODev"}, "")

	cont.apicConn.SetSubscriptionHooks("opflexODev",
		func(obj apicapi.ApicObject) bool {
			cont.opflexDeviceChanged(obj)
			return true
		},
		func(dn string) {
			cont.opflexDeviceDeleted(dn)
		})
	go cont.apicConn.Run(stopCh)
}

func (cont *AciController) syncOpflexDevices(stopCh <-chan struct{}, seconds time.Duration) {
	cont.log.Debug("Go routine to periodically delete old opflexdevices started")
	ticker := time.NewTicker(seconds * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cont.deleteOldOpflexDevices()
		case <-stopCh:
			return
		}
	}
}

func (cont *AciController) syncDelayedEpSlices(stopCh <-chan struct{}, seconds time.Duration) {
	cont.log.Debug("Go routine to periodically check and process the epslices having delay adding in service")
	ticker := time.NewTicker(seconds * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cont.processDelayedEpSlices()
		case <-stopCh:
			return
		}
	}
}

func (cont *AciController) snatGlobalInfoSync(stopCh <-chan struct{}, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	cont.log.Debug("Go routine to periodically sync globalinfo and nodeinfo started")
	iteration := 0
	for {
		// To avoid noisy logs, only printing once in 5 minutes
		if iteration%5 == 0 {
			cont.log.Debug("Syncing GlobalInfo with Node infos")
		}
		var nodeInfos []*nodeinfo.NodeInfo
		cont.indexMutex.Lock()
		cache.ListAll(cont.snatNodeInfoIndexer, labels.Everything(),
			func(nodeInfoObj interface{}) {
				nodeInfo := nodeInfoObj.(*nodeinfo.NodeInfo)
				nodeInfos = append(nodeInfos, nodeInfo)
			})
		expectedmap := make(map[string]map[string]bool)
		for _, glinfo := range cont.snatGlobalInfoCache {
			for nodename, entry := range glinfo {
				if _, found := expectedmap[nodename]; !found {
					newentry := make(map[string]bool)
					newentry[entry.SnatPolicyName] = true
					expectedmap[nodename] = newentry
				} else {
					currententry := expectedmap[nodename]
					currententry[entry.SnatPolicyName] = true
					expectedmap[nodename] = currententry
				}
			}
		}
		cont.indexMutex.Unlock()

		for _, value := range nodeInfos {
			marked := false
			policyNames := value.Spec.SnatPolicyNames
			nodeName := value.ObjectMeta.Name
			_, ok := expectedmap[nodeName]
			if !ok && len(policyNames) > 0 {
				cont.log.Info("Adding missing entry in snatglobalinfo for node: ", nodeName)
				cont.log.Debug("No snat policies found in snatglobalinfo")
				cont.log.Debug("Snatpolicy list according to nodeinfo: ", policyNames)
				marked = true
			} else if len(policyNames) != len(expectedmap[nodeName]) {
				cont.log.Info("Adding missing snatpolicy entry in snatglobalinfo for node: ", nodeName)
				cont.log.Debug("Snatpolicy list according to snatglobalinfo: ", expectedmap[nodeName])
				cont.log.Debug("Snatpolicy list according to nodeinfo: ", policyNames)
				marked = true
			} else {
				if len(policyNames) == 0 && len(expectedmap[nodeName]) == 0 {
					// No snatpolicies present
					continue
				}
				eq := reflect.DeepEqual(expectedmap[nodeName], policyNames)
				if !eq {
					cont.log.Debug("Syncing inconsistent snatpolicy entry in snatglobalinfo for node: ", nodeName)
					cont.log.Debug("Snatpolicy list according to snatglobalinfo: ", expectedmap[nodeName])
					cont.log.Debug("Snatpolicy list according to nodeinfo: ", policyNames)
					marked = true
				}
			}
			if marked {
				cont.log.Info("Nodeinfo and globalinfo out of sync for node: ", nodeName)
				nodeinfokey, err := cache.MetaNamespaceKeyFunc(value)
				if err != nil {
					cont.log.Error("Not able to get key for node: ", nodeName)
					continue
				}
				cont.log.Info("Queuing nodeinfokey for globalinfo sync: ", nodeinfokey)
				cont.queueNodeInfoUpdateByKey(nodeinfokey)
			} else {
				if iteration%5 == 0 {
					cont.log.Info("Nodeinfo and globalinfo in sync for node: ", nodeName)
				}
			}
		}
		time.Sleep(time.Duration(seconds) * time.Second)
		iteration = iteration + 1
	}
}

func (cont *AciController) processSyncQueue(queue workqueue.RateLimitingInterface,
	queueStop <-chan struct{}) {

	go wait.Until(func() {
		for {
			syncType, quit := queue.Get()
			if quit {
				break
			}
			var requeue bool
			switch syncType := syncType.(type) {
			case string:
				if f, ok := cont.syncProcessors[syncType]; ok {
					requeue = f()
				}
			}
			if requeue {
				queue.AddRateLimited(syncType)
			} else {
				queue.Forget(syncType)
			}
			queue.Done(syncType)

		}
	}, time.Second, queueStop)
	<-queueStop
	queue.ShutDown()
}

func (cont *AciController) scheduleSyncGlobalInfo() {
	cont.syncQueue.AddRateLimited("snatGlobalInfo")
}
func (cont *AciController) scheduleRdConfig() {
	cont.syncQueue.AddRateLimited("rdConfig")
}
func (cont *AciController) scheduleCreateIstioCR() {
	cont.syncQueue.AddRateLimited("istioCR")
}

func (cont *AciController) addVmmInjectedLabel() {
	if apicapi.ApicVersion >= "5.2" {
		err := apicapi.AddMetaDataChild("vmmInjectedNs", "vmmInjectedLabel")
		if err != nil {
			panic(err.Error())
		}
		err = apicapi.AddMetaDataChild("vmmInjectedSvc", "vmmInjectedLabel")
		if err != nil {
			panic(err.Error())
		}
	}
	if apicapi.ApicVersion >= "5.0" {
		err := apicapi.AddMetaDataChild("vmmInjectedReplSet", "vmmInjectedLabel")
		if err != nil {
			panic(err.Error())
		}
		err = apicapi.AddMetaDataChild("vmmInjectedContGrp", "vmmInjectedLabel")
		if err != nil {
			panic(err.Error())
		}
		err = apicapi.AddMetaDataChild("vmmInjectedDepl", "vmmInjectedLabel")
		if err != nil {
			panic(err.Error())
		}
	}
}
