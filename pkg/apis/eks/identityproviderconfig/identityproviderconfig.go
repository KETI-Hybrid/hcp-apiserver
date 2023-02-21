package identityproviderconfig

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func IdentityProviderConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AssociateResource))
	apis.AddResource(router, new(DiassociateResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
}
