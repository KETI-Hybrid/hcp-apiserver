package main

import (

	// "Hybrid_Cloud/hybridctl/util"

	"Hybrid_Cloud/hcp-apiserver/pkg/handler"
	aksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/aks"
	eksFunc "Hybrid_Cloud/hcp-apiserver/pkg/main/eks"
	"fmt"
	"log"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func AKSHandler() {
	// aks

	// addon
	http.HandleFunc("/addonDisable", aksFunc.AddonDisable)
	http.HandleFunc("/addonEnable", aksFunc.AddonEnable)
	http.HandleFunc("/addonList", aksFunc.AddonList)
	http.HandleFunc("/addonListAvailable", aksFunc.AddonListAvailable)
	http.HandleFunc("/addonShow", aksFunc.AddonShow)
	http.HandleFunc("/addonUpdate", aksFunc.AddonUpdate)

	// pod-identity
	http.HandleFunc("/podIdentityAdd", aksFunc.PodIdentityAdd)
	http.HandleFunc("/podIdentityDelete", aksFunc.PodIdentityDelete)
	http.HandleFunc("/podIdentityList", aksFunc.PodIdentityList)
	http.HandleFunc("/podIdentityExceptionAdd", aksFunc.PodIdentityExceptionAdd)
	http.HandleFunc("/podIdentityExceptionDelete", aksFunc.PodIdentityExceptionDelete)
	http.HandleFunc("/podIdentityExceptionList", aksFunc.PodIdentityExceptionList)
	http.HandleFunc("/podIdentityExceptionUpdate", aksFunc.PodIdentityExceptionUpdate)

	// maintenanceconfiguration
	http.HandleFunc("/maintenanceconfigurationCreateOrUpdate", aksFunc.MaintenanceconfigurationCreateOrUpdate)
	http.HandleFunc("/maintenanceconfigurationDelete", aksFunc.MaintenanceconfigurationDelete)
	http.HandleFunc("/maintenanceconfigurationList", aksFunc.MaintenanceconfigurationList)
	http.HandleFunc("/maintenanceconfigurationShow", aksFunc.MaintenanceconfigurationShow)

	// k8sconfiguration
	http.HandleFunc("/configurationCreate", aksFunc.ConfigurationCreate)
	http.HandleFunc("/configurationDelete", aksFunc.ConfigurationDelete)
	http.HandleFunc("/configurationList", aksFunc.ConfigurationList)
	http.HandleFunc("/configurationShow", aksFunc.ConfigurationShow)

	// connectedk8s
	http.HandleFunc("/connectedk8sConnect", aksFunc.Connectedk8sConnect)
	http.HandleFunc("/connectedk8sDelete", aksFunc.Connectedk8sDelete)
	http.HandleFunc("/connectedk8sDisableFeatures", aksFunc.Connectedk8sDisableFeatures)
	http.HandleFunc("/connectedk8sEnableFeatures", aksFunc.Connectedk8sEnableFeatures)
	http.HandleFunc("/connectedk8sList", aksFunc.Connectedk8sList)
	http.HandleFunc("/connectedk8sProxy", aksFunc.Connectedk8sProxy)
	http.HandleFunc("/connectedk8sShow", aksFunc.Connectedk8sShow)
	http.HandleFunc("/connectedk8sUpdate", aksFunc.Connectedk8sUpdate)
	http.HandleFunc("/connectedk8sUpgrade", aksFunc.Connectedk8sUpgrade)

	// etc
	http.HandleFunc("/aksStart", aksFunc.AksStart)
	http.HandleFunc("/aksStop", aksFunc.AksStop)
	http.HandleFunc("/aksRotateCerts", aksFunc.AksRotateCerts)
	http.HandleFunc("/aksGetOSoptions", aksFunc.AksGetOSoptions)
	http.HandleFunc("/appUp", aksFunc.AppUp)
	http.HandleFunc("/browse", aksFunc.Browse)
	http.HandleFunc("/checkAcr", aksFunc.CheckAcr)
	http.HandleFunc("/getUpgrades", aksFunc.GetUpgrades)
	http.HandleFunc("/getVersions", aksFunc.GetVersions)
	http.HandleFunc("/kanalyze", aksFunc.Kanalyze)
	http.HandleFunc("/kollect", aksFunc.Kollect)
	http.HandleFunc("/nodepoolGetUpgrades", aksFunc.NodepoolGetUpgrades)
	http.HandleFunc("/installCLI", aksFunc.InstallCLI)
}

func EKSHandler() {
	// eks

	// addon
	http.HandleFunc("/createAddon", eksFunc.CreateAddon)
	http.HandleFunc("/listAddon", eksFunc.ListAddon)
	http.HandleFunc("/deleteAddon", eksFunc.DeleteAddon)
	http.HandleFunc("/describeAddon", eksFunc.DescribeAddon)
	http.HandleFunc("/updateAddon", eksFunc.UpdateAddon)
	http.HandleFunc("/describeAddonVersions", eksFunc.DescribeAddonVersions)

	// identity-provider
	http.HandleFunc("/associateIdentityProviderConfig", eksFunc.AssociateIdentityProviderConfig)
	http.HandleFunc("/disassociateIdentityProviderConfig", eksFunc.DisassociateIdentityProviderConfig)
	http.HandleFunc("/describeIdentityProviderConfig", eksFunc.DescribeIdentityProviderConfig)
	http.HandleFunc("/listIdentityProviderConfigs", eksFunc.ListIdentityProviderConfigs)

	// tag
	http.HandleFunc("/tagResource", eksFunc.TagResource)
	http.HandleFunc("/untagResource", eksFunc.UntagResource)
	http.HandleFunc("/listTagsForResource", eksFunc.ListTagsForResource)

	// update
	http.HandleFunc("/updateNodegroupConfig", eksFunc.UpdateNodegroupConfig)
	http.HandleFunc("/updateClusterConfig", eksFunc.UpdateClusterConfig)

	// etc
	http.HandleFunc("/listUpdate", eksFunc.ListUpdate)
	http.HandleFunc("/describeUpdate", eksFunc.DescribeUpdate)
	http.HandleFunc("/associateEncryptionConfig", eksFunc.AssociateEncryptionConfig)
}

func main() {

	fmt.Println("start server")
	AKSHandler()
	EKSHandler()
	http.HandleFunc("/resources/deployment", handler.CreateDeploymentHandler)
	http.ListenAndServe(":8080", nil)
}
