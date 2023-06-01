package maintenanceconfig

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var MaintenanceConfigurationsClient *armcontainerservice.MaintenanceConfigurationsClient

func MaintenanceConfigurationResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateOrUpdateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(ListByManagedClusterResource))
	maintenanceConfigurationClientInit()
}

func maintenanceConfigurationClientInit() {
	sess := types.GetAKSClient()
	MaintenanceConfigurationsClient = sess.MaintenanceConfigClient
}
