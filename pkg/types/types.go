package types

import (
	"hcp-apiserver/pkg/client/aks"
	"hcp-apiserver/pkg/client/eks"
	"hcp-apiserver/pkg/client/gke"
	"hcp-apiserver/pkg/client/nks"

	"k8s.io/client-go/kubernetes"
)

type HCPClient struct {
	KubeClient *kubernetes.Clientset
	EKSClient  *eks.EKSClient
	AKSClient  *aks.AKSClientSet
	GKEClient  *gke.GKEClient
	NKSClient  *nks.NKSClient
}

var GlobalClient *HCPClient

func GetEKSClient() *eks.EKSClient {
	return GlobalClient.EKSClient
}

func GetAKSClient() *aks.AKSClientSet {
	return GlobalClient.AKSClient
}

func GetGKEClient() *gke.GKEClient {
	return GlobalClient.GKEClient
}

func GetNKSClient() *nks.NKSClient {
	return GlobalClient.NKSClient
}
