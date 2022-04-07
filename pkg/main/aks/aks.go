package aks

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/handler"
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// addon
func AddonDisable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "disable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon.Addon)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func AddonEnable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "enable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon.Addon)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func AddonList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "list", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func AddonListAvailable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "list-available")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func AddonShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "show", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon.Addon)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func AddonUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "addon", "update", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon.Addon)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// pod-identity
func PodIdentityAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--identity-resource-id", input.PodIdentity.IdentityResourceID, "--namespace", input.PodIdentity.Namespace, "--resource-group", input.ResourceGroupName}
	if input.PodIdentity.Name != "" {
		args = append(args, "--name", input.PodIdentity.Name)
	}

	if input.PodIdentity.BindingSelector != "" {
		args = append(args, "--binding-selector", input.PodIdentity.BindingSelector)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func PodIdentityDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "delete", "--cluster-name", input.ClusterName, "--name", input.PodIdentity.Name, "--namespace", input.PodIdentity.Namespace, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func PodIdentityList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func PodIdentityExceptionAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodIdentity.PodLabels, "--namespace", input.PodIdentity.Namespace, "--resource-group", input.ResourceGroupName}
	if input.PodIdentity.Name != "" {
		args = append(args, "--name", input.PodIdentity.Name)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
func PodIdentityExceptionDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "delete", "--cluster-name", input.ClusterName, "--name", input.PodIdentity.Name, "--namespace", input.PodIdentity.Namespace, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
func PodIdentityExceptionList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func PodIdentityExceptionUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodIdentity.PodLabels, "--name", input.PodIdentity.Name, "--namespace", input.PodIdentity.Namespace, "--resource-group", input.ResourceGroupName)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// maintenanceconfiguration
func MaintenanceconfigurationCreateOrUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.MaintenanceconfigurationCreateOrUpdate(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func MaintenanceconfigurationList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.MaintenanceconfigurationList(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func MaintenanceconfigurationDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.MaintenanceconfigurationDelete(input)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.Status))
}

func MaintenanceconfigurationShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.MaintenanceconfigurationShow(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// k8sconfiguration
func ConfigurationCreate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"k8sconfiguration", "create", "-g", input.ResourceGroupName, "-c", input.ClusterName, "--cluster-type", input.K8sConfiguration.ClusterType, "-n", input.K8sConfiguration.Name, "-u", input.K8sConfiguration.RepositoryURL, "--scope", input.K8sConfiguration.Scope}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ConfigurationDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "k8sconfiguration", "delete", "-g", input.ResourceGroupName, "-c", input.ClusterName, "--cluster-type", input.K8sConfiguration.ClusterType, "-n", input.K8sConfiguration.Name, "--yes")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
func ConfigurationShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "k8sconfiguration", "show", "-g", input.ResourceGroupName, "-c", input.ClusterName, "--cluster-type", input.K8sConfiguration.ClusterType, "-n", input.K8sConfiguration.Name)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ConfigurationList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	cmd := exec.Command("az", "k8s-configuration", "list", "-g", input.ResourceGroupName, "-c", input.ClusterName, "-t", input.K8sConfiguration.ClusterType)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// connectedk8s

func Connectedk8sConnect(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "connect", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "delete", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sDisableFeatures(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "disable-features", "-n", input.ClusterName, "-g", input.ResourceGroupName, "--features"}
	for i := range input.Features {
		f := input.Features[i]
		args = append(args, f)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sEnableFeatures(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "enable-features", "-n", input.ClusterName, "-g", input.ResourceGroupName, "--features"}
	for i := range input.Features {
		f := input.Features[i]
		args = append(args, f)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"connectedk8s", "list"}
	if input.ResourceGroupName != "" {
		args = append(args, "-g", input.ResourceGroupName)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sProxy(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "proxy", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "show", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "update", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Connectedk8sUpgrade(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, input)
	args := []string{"connectedk8s", "upgrade", "-n", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// etc
func AksStart(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.AksStart(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}

func AksStop(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.AksStop(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func AksRotateCerts(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.AksRotateCerts(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func AksGetOSoptions(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	response, err := handler.AksGetOSoptions(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func AppUp(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
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
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Browse(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "browse", "--name", input.ClusterName, "-g", input.ResourceGroupName}
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
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func CheckAcr(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "check-acr", "--name", input.ClusterName, "-g", input.ResourceGroupName, "--acr", input.Acr}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func GetUpgrades(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "get-upgrades", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func GetVersions(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "get-versions", "-l", input.Location}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	fmt.Println(args)
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Kanalyze(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "kanalyze", "--name", input.ClusterName, "-g", input.ResourceGroupName}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Kollect(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "kollect", "-n", input.ClusterName, "-g", input.ResourceGroupName, "--storage-account", input.StorageAccount}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func NodepoolGetUpgrades(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(req, &input)
	args := []string{"aks", "nodepool", "get-upgrades", "--cluster-name", input.ClusterName, "-g", input.ResourceGroupName, "--nodepool-name", input.NodepoolName}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func InstallCLI(w http.ResponseWriter, req *http.Request) {
	var input util.AKSInstallCLI
	util.Parser(req, &input)
	args := []string{"aks", "install-cli"}

	if input.BaseSrcURL != "" {
		args = append(args, "--base-src-url", input.Subscription)
	}
	if input.ClientVersion != "" {
		args = append(args, "--client-version", input.Subscription)
	}
	if input.KubeloginBaseSrcURL != "" {
		args = append(args, "--kubelogin-base-src-url", input.Subscription)
	}
	if input.KubeloginBaseSrcURL != "" {
		args = append(args, "--kubelogin-install-location", input.Subscription)
	}
	if input.KubeloginVersion != "" {
		args = append(args, "--kubelogin-version", input.Subscription)
	}
	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	cmd := exec.Command("az", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
