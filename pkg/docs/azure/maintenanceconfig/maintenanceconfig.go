package maintenanceconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var MaintenanceConfigurationsClient *armcontainerservice.MaintenanceConfigurationsClient

func MaintenanceConfigurationResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateOrUpdateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(ListByManagedClusterResource))
	maintenanceConfigurationClientInit()
}

func maintenanceConfigurationClientInit() {
	sess := types.GetAKSClient()
	MaintenanceConfigurationsClient = sess.MaintenanceConfigClient
}
