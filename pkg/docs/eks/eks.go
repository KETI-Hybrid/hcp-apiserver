package eks

import (
	"hcp-apiserver/pkg/apis/eks/addon"
	"hcp-apiserver/pkg/apis/eks/cluster"
	"hcp-apiserver/pkg/apis/eks/clusterconfig"
	"hcp-apiserver/pkg/apis/eks/encryptionconfig"
	"hcp-apiserver/pkg/apis/eks/fargateprofile"
	"hcp-apiserver/pkg/apis/eks/identityproviderconfig"
	"hcp-apiserver/pkg/apis/eks/k8sapi"
	"hcp-apiserver/pkg/apis/eks/nodegroup"
	"hcp-apiserver/pkg/apis/eks/nodegroupconfig"
	"hcp-apiserver/pkg/apis/eks/resource"
	"hcp-apiserver/pkg/apis/eks/update"

	"github.com/julienschmidt/httprouter"
)

func InitEKSEndPoint(router *httprouter.Router) {
	addon.AddonResourceAttach(router)
	cluster.ClusterResourceAttach(router)
	clusterconfig.ClusterConfigResourceAttach(router)
	encryptionconfig.EncryptionConfigResourceAttach(router)
	fargateprofile.FargateProfileResourceAttach(router)
	identityproviderconfig.IdentityProviderConfigResourceAttach(router)
	k8sapi.KubernetesAPIResourceAttach(router)
	nodegroup.NodeGroupResourceAttach(router)
	nodegroupconfig.NodeGroupConfigResourceAttach(router)
	resource.ResourceResourceAttach(router)
	update.UpdateResourceAttach(router)
}
