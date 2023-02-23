package clusterconfig

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var ClusterConfigClient *eks.EKS

func ClusterConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(UpdateResource))
	clusterConfigClientInit()
}

func clusterConfigClientInit() {
	sess := types.GetEKSClient()
	ClusterConfigClient = eks.New(sess.Client)
}
