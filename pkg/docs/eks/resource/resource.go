package resource

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var ResourceClient *eks.EKS

func ResourceResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(TagResource))
	docs.AddResource(router, new(UntagResource))
	resourceClientInit()
}

func resourceClientInit() {
	sess := types.GetEKSClient()
	ResourceClient = eks.New(sess.Client)
}
