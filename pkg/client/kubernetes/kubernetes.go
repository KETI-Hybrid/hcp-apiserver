package kubernetes

import (
	"hcp-apiserver/pkg/client/aks"
	"hcp-apiserver/pkg/client/eks"
	"hcp-apiserver/pkg/client/gke"
	"hcp-apiserver/pkg/client/nks"
	"hcp-apiserver/pkg/types"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // 윈도우에서는 USERPROFILE을 사용
}

func InitHCPClient() {
	var kubeconfig string
	config := new(rest.Config)
	var err error

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
	types.GlobalClient.KubeClient = clientset
	types.GlobalClient.EKSClient = eks.NewEKSClient(clientset)
	types.GlobalClient.AKSClient = aks.NewAKSClient(clientset)
	types.GlobalClient.GKEClient = gke.NewGKEClient(clientset)
	types.GlobalClient.NKSClient = nks.NewNKSClient(clientset)
}
