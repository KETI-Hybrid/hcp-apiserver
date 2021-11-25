package main

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/converter"
	// "Hybrid_Cluster/hybridctl/util"
	util "Hybrid_Cluster/hcp-apiserver/pkg/util"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "Hybrid_Cluster/hcp-apiserver/pkg/handler"
	aksFunc "Hybrid_Cluster/hcp-apiserver/pkg/main/aks"
	eksFunc "Hybrid_Cluster/hcp-apiserver/pkg/main/eks"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func join(w http.ResponseWriter, req *http.Request) {

	fmt.Println("---ok---")
	clusterInfo := converter.ClusterInfo{}
	util.Parser(w, req, &clusterInfo)
	var info = converter.ClusterInfo{
		PlatformName: clusterInfo.PlatformName,
		ClusterName:  clusterInfo.ClusterName,
	}
	handler.Join(info)
	w.Header().Set("Content-Type", "application/json")
}

func unjoin(w http.ResponseWriter, req *http.Request) {
	clusterInfo := converter.ClusterInfo{}
	util.Parser(w, req, &clusterInfo)
	var info = converter.ClusterInfo{
		PlatformName: clusterInfo.PlatformName,
		ClusterName:  clusterInfo.ClusterName,
	}
	handler.Unjoin(info)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/unjoin", unjoin)
	http.HandleFunc("/createAddon", eksFunc.CreateAddon)
	http.HandleFunc("/listAddon", eksFunc.ListAddon)
	http.HandleFunc("/deleteAddon", eksFunc.DeleteAddon)
	http.HandleFunc("/describeAddon", eksFunc.DescribeAddon)
	http.HandleFunc("/updateNodegroupConfig", eksFunc.UpdateNodegroupConfig)
	http.HandleFunc("/describeAddonVersions", eksFunc.DescribeAddonVersions)
	http.HandleFunc("/updateAddon", eksFunc.UpdateAddon)
	http.HandleFunc("/listUpdate", eksFunc.ListUpdate)
	http.HandleFunc("/describeUpdate", eksFunc.DescribeUpdate)
	http.HandleFunc("/listTagsForResource", eksFunc.ListTagsForResource)
	http.HandleFunc("/associateEncryptionConfig", eksFunc.AssociateEncryptionConfig)
	http.HandleFunc("/associateIdentityProviderConfig", eksFunc.AssociateIdentityProviderConfig)
	http.HandleFunc("/disassociateIdentityProviderConfig", eksFunc.DisassociateIdentityProviderConfig)
	http.HandleFunc("/describeIdentityProviderConfig", eksFunc.DescribeIdentityProviderConfig)
	http.HandleFunc("/listIdentityProviderConfigs", eksFunc.ListIdentityProviderConfigs)
	http.HandleFunc("/tagResource", eksFunc.TagResource)
	http.HandleFunc("/untagResource", eksFunc.UntagResource)
	http.HandleFunc("/updateClusterConfig", eksFunc.UpdateClusterConfig)

	http.HandleFunc("/aksStart", aksFunc.AksStart)
	http.HandleFunc("/aksStop", aksFunc.AksStop)
	http.HandleFunc("/aksRotateCerts", aksFunc.AksRotateCerts)
	http.HandleFunc("/aksGetOSoptions", aksFunc.AksGetOSoptions)
	http.HandleFunc("/maintenanceconfigurationCreateOrUpdate", aksFunc.MaintenanceconfigurationCreateOrUpdate)
	// maintenanceconfiguration add + update
	http.HandleFunc("/maintenanceconfigurationDelete", aksFunc.MaintenanceconfigurationDelete)
	http.HandleFunc("/maintenanceconfigurationList", aksFunc.MaintenanceconfigurationList)
	http.HandleFunc("/maintenanceconfigurationShow", aksFunc.MaintenanceconfigurationShow)
	http.HandleFunc("/addonDisable", aksFunc.AddonDisable)
	http.HandleFunc("/addonEnable", aksFunc.AddonEnable)
	http.HandleFunc("/addonList", aksFunc.AddonList)
	http.HandleFunc("/addonListAvailable", aksFunc.AddonListAvailable)
	http.HandleFunc("/addonShow", aksFunc.AddonShow)
	http.HandleFunc("/addonUpdate", aksFunc.AddonUpdate)
	http.HandleFunc("/podIdentityAdd", aksFunc.PodIdentityAdd)
	http.HandleFunc("/podIdentityDelete", aksFunc.PodIdentityDelete)
	http.HandleFunc("/podIdentityList", aksFunc.PodIdentityList)
	http.HandleFunc("/podIdentityExceptionAdd", aksFunc.PodIdentityExceptionAdd)
	http.HandleFunc("/podIdentityExceptionDelete", aksFunc.PodIdentityExceptionDelete)
	http.HandleFunc("/podIdentityExceptionList", aksFunc.PodIdentityExceptionList)
	http.HandleFunc("/podIdentityExceptionUpdate", aksFunc.PodIdentityExceptionUpdate)
	http.HandleFunc("/appUp", aksFunc.AppUp)
	http.HandleFunc("/browse", aksFunc.Browse)
	http.HandleFunc("/checkAcr", aksFunc.CheckAcr)
	http.HandleFunc("/getUpgrades", aksFunc.GetUpgrades)
	http.HandleFunc("/getVersions", aksFunc.GetVersions)
	http.HandleFunc("/kanalyze", aksFunc.Kanalyze)
	http.HandleFunc("/nodepoolGetUpgrades", aksFunc.NodepoolGetUpgrades)
	http.HandleFunc("/installCLI", aksFunc.InstallCLI)
	// http.HandleFunc("/connectedConnect", connectedConnect)
	// http.HandleFunc("/connectedk8sDelete", connectedk8sDelete)
	http.HandleFunc("/connectedDisableFeatures", aksFunc.ConnectedDisableFeatures)
	// http.HandleFunc("/connectedEnableFeatures", connectedEnableFeatures)
	http.HandleFunc("/connectedList", aksFunc.ConnectedList)
	// http.HandleFunc("/connectedProxy", connectedProxy)
	// http.HandleFunc("/connectedShow", connectedShow)
	// http.HandleFunc("/connectedUpdate", connectedUpdate)
	// http.HandleFunc("/connectedUpgrade", connectedUpgrade)
	http.HandleFunc("/configurationCreate", aksFunc.ConfigurationCreate)
	http.HandleFunc("/configurationDelete", aksFunc.ConfigurationDelete)
	// http.HandleFunc("/configurationCreate", configurationList)
	// http.HandleFunc("/configurationCreate", configurationShow)
	http.ListenAndServe(":8080", nil)
}

func init() {
	os.Setenv("ClientId", "5a7002e5-86e6-42c8-a844-976f4b95760d")
	os.Setenv("ClientSecret", "I.E76p.jvKWFJxf3Ufqf1H_c66--ww53J2")
	os.Setenv("SubscriptionId", "ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73")
	os.Setenv("TenantId", "c8ea91b5-6aac-4c5c-ae34-9717a872159f")
}
