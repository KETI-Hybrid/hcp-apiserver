package nodegroupconfig

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var NodeGroupConfigClient *eks.EKS

func NodeGroupConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(UpdateResource))
	nodeGroupConfigClientInit()
}

func nodeGroupConfigClientInit() {
	sess := types.GetEKSClient()
	NodeGroupConfigClient = eks.New(sess.Client)
}
