package cluster

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var ClusterClient *eks.EKS

func ClusterResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DeregisterResource))
	apis.AddResource(router, new(RegisterResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpgradeResource))
	clusterClientInit()
}

func clusterClientInit() {
	sess := types.GetEKSClient()
	ClusterClient = eks.New(sess.Client)
}
