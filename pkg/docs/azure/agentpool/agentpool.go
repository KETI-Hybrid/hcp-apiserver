package agentpool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var AgentPoolsClient *armcontainerservice.AgentPoolsClient

func AgentPoolResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateOrUpdateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(GetAvailableAgentPoolVersionsResource))
	docs.AddResource(router, new(GetUpgradeProfileResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(UpdateNodeImageVersionResource))
	agentpoolClientInit()
}

func agentpoolClientInit() {
	sess := types.GetAKSClient()
	AgentPoolsClient = sess.AgentPoolClient
}
