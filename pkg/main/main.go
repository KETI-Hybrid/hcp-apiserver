package main

import (

	// "Hybrid_Cloud/hybridctl/util"
	"Hybrid_Cloud/hcp-apiserver/pkg/handler"
	aksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/aks"
	eksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/eks"
	gkeFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/gke"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func handlerRequests() http.Handler {

	mux := mux.NewRouter()
	// aks

	// addon
	mux.HandleFunc("/aks/addon/disable", aksFunc.AddonDisable)
	mux.HandleFunc("/aks/addon/enable", aksFunc.AddonEnable)
	mux.HandleFunc("/aks/addon/list", aksFunc.AddonList)
	mux.HandleFunc("/aks/addon/list-available", aksFunc.AddonListAvailable)
	mux.HandleFunc("/aks/addon/show", aksFunc.AddonShow)
	mux.HandleFunc("/aks/addon/update", aksFunc.AddonUpdate)

	// pod-identity
	mux.HandleFunc("/aks/pod-identity/add", aksFunc.PodIdentityAdd)
	mux.HandleFunc("/aks/pod-identity/delete", aksFunc.PodIdentityDelete)
	mux.HandleFunc("/aks/pod-identity/list", aksFunc.PodIdentityList)
	mux.HandleFunc("/aks/pod-identity/exception/add", aksFunc.PodIdentityExceptionAdd)
	mux.HandleFunc("/aks/pod-identity/exception/delete", aksFunc.PodIdentityExceptionDelete)
	mux.HandleFunc("/aks/pod-identity/exception/list", aksFunc.PodIdentityExceptionList)
	mux.HandleFunc("/aks/pod-identity/exception/update", aksFunc.PodIdentityExceptionUpdate)

	// maintenanceconfiguration
	mux.HandleFunc("/aks/maintenance-configuration/create", aksFunc.MaintenanceconfigurationCreateOrUpdate)
	mux.HandleFunc("/aks/maintenance-configuration/delete", aksFunc.MaintenanceconfigurationDelete)
	mux.HandleFunc("/aks/maintenance-configuration/list", aksFunc.MaintenanceconfigurationList)
	mux.HandleFunc("/aks/maintenance-configuration/show", aksFunc.MaintenanceconfigurationShow)

	// k8sconfiguration
	mux.HandleFunc("/aks/configuration/create", aksFunc.ConfigurationCreate)
	mux.HandleFunc("/aks/configuration/delete", aksFunc.ConfigurationDelete)
	mux.HandleFunc("/aks/configuration/list", aksFunc.ConfigurationList)
	mux.HandleFunc("/aks/configuration/show", aksFunc.ConfigurationShow)

	// connectedk8s
	mux.HandleFunc("/aks/connectedk8s/connect", aksFunc.Connectedk8sConnect)
	mux.HandleFunc("/aks/connectedk8s/delete", aksFunc.Connectedk8sDelete)
	mux.HandleFunc("/aks/connectedk8s/disable-features", aksFunc.Connectedk8sDisableFeatures)
	mux.HandleFunc("/aks/connectedk8s/enable-features", aksFunc.Connectedk8sEnableFeatures)
	mux.HandleFunc("/aks/connectedk8s/list", aksFunc.Connectedk8sList)
	mux.HandleFunc("/aks/connectedk8s/proxy", aksFunc.Connectedk8sProxy)
	mux.HandleFunc("/aks/connectedk8s/show", aksFunc.Connectedk8sShow)
	mux.HandleFunc("/aks/connectedk8s/update", aksFunc.Connectedk8sUpdate)
	mux.HandleFunc("/aks/connectedk8s/upgrade", aksFunc.Connectedk8sUpgrade)

	// etc
	mux.HandleFunc("/aks/start", aksFunc.AksStart)
	mux.HandleFunc("/aks/stop", aksFunc.AksStop)
	mux.HandleFunc("/aks/rotate-certs", aksFunc.AksRotateCerts)
	mux.HandleFunc("/aks/get-os-options", aksFunc.AksGetOSoptions)
	mux.HandleFunc("/aks/app-up", aksFunc.AppUp)
	mux.HandleFunc("/aks/browse", aksFunc.Browse)
	mux.HandleFunc("/aks/check-acr", aksFunc.CheckAcr)
	mux.HandleFunc("/aks/get-upgrades", aksFunc.GetUpgrades)
	mux.HandleFunc("/aks/get-versions", aksFunc.GetVersions)
	mux.HandleFunc("/aks/kanalyze", aksFunc.Kanalyze)
	mux.HandleFunc("/aks/kollect", aksFunc.Kollect)
	mux.HandleFunc("/aks/nodepool-get-upgrades", aksFunc.NodepoolGetUpgrades)
	mux.HandleFunc("/aks/install-cli", aksFunc.InstallCLI)

	// eks

	// addon
	mux.HandleFunc("/eks/addon/create", eksFunc.CreateAddon)
	mux.HandleFunc("/eks/addon/list", eksFunc.ListAddon)
	mux.HandleFunc("/eks/addon/delete", eksFunc.DeleteAddon)
	mux.HandleFunc("/eks/addon/describe", eksFunc.DescribeAddon)
	mux.HandleFunc("/eks/addon/update", eksFunc.UpdateAddon)
	mux.HandleFunc("/eks/addon/describe-versions", eksFunc.DescribeAddonVersions)

	// identity-provider
	mux.HandleFunc("/eks/identity-provider-config/associate", eksFunc.AssociateIdentityProviderConfig)
	mux.HandleFunc("/eks/identity-provider-config/disassociate", eksFunc.DisassociateIdentityProviderConfig)
	mux.HandleFunc("/eks/identity-provider-config/describe", eksFunc.DescribeIdentityProviderConfig)
	mux.HandleFunc("/eks/identity-provider-config/list", eksFunc.ListIdentityProviderConfigs)

	// tag
	mux.HandleFunc("/eks/resource/tag", eksFunc.TagResource)
	mux.HandleFunc("/eks/resource/untag", eksFunc.UntagResource)
	mux.HandleFunc("/eks/resource/list", eksFunc.ListTagsForResource)

	// update
	mux.HandleFunc("/eks/nodegroup-config/update", eksFunc.UpdateNodegroupConfig)
	mux.HandleFunc("/eks/cluster-config/update", eksFunc.UpdateClusterConfig)

	// etc
	mux.HandleFunc("/eks/list/update", eksFunc.ListUpdate)
	mux.HandleFunc("/eks/describe/update", eksFunc.DescribeUpdate)
	mux.HandleFunc("/eks/encryption-config/associate", eksFunc.AssociateEncryptionConfig)

	// gke

	mux.HandleFunc("/gke/container/images/tag/add", gkeFunc.ImagesAddTag)
	mux.HandleFunc("/gke/container/images/delete", gkeFunc.ImagesDelete)
	mux.HandleFunc("/gke/container/images/describe", gkeFunc.ImagesDescribe)
	mux.HandleFunc("/gke/container/images/list", gkeFunc.ImagesList)
	mux.HandleFunc("/gke/container/images/tag/list", gkeFunc.ImagesListTags)
	mux.HandleFunc("/gke/container/images/untags", gkeFunc.ImagesUnTags)

	mux.HandleFunc("/gke/container/operations/describe", gkeFunc.GetOperation)
	mux.HandleFunc("/gke/container/operations/list", gkeFunc.ListOperations)
	mux.HandleFunc("/gke/container/operations/wait", gkeFunc.WaitOperations)
	mux.HandleFunc("/gke/container/server-config/get", gkeFunc.GetServerConfig)
	mux.HandleFunc("/gke/container/nodepool-upgrade/rollback", gkeFunc.RollbackNodePoolUpgrade)

	mux.HandleFunc("/gke/auth/configure-docker", gkeFunc.ConfigureDocker)
	mux.HandleFunc("/gke/auth/list", gkeFunc.AuthList)
	mux.HandleFunc("/gke/auth/revoke", gkeFunc.AuthRevoke)
	mux.HandleFunc("/gke/auth/login", gkeFunc.AuthLogin)

	mux.HandleFunc("/gke/docker", gkeFunc.GDocker)

	mux.HandleFunc("/gke/config/set", gkeFunc.ConfigSet)

	// HCPResource
	mux.HandleFunc("/resources/deployment", handler.CreateDeploymentHandler).Methods("POST")
	mux.HandleFunc("/resources/pod", handler.CreatePodHandler).Methods("POST")

	return mux
}

func main() {

	fmt.Println("start server")
	http.ListenAndServe(":8080", handlerRequests())
}
