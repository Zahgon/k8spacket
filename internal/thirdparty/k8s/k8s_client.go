package k8sclient

import (
	"os"
	"strconv"
	"sync"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

type ipResourceInfoType string

const (
	Node ipResourceInfoType = "Node"
	Pod  ipResourceInfoType = "Pod"
	Svc  ipResourceInfoType = "Svc"
)

type ipResourceInfo struct {
	ipResourceInfoType ipResourceInfoType
	Name               string
	Namespace          string
}

type K8SClient struct {
	Client
}

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]ipResourceInfo
}

var k8sInfo *SafeMap

var clientset *kubernetes.Clientset

var disabledK8sResource, _ = strconv.ParseBool(os.Getenv("K8S_PACKET_K8S_RESOURCES_DISABLED"))

func init() {
	k8sInfo = &SafeMap{data: make(map[string]ipResourceInfo)}
	if !disabledK8sResource {
		_, clientset = configClusterClient()
		factory := informers.NewSharedInformerFactoryWithOptions(clientset, 5*time.Minute)
		stopChan := make(chan struct{})
		createPodInformer(factory)
		createSvcInformer(factory)
		createNodeInformer(factory)
		factory.Start(stopChan)
	}
}

func (k8sClient *K8SClient) GetPodIPsBySelectors(fieldSelector string, labelSelector string) []string {
	_ = "STUB: not implemented"
	return nil
}

func configClusterClient() (error, *kubernetes.Clientset) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPodInformer(factory informers.SharedInformerFactory) { _ = "STUB: not implemented"; return }

func addPod(obj interface{}) { _ = "STUB: not implemented"; return }

func createSvcInformer(factory informers.SharedInformerFactory) { _ = "STUB: not implemented"; return }

func addSvc(obj interface{}) { _ = "STUB: not implemented"; return }

func createNodeInformer(factory informers.SharedInformerFactory) { _ = "STUB: not implemented"; return }

func addNode(obj interface{}) { _ = "STUB: not implemented"; return }

func GetNameAndNamespace(id string) (string, string) { _ = "STUB: not implemented"; return "", "" }

func addItem(id string, info ipResourceInfo) { _ = "STUB: not implemented"; return }
