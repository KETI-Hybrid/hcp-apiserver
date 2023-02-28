package identityproviderconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var IdentityProviderConfigClient *eks.EKS

func IdentityProviderConfigResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(AssociateResource))
	docs.AddResource(router, new(DisassociateResource))
	docs.AddResource(router, new(DescribeResource))
	docs.AddResource(router, new(ListResource))
	identityProviderConfigClientInit()
}

func identityProviderConfigClientInit() {
	sess := types.GetEKSClient()
	IdentityProviderConfigClient = eks.New(sess.Client)
}
