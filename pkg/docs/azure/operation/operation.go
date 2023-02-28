package operation

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var OperationsClient *armcontainerservice.OperationsClient

func OperationsResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(ListResource))
	operationsClientInit()
}

func operationsClientInit() {
	sess := types.GetAKSClient()
	OperationsClient = sess.OperationClient
}
