package k8sapi

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var KubernetesClient *eks.EKS

func KubernetesAPIResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AccessResource))
	k8sClientInit()
}

func k8sClientInit() {
	sess := types.GetEKSClient()
	KubernetesClient = eks.New(sess.Client)
}
