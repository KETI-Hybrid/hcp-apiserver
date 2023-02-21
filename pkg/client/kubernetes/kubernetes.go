package kubernetes

import (
	"hcp-apiserver/pkg/client/aks"
	"hcp-apiserver/pkg/client/eks"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/aws/aws-sdk-go/aws/session"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

type HCPClient struct {
	KubeClient *kubernetes.Clientset
	EKSClient  *session.Session
	AKSClient  *armcontainerservice.ManagedClustersClient
	GKEClient  string
	NKSClient  string
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // 윈도우에서는 USERPROFILE을 사용
}

func InitHCPClient() *HCPClient {
	var kubeconfig string
	config := new(rest.Config)
	var err error

	hcpClient := &HCPClient{}

	if home := homeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}

	if len(kubeconfig) == 0 {
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.Errorln(err.Error())
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			klog.Errorln(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Errorln(err.Error())
	}
	hcpClient.KubeClient = clientset
	hcpClient.EKSClient = eks.NewEKSClient(clientset)
	hcpClient.AKSClient = aks.NewAKSClient(clientset)
	return nil
}
