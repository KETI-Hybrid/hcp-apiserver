package agentpool

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var AgentPoolsClient *armcontainerservice.AgentPoolsClient

func AgentPoolResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateOrUpdateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(GetAvailableAgentPoolVersionsResource))
	apis.AddResource(router, new(GetUpgradeProfileResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpdateNodeImageVersionResource))
	agentpoolClientInit()
}

func agentpoolClientInit() {
	sess := types.GetAKSClient()
	AgentPoolsClient = sess.AgentPoolClient
}
