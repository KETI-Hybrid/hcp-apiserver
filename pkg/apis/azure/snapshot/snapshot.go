package snapshot

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var SnapshotsClient *armcontainerservice.SnapshotsClient

func SnapshotResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateOrUpdateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(ListByResourceGroupResource))
	apis.AddResource(router, new(UpdateTagsResource))
	snapshotClientInit()
}

func snapshotClientInit() {
	sess := types.GetAKSClient()
	SnapshotsClient = sess.SnapshotClient
}
