package identityproviderconfig

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

var IdentityProviderConfigClient *eks.EKS

func IdentityProviderConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AssociateResource))
	apis.AddResource(router, new(DisassociateResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	identityProviderConfigClientInit()
}

func identityProviderConfigClientInit() {
	sess := types.GetEKSClient()
	IdentityProviderConfigClient = eks.New(sess.Client)
}
