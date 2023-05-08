package fargateprofile

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var FargateProfileClient *eks.EKS

func FargateProfileResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	fargateProfileClientInit()
}
func fargateProfileClientInit() {
	sess := types.GetEKSClient()
	FargateProfileClient = eks.New(sess.Client)
}
