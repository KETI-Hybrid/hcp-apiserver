package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var AddonClient *eks.EKS

func AddonResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(DescribeConfigResource))
	docs.AddResource(router, new(DescribeVersionResource))
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(UpdateResource))
	addOnClientInit()
}

func addOnClientInit() {
	sess := types.GetEKSClient()
	AddonClient = eks.New(sess.Client)
}
