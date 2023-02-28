package google

import (
	"hcp-apiserver/pkg/apis/google/location"
	"hcp-apiserver/pkg/apis/google/usablesubnet"
	"hcp-apiserver/pkg/apis/google/zone"

	"github.com/julienschmidt/httprouter"
)

func InitGKEEndPoint(router *httprouter.Router) {
	location.LocationResourceAttach(router)
	usablesubnet.UsableSubnetworksResourceAttach(router)
	zone.ZoneResourceAttach(router)
}
