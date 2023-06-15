package privatelink

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var PrivateLinkClient *armcontainerservice.PrivateLinkResourcesClient

func PrivateLinkResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(ListResource))
	privateLinkClientInit()
}

func privateLinkClientInit() {
	sess := types.GetAKSClient()
	PrivateLinkClient = sess.PrivateLinkClient
}
