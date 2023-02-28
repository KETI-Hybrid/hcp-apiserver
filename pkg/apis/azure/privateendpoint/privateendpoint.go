package privateendpoint

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var PrivateEndpointConnectionsClient *armcontainerservice.PrivateEndpointConnectionsClient

func PrivateEndpointConnectionsResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(UpdateResource))
	privateEndpointConnectionsClientInit()
}

func privateEndpointConnectionsClientInit() {
	sess := types.GetAKSClient()
	PrivateEndpointConnectionsClient = sess.PrivateEndpointClient
}
