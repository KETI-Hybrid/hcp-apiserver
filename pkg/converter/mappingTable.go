package converter

var AKSAPI map[string]string = map[string]string{
	"start":                                 "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/start?api-version=2021-05-01",
	"stop":                                  "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/stop?api-version=2021-05-01",
	"rotateCerts":                           "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/rotateClusterCertificates?api-version=2021-05-01",
	"getOSoptions":                          "https://management.azure.com/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/osOptions/default?api-version=2021-05-01",
	"maintenanceconfigurationCreate/Update": "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}?api-version=2021-05-01",
	"maintenanceconfigurationDelete":        "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}?api-version=2021-05-01",
	"maintenanceconfigurationList":          "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations?api-version=2021-05-01",
	"maintenanceconfigurationShow":          "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/maintenanceConfigurations/{configName}?api-version=2021-05-01",

	"getCluster":    "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}?api-version=2022-04-01",
	"createCluster": "https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}?api-version=2022-04-01",
}
