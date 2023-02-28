package encryptionconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var EncryptionConfigClient *eks.EKS

func EncryptionConfigResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(AssociateResource))
	encryptionConfigClientInit()
}

func encryptionConfigClientInit() {
	sess := types.GetEKSClient()
	EncryptionConfigClient = eks.New(sess.Client)
}
