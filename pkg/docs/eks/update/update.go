package update

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var UpdateClient *eks.EKS

func UpdateResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	updateClientInit()
}

func updateClientInit() {
	sess := types.GetEKSClient()
	UpdateClient = eks.New(sess.Client)
}
