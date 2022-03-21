package main

import (

	// "Hybrid_Cloud/hybridctl/util"
	"Hybrid_Cloud/hcp-apiserver/pkg/handler"
	aksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/aks"
	eksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/eks"
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
	mux.HandleFunc("/addonDisable", aksFunc.AddonDisable)
	mux.HandleFunc("/addonEnable", aksFunc.AddonEnable)
	mux.HandleFunc("/addonList", aksFunc.AddonList)
	mux.HandleFunc("/addonListAvailable", aksFunc.AddonListAvailable)
	mux.HandleFunc("/addonShow", aksFunc.AddonShow)
	mux.HandleFunc("/addonUpdate", aksFunc.AddonUpdate)

	// pod-identity
	mux.HandleFunc("/podIdentityAdd", aksFunc.PodIdentityAdd)
	mux.HandleFunc("/podIdentityDelete", aksFunc.PodIdentityDelete)
	mux.HandleFunc("/podIdentityList", aksFunc.PodIdentityList)
	mux.HandleFunc("/podIdentityExceptionAdd", aksFunc.PodIdentityExceptionAdd)
	mux.HandleFunc("/podIdentityExceptionDelete", aksFunc.PodIdentityExceptionDelete)
	mux.HandleFunc("/podIdentityExceptionList", aksFunc.PodIdentityExceptionList)
	mux.HandleFunc("/podIdentityExceptionUpdate", aksFunc.PodIdentityExceptionUpdate)

	// maintenanceconfiguration
	mux.HandleFunc("/maintenanceconfigurationCreateOrUpdate", aksFunc.MaintenanceconfigurationCreateOrUpdate)
	mux.HandleFunc("/maintenanceconfigurationDelete", aksFunc.MaintenanceconfigurationDelete)
	mux.HandleFunc("/maintenanceconfigurationList", aksFunc.MaintenanceconfigurationList)
	mux.HandleFunc("/maintenanceconfigurationShow", aksFunc.MaintenanceconfigurationShow)

	// k8sconfiguration
	mux.HandleFunc("/configurationCreate", aksFunc.ConfigurationCreate)
	mux.HandleFunc("/configurationDelete", aksFunc.ConfigurationDelete)
	mux.HandleFunc("/configurationList", aksFunc.ConfigurationList)
	mux.HandleFunc("/configurationShow", aksFunc.ConfigurationShow)

	// connectedk8s
	mux.HandleFunc("/connectedk8sConnect", aksFunc.Connectedk8sConnect)
	mux.HandleFunc("/connectedk8sDelete", aksFunc.Connectedk8sDelete)
	mux.HandleFunc("/connectedk8sDisableFeatures", aksFunc.Connectedk8sDisableFeatures)
	mux.HandleFunc("/connectedk8sEnableFeatures", aksFunc.Connectedk8sEnableFeatures)
	mux.HandleFunc("/connectedk8sList", aksFunc.Connectedk8sList)
	mux.HandleFunc("/connectedk8sProxy", aksFunc.Connectedk8sProxy)
	mux.HandleFunc("/connectedk8sShow", aksFunc.Connectedk8sShow)
	mux.HandleFunc("/connectedk8sUpdate", aksFunc.Connectedk8sUpdate)
	mux.HandleFunc("/connectedk8sUpgrade", aksFunc.Connectedk8sUpgrade)

	// etc
	mux.HandleFunc("/aksStart", aksFunc.AksStart)
	mux.HandleFunc("/aksStop", aksFunc.AksStop)
	mux.HandleFunc("/aksRotateCerts", aksFunc.AksRotateCerts)
	mux.HandleFunc("/aksGetOSoptions", aksFunc.AksGetOSoptions)
	mux.HandleFunc("/appUp", aksFunc.AppUp)
	mux.HandleFunc("/browse", aksFunc.Browse)
	mux.HandleFunc("/checkAcr", aksFunc.CheckAcr)
	mux.HandleFunc("/getUpgrades", aksFunc.GetUpgrades)
	mux.HandleFunc("/getVersions", aksFunc.GetVersions)
	mux.HandleFunc("/kanalyze", aksFunc.Kanalyze)
	mux.HandleFunc("/kollect", aksFunc.Kollect)
	mux.HandleFunc("/nodepoolGetUpgrades", aksFunc.NodepoolGetUpgrades)
	mux.HandleFunc("/installCLI", aksFunc.InstallCLI)

	// eks

	// addon
	mux.HandleFunc("/createAddon", eksFunc.CreateAddon)
	mux.HandleFunc("/listAddon", eksFunc.ListAddon)
	mux.HandleFunc("/deleteAddon", eksFunc.DeleteAddon)
	mux.HandleFunc("/describeAddon", eksFunc.DescribeAddon)
	mux.HandleFunc("/updateAddon", eksFunc.UpdateAddon)
	mux.HandleFunc("/describeAddonVersions", eksFunc.DescribeAddonVersions)

	// identity-provider
	mux.HandleFunc("/associateIdentityProviderConfig", eksFunc.AssociateIdentityProviderConfig)
	mux.HandleFunc("/disassociateIdentityProviderConfig", eksFunc.DisassociateIdentityProviderConfig)
	mux.HandleFunc("/describeIdentityProviderConfig", eksFunc.DescribeIdentityProviderConfig)
	mux.HandleFunc("/listIdentityProviderConfigs", eksFunc.ListIdentityProviderConfigs)

	// tag
	mux.HandleFunc("/tagResource", eksFunc.TagResource)
	mux.HandleFunc("/untagResource", eksFunc.UntagResource)
	mux.HandleFunc("/listTagsForResource", eksFunc.ListTagsForResource)

	// update
	mux.HandleFunc("/updateNodegroupConfig", eksFunc.UpdateNodegroupConfig)
	mux.HandleFunc("/updateClusterConfig", eksFunc.UpdateClusterConfig)

	// etc
	mux.HandleFunc("/listUpdate", eksFunc.ListUpdate)
	mux.HandleFunc("/describeUpdate", eksFunc.DescribeUpdate)
	mux.HandleFunc("/associateEncryptionConfig", eksFunc.AssociateEncryptionConfig)

	// HCPResource
	mux.HandleFunc("/resources/deployment", handler.CreateDeploymentHandler).Methods("POST")
	mux.HandleFunc("/resources/pod", handler.CreatePodHandler).Methods("POST")

	return mux
}

func main() {

	fmt.Println("start server")
	http.ListenAndServe(":8080", handlerRequests())
}
