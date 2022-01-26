package handler

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/converter"
	util "Hybrid_Cluster/hcp-apiserver/pkg/util"

	"fmt"
	"net/http"
	"os"
	"strings"
)

func AksStart(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["start"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	fmt.Println(api)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksStop(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["stop"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksRotateCerts(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["rotateCerts"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	fmt.Println(api)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksGetOSoptions(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["getOSoptions"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{location}", input.Location)
	hosturl := api
	fmt.Println(api)
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationCreateOrUpdate(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationCreate/Update"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	fmt.Println(api)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("PUT", hosturl, input.ConfigFile)
	return response, err
}

func MaintenanceconfigurationDelete(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationDelete"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	fmt.Println(api)
	response, err := util.AuthorizationAndHTTP("DELETE", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationList(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationList"]
	fmt.Println(input)
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	fmt.Println(api)
	hosturl := api
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationShow(input util.AKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationShow"]
	fmt.Println(input)
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ClusterName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	// fmt.Println(api)
	response, err := util.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}
