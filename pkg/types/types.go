package types

import (
	"hcp-apiserver/pkg/client/eks"
	"hcp-apiserver/pkg/client/gke"
	"hcp-apiserver/pkg/client/nks"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"k8s.io/client-go/kubernetes"
)

type HCPClient struct {
	KubeClient *kubernetes.Clientset
	EKSClient  *eks.EKSClient
	AKSClient  *armcontainerservice.ManagedClustersClient
	GKEClient  *gke.GKEClient
	NKSClient  *nks.NKSClient
}

var GlobalClient *HCPClient

func GetEKSClient() *eks.EKSClient {
	return GlobalClient.EKSClient
}

func GetAKSClient() *armcontainerservice.ManagedClustersClient {
	return GlobalClient.AKSClient
}

func GetGKEClient() *gke.GKEClient {
	return GlobalClient.GKEClient
}

func GetNKSClient() *nks.NKSClient {
	return GlobalClient.NKSClient
}
