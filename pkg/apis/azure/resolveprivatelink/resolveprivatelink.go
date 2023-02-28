package resolveprivatelink

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var ResolvePrivateLinkServiceIdClient *armcontainerservice.ResolvePrivateLinkServiceIDClient

func ResolvePrivateLinkServiceIdResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(POSTResource))
	resolvePrivateLinkServiceIdClientInit()
}

func resolvePrivateLinkServiceIdClientInit() {
	sess := types.GetAKSClient()
	ResolvePrivateLinkServiceIdClient = sess.ResolvePrivateLinkClient
}
