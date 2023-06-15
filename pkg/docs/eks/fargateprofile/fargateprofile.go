package fargateprofile

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var FargateProfileClient *eks.EKS

func FargateProfileResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	fargateProfileClientInit()
}
func fargateProfileClientInit() {
	sess := types.GetEKSClient()
	FargateProfileClient = eks.New(sess.Client)
}
