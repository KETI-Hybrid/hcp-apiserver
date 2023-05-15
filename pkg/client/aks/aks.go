package aks

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type AKSClientSet struct {
	AgentPoolClient          *armcontainerservice.AgentPoolsClient
	MaintenanceConfigClient  *armcontainerservice.MaintenanceConfigurationsClient
	ManagedClusterClient     *armcontainerservice.ManagedClustersClient
	OperationClient          *armcontainerservice.OperationsClient
	PrivateEndpointClient    *armcontainerservice.PrivateEndpointConnectionsClient
	PrivateLinkClient        *armcontainerservice.PrivateLinkResourcesClient
	ResolvePrivateLinkClient *armcontainerservice.ResolvePrivateLinkServiceIDClient
	SnapshotClient           *armcontainerservice.SnapshotsClient
}

func NewAKSClient(k8sclient *kubernetes.Clientset) *AKSClientSet {
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

	client := &AKSClientSet{}
	client.AgentPoolClient, err = armcontainerservice.NewAgentPoolsClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.MaintenanceConfigClient, err = armcontainerservice.NewMaintenanceConfigurationsClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.ManagedClusterClient, err = armcontainerservice.NewManagedClustersClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.OperationClient, err = armcontainerservice.NewOperationsClient(credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.PrivateEndpointClient, err = armcontainerservice.NewPrivateEndpointConnectionsClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.PrivateLinkClient, err = armcontainerservice.NewPrivateLinkResourcesClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.ResolvePrivateLinkClient, err = armcontainerservice.NewResolvePrivateLinkServiceIDClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.SnapshotClient, err = armcontainerservice.NewSnapshotsClient(subscriptionID, credention, new(policy.ClientOptions))
	if err != nil {
		klog.Errorln(err.Error())
	}

	return client
}
