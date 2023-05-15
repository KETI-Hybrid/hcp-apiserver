package resource

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var ResourceClient *eks.EKS

func ResourceResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(TagResource))
	apis.AddResource(router, new(UntagResource))
	resourceClientInit()
}

func resourceClientInit() {
	sess := types.GetEKSClient()
	ResourceClient = eks.New(sess.Client)
}
