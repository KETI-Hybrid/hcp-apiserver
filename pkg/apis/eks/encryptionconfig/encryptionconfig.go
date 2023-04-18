package encryptionconfig

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func EncryptionConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AssociateResource))
}
