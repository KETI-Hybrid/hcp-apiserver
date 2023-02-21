package aks

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func NewAKSClient(k8sclient *kubernetes.Clientset) *armcontainerservice.ManagedClustersClient {
	config, err := k8sclient.CoreV1().ConfigMaps("public-auth").Get(context.Background(), "aks-auth", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err.Error())
	}

	tanantID := config.Data["tenant-id"]
	clientID := config.Data["client-id"]
	clientSecret := config.Data["client-secret"]
	subscriptionID := config.Data["ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73"]

	credention, err := azidentity.NewClientSecretCredential(tanantID, clientID, clientSecret, new(azidentity.ClientSecretCredentialOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	return client
}
