package nodegroup

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var NodeGroupClient *eks.EKS

func NodeGroupResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(UpgradeResource))
	nodeGroupClientInit()
}

func nodeGroupClientInit() {
	sess := types.GetEKSClient()
	NodeGroupClient = eks.New(sess.Client)
}
