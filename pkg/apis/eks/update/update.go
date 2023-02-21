package update

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func UpdateResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
}
