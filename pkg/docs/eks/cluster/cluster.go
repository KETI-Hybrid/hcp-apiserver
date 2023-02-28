package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var ClusterClient *eks.EKS

func ClusterResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(DeregisterResource))
	docs.AddResource(router, new(RegisterResource))
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(UpgradeResource))
	clusterClientInit()
}

func clusterClientInit() {
	sess := types.GetEKSClient()
	ClusterClient = eks.New(sess.Client)
}
