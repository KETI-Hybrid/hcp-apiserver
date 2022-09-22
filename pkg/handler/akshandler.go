package handler

import (
	"context"
	"io/ioutil"

	"github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/converter"
	util "github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/util"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"

	"net/http"
	"os"
	"strings"

	"k8s.io/klog"
)

func AksNotusedCreateCluster() {
	klog.Infoln("create test cluster")

	type test struct {
		Location string `json:"location"`
	}
	api := converter.AKSAPI["createCluster"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", "hcpResourceGroup")
	api = strings.ReplaceAll(api, "{resourceName}", "aks-keti-cluster31")
	hosturl := api

	aa := &test{}
	aa.Location = "East US"

	lo := "East US"

	cc := armcontainerservice.ManagedCluster{}
	klog.Infoln(cc)
	cc.Location = &lo
	klog.Infoln(cc)
	cc.Properties.ServicePrincipalProfile.ClientID = to.Ptr("")
	klog.Infoln(cc)
	cc.Properties.ServicePrincipalProfile.Secret = to.Ptr("")

	res, err := util.AuthorizationAndHTTP("PUT", hosturl, cc)
	if err != nil {
		klog.Infoln(err)
	}
	_ = res
	klog.Infoln(res)
	klog.Infoln(to.Ptr("production"))

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		klog.Infoln(err)
	}
	klog.Infoln(string(data))
}

func AksTestCreateCluster(input util.AKSAPIParameter) (string, error) {

	resourceGroup := input.ResourceGroupName
	clusterName := input.ClusterName
	location := input.Location

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		klog.Infoln(err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewManagedClustersClient("", cred, nil)
	if err != nil {
		klog.Infoln(err)
	}

	poller, err := client.BeginCreateOrUpdate(ctx,
		resourceGroup,
		clusterName,
		armcontainerservice.ManagedCluster{
			Location: to.Ptr(location),
			Properties: &armcontainerservice.ManagedClusterProperties{
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
						Count:              to.Ptr[int32](3),
						EnableFIPS:         to.Ptr(true),
						EnableNodePublicIP: to.Ptr(true),
						Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
						OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
						VMSize:             to.Ptr("Standard_DS2_v2"),
						Name:               to.Ptr("aksketinp1"),
					}},
				DNSPrefix: to.Ptr("aks-keti-cluster33-dns"),
				ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.Ptr(""),
					Secret:   to.Ptr(""),
				},
			},
		},
		nil)
	if err != nil {
		klog.Infoln(err)
	}

	result := ""
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		klog.Infoln(err)
		result = "Aks Cluster Not Created"
	} else {
		result = "Aks Cluster Created"
	}
	_ = res
	klog.Infoln(res)

	return result, err
}

func AksDescribeCluster(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["getCluster"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	// api = strings.ReplaceAll(api, "{resourceGroupName}", "hcpResourceGroup")
	// api = strings.ReplaceAll(api, "{resourceName}", "aks-keti-cluster1")
	hosturl := api
	klog.Infoln(hosturl)
	res, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	if err != nil {
		klog.Infoln(err)
	}
	klog.Infoln(res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		klog.Infoln(err)
	}
	klog.Infoln(string(data))
	return res, err
}

func AKSStart(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["start"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AKSStop(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["stop"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AKSRotateCerts(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["rotateCerts"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AKSGetOSoptions(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["getOSoptions"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{location}", input.Location)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationCreateOrUpdate(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["maintenanceconfigurationCreate/Update"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("PUT", hosturl, input.ConfigFile)
	return response, err
}

func MaintenanceconfigurationDelete(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["maintenanceconfigurationDelete"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("DELETE", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationList(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["maintenanceconfigurationList"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationShow(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AKSAPI["maintenanceconfigurationShow"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}
