package addon

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var AddonClient *eks.EKS

func AddonResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DescribeConfigResource))
	apis.AddResource(router, new(DescribeVersionResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpdateResource))
	addOnClientInit()
}

func addOnClientInit() {
	sess := types.GetEKSClient()
	AddonClient = eks.New(sess.Client)
}
