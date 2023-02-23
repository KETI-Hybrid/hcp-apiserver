package nodegroup

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var NodeGroupClient *eks.EKS

func NodeGroupResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpgradeResource))
	nodeGroupClientInit()
}

func nodeGroupClientInit() {
	sess := types.GetEKSClient()
	NodeGroupClient = eks.New(sess.Client)
}
