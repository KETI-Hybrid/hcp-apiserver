package snapshot

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var SnapshotsClient *armcontainerservice.SnapshotsClient

func SnapshotResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateOrUpdateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(ListByResourceGroupResource))
	docs.AddResource(router, new(UpdateTagsResource))
	snapshotClientInit()
}

func snapshotClientInit() {
	sess := types.GetAKSClient()
	SnapshotsClient = sess.SnapshotClient
}
