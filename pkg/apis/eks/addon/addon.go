package addon

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func AddonResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateReaource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DescribeConfigResource))
	apis.AddResource(router, new(DescribeVersionResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpdateResource))
}
