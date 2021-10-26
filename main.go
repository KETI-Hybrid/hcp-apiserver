package main

import (
	"Hybrid_Cluster/hcp-apiserver/converter"
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	handler "Hybrid_Cluster/hcp-apiserver/handler"

	"github.com/aws/aws-sdk-go/service/eks"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func parser(w http.ResponseWriter, req *http.Request, input interface{}) {
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	fmt.Printf(string(jsonDataFromHttp))
	json.Unmarshal(jsonDataFromHttp, input)
	defer req.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
}

func join(w http.ResponseWriter, req *http.Request) {

	fmt.Println("---ok---")
	clusterInfo := converter.ClusterInfo{}
	parser(w, req, &clusterInfo)
	var info = converter.ClusterInfo{
		PlatformName: clusterInfo.PlatformName,
		ClusterName:  clusterInfo.ClusterName,
	}
	handler.Join(info)
	w.Header().Set("Content-Type", "application/json")
}

func unjoin(w http.ResponseWriter, req *http.Request) {
	clusterInfo := converter.ClusterInfo{}
	parser(w, req, &clusterInfo)
	var info = converter.ClusterInfo{
		PlatformName: clusterInfo.PlatformName,
		ClusterName:  clusterInfo.ClusterName,
	}
	handler.Unjoin(info)
	w.Header().Set("Content-Type", "application/json")
}

func createAddon(w http.ResponseWriter, req *http.Request) {

	var createAddonInput eks.CreateAddonInput

	parser(w, req, &createAddonInput)
	out, err := handler.CreateAddon(createAddonInput)
	var jsonData []byte
	if err != nil {
		log.Println(err)
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func deleteAddon(w http.ResponseWriter, req *http.Request) {

	var deleteAddonInput eks.DeleteAddonInput

	parser(w, req, &deleteAddonInput)
	out, err := handler.DeleteAddon(deleteAddonInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func describeAddon(w http.ResponseWriter, req *http.Request) {

	var describeAddonInput eks.DescribeAddonInput

	parser(w, req, &describeAddonInput)
	out, err := handler.DescribeAddon(describeAddonInput)
	checkErr(err)
	if err != nil {
		jsonData, _ := json.Marshal(err)
		w.Write(jsonData)
	} else {
		jsonData, _ := json.Marshal(&out)
		fmt.Println(jsonData)
		w.Write(jsonData)
	}
}

func describeAddonVersions(w http.ResponseWriter, req *http.Request) {

	var describeAddonVersionsInput eks.DescribeAddonVersionsInput

	parser(w, req, &describeAddonVersionsInput)
	out, err := handler.DescribeAddonVersions(describeAddonVersionsInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write(jsonData)

}

func listAddon(w http.ResponseWriter, req *http.Request) {

	var listAddonInput eks.ListAddonsInput

	parser(w, req, &listAddonInput)
	out, err := handler.ListAddon(listAddonInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))

}

func updateAddon(w http.ResponseWriter, req *http.Request) {

	var updateAddonInput eks.UpdateAddonInput

	parser(w, req, &updateAddonInput)
	out, err := handler.UpdateAddon(updateAddonInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func listUpdate(w http.ResponseWriter, req *http.Request) {
	var listUpdateInput eks.ListUpdatesInput

	parser(w, req, &listUpdateInput)
	out, err := handler.ListUpdate(listUpdateInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func describeUpdate(w http.ResponseWriter, req *http.Request) {
	var describeUpdateInput eks.DescribeUpdateInput

	parser(w, req, &describeUpdateInput)
	out, err := handler.DescribeUpdate(describeUpdateInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func listTagsForResource(w http.ResponseWriter, req *http.Request) {
	var listTagsForResourceInput eks.ListTagsForResourceInput

	parser(w, req, &listTagsForResourceInput)
	out, err := handler.ListTagsForResource(listTagsForResourceInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func associateIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var associateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput

	parser(w, req, &associateIdentityProviderConfigInput)
	out, err := handler.AssociateIdentityProviderConfig(associateIdentityProviderConfigInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func disassociateIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var disassociateIdentityProviderConfigInput eks.DisassociateIdentityProviderConfigInput

	parser(w, req, &disassociateIdentityProviderConfigInput)
	out, err := handler.DisassociateIdentityProviderConfig(disassociateIdentityProviderConfigInput)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func describeIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var input eks.DescribeIdentityProviderConfigInput

	parser(w, req, &input)
	out, err := handler.DescribeIdentityProviderConfig(input)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func listIdentityProviderConfigs(w http.ResponseWriter, req *http.Request) {
	var input eks.ListIdentityProviderConfigsInput

	parser(w, req, &input)
	out, err := handler.ListIdentityProviderConfigs(input)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func tagResource(w http.ResponseWriter, req *http.Request) {
	var input eks.TagResourceInput

	parser(w, req, &input)
	out, err := handler.TagResource(input)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func untagResource(w http.ResponseWriter, req *http.Request) {
	var input eks.UntagResourceInput

	parser(w, req, &input)
	out, err := handler.UntagResource(input)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func updateClusterConfig(w http.ResponseWriter, req *http.Request) {
	var input eks.UpdateClusterConfigInput

	parser(w, req, &input)
	out, err := handler.UpdateClusterConfig(input)
	checkErr(err)
	jsonData, _ := json.Marshal(&out)
	w.Write([]byte(jsonData))
}

func aksStart(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.AksStart(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	fmt.Println(string(bytes))
	w.Write(bytes)
}

func aksStop(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.AksStop(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func aksRotateCerts(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.AksRotateCerts(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func aksGetOSoptions(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.AksGetOSoptions(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func maintenanceconfigurationCreateOrUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationCreateOrUpdate(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func maintenanceconfigurationList(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationList(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func maintenanceconfigurationDelete(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationDelete(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func maintenanceconfigurationShow(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationShow(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func addonDisable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "disable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func addonEnable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "enable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func addonList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "list", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func addonListAvailable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "list-available")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func addonShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "show", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func addonUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "update", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func podIdentityAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	args := []string{"az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--identity-resource-id", input.IdentityResourceID, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName}
	if input.Name != "" {
		args = append(args, "--name", input.Name)
	}
	cmd := exec.Command("podIdentityAdd", args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func podIdentityDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "delete", "--cluster-name", input.ClusterName, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func podIdentityList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
		w.Write(output)
	}
}
func podIdentityExceptionAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodLabels, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func podIdentityExceptionDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "delete", "--cluster-name", input.ClusterName, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func podIdentityExceptionList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func podIdentityExceptionUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodLabels, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func appUp(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	// args := []string{"az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--identity-resource-id", input.IdentityResourceID, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName}
	// if input.Name != "" {
	// 	args = append(args, "--name", input.Name)
	// }
	// cmd := exec.Command("podIdentityAdd", args...)
	args := []string{"aks", "app", "up"}
	if input.Acr != "" {
		args = append(args, "--acr", input.Acr)
	}
	if input.AksCluster != "" {
		args = append(args, "--aks-cluster", input.AksCluster)
	}
	if input.BranchName != "" {
		args = append(args, "--branch-name", input.BranchName)
	}
	if input.DoNotWait != "" {
		args = append(args, "--do-not-wait", input.DoNotWait)
	}
	if input.BranchName != "" {
		args = append(args, "--port", input.Port)
	}
	if input.Repository != "" {
		args = append(args, "--repository", input.Repository)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func browse(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	args := []string{"aks", "browse", "--name", input.Name, "-g", input.ResourceGroup}
	if input.DisableBrowser {
		args = append(args, "--disable-browser")
	}
	if input.ListenAddress != "" {
		args = append(args, "--listen-address", input.ListenAddress)
	}
	if input.ListenPort != "" {
		args = append(args, "--listen-port", input.ListenPort)
	}
	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func checkAcr(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	args := []string{"aks", "check-acr", "--name", input.Name, "-g", input.ResourceGroup, "--acr", input.Acr}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func getUpgrades(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	args := []string{"aks", "get-upgrades", "--name", input.Name, "-g", input.ResourceGroup}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func getVersions(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	args := []string{"aks", "get-versions", "-l", input.Location}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func kanalyze(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	parser(w, req, &input)
	args := []string{"aks", "kanalyze", "--name", input.Name, "-g", input.ResourceGroup}
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/unjoin", unjoin)
	http.HandleFunc("/listAddon", listAddon)
	http.HandleFunc("/updateAddon", updateAddon)
	http.HandleFunc("/listUpdate", listUpdate)
	http.HandleFunc("/describeUpdate", describeUpdate)
	http.HandleFunc("/listTagsForResource", listTagsForResource)
	http.HandleFunc("/associateIdentityProvicerConfig", associateIdentityProviderConfig)
	http.HandleFunc("/disassociateIdentityProviderConfig", disassociateIdentityProviderConfig)
	http.HandleFunc("/describeIdentityProviderConfig", describeIdentityProviderConfig)
	http.HandleFunc("/listIdentityProviderConfigs", listIdentityProviderConfigs)
	http.HandleFunc("/tagResource", tagResource)
	http.HandleFunc("/untagResource", untagResource)
	http.HandleFunc("/updateClusterConfig", updateClusterConfig)
	http.HandleFunc("/aksStart", aksStart)
	http.HandleFunc("/aksStop", aksStop)
	http.HandleFunc("/aksRotateCerts", aksRotateCerts)
	http.HandleFunc("/aksGetOSoptions", aksGetOSoptions)
	http.HandleFunc("/maintenanceconfigurationCreateOrUpdate", maintenanceconfigurationCreateOrUpdate)
	// maintenanceconfiguration add + update
	http.HandleFunc("/maintenanceconfigurationDelete", maintenanceconfigurationDelete)
	http.HandleFunc("/maintenanceconfigurationList", maintenanceconfigurationList)
	http.HandleFunc("/maintenanceconfigurationShow", maintenanceconfigurationShow)
	http.HandleFunc("/addonDisable", addonDisable)
	http.HandleFunc("/addonEnable", addonEnable)
	http.HandleFunc("/addonList", addonList)
	http.HandleFunc("/addonListAvailable", addonListAvailable)
	http.HandleFunc("/addonShow", addonShow)
	http.HandleFunc("/addonUpdate", addonUpdate)
	http.HandleFunc("/podIdentityAdd", podIdentityAdd)
	http.HandleFunc("/podIdentityDelete", podIdentityDelete)
	http.HandleFunc("/podIdentityList", podIdentityList)
	http.HandleFunc("/podIdentityExceptionAdd", podIdentityExceptionAdd)
	http.HandleFunc("/podIdentityExceptionDelete", podIdentityExceptionDelete)
	http.HandleFunc("/podIdentityExceptionList", podIdentityExceptionList)
	http.HandleFunc("/podIdentityExceptionUpdate", podIdentityExceptionUpdate)
	http.HandleFunc("/appUp", appUp)
	http.HandleFunc("/browse", browse)
	http.HandleFunc("/checkAcr", checkAcr)
	http.HandleFunc("/getUpgrades", getUpgrades)
	http.HandleFunc("/getVersions", getVersions)
	http.HandleFunc("/kanalyze", kanalyze)
	http.ListenAndServe(":8080", nil)
}

func init() {
	os.Setenv("ClientId", "5a7002e5-86e6-42c8-a844-976f4b95760d")
	os.Setenv("ClientSecret", "I.E76p.jvKWFJxf3Ufqf1H_c66--ww53J2")
	os.Setenv("SubscriptionId", "ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73")
	os.Setenv("TenantId", "c8ea91b5-6aac-4c5c-ae34-9717a872159f")
}