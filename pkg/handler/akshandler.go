package handler

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/converter"
	auth "Hybrid_Cluster/hcp-apiserver/pkg/util"
	"Hybrid_Cluster/hybridctl/util"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func AksStart(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["start"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	fmt.Println(api)
	hosturl := api
	response, err := auth.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksStop(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["stop"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	hosturl := api
	response, err := auth.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksRotateCerts(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["rotateCerts"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	fmt.Println(api)
	hosturl := api
	response, err := auth.AuthorizationAndHTTP("POST", hosturl, nil)
	return response, err
}

func AksGetOSoptions(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["getOSoptions"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{location}", input.Location)
	hosturl := api
	fmt.Println(api)
	response, err := auth.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationCreateOrUpdate(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationCreate/Update"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	fmt.Println(api)
	hosturl := api
	response, err := auth.AuthorizationAndHTTP("PUT", hosturl, input.ConfigFile)
	return response, err
}

func MaintenanceconfigurationDelete(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationDelete"]
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	fmt.Println(api)
	response, err := auth.AuthorizationAndHTTP("DELETE", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationList(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationList"]
	fmt.Println(input)
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	fmt.Println(api)
	hosturl := api
	response, err := auth.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}

func MaintenanceconfigurationShow(input util.EKSAPIParameter) (*http.Response, error) {
	api := converter.AksAPI["maintenanceconfigurationShow"]
	fmt.Println(input)
	api = strings.ReplaceAll(api, "{subscriptionId}", os.Getenv("SubscriptionId"))
	api = strings.ReplaceAll(api, "{resourceGroupName}", input.ResourceGroupName)
	api = strings.ReplaceAll(api, "{resourceName}", input.ResourceName)
	api = strings.ReplaceAll(api, "{configName}", input.ConfigName)
	hosturl := api
	// fmt.Println(api)
	response, err := auth.AuthorizationAndHTTP("GET", hosturl, nil)
	return response, err
}
