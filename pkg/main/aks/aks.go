package aks

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/handler"
	"Hybrid_Cluster/hcp-apiserver/pkg/util"
	"fmt"
	"io"
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

func AksStart(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.AksStart(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	fmt.Println(string(bytes))
	w.Write(bytes)
}

func AksStop(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.AksStop(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func AksRotateCerts(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.AksRotateCerts(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func AksGetOSoptions(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.AksGetOSoptions(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func MaintenanceconfigurationCreateOrUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationCreateOrUpdate(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func MaintenanceconfigurationList(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationList(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func MaintenanceconfigurationDelete(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationDelete(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func MaintenanceconfigurationShow(w http.ResponseWriter, req *http.Request) {
	var input util.EKSAPIParameter
	util.Parser(w, req, &input)
	response, err := handler.MaintenanceconfigurationShow(input)
	CheckErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func AddonDisable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "disable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func AddonEnable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "enable", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func AddonList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "list", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName)
	// errMsg := ExampleCmd_StderrPipe(cmd)
	// fmt.Println(string(errMsg))
	errMsg := ExampleCmd_StderrPipe(cmd)
	output, err := cmd.Output()
	// errMsg := ExampleCmd_StderrPipe(cmd)
	w.Write(errMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	w.Write(output)
	// }
}

func AddonListAvailable(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "list-available")
	errMsg := ExampleCmd_StderrPipe(cmd)
	fmt.Println(string(errMsg))
	output, err := cmd.Output()
	// errMsg := ExampleCmd_StderrPipe(cmd)
	// w.Write(errMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func ExampleCmd_StderrPipe(cmd *exec.Cmd) []byte {
	// cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	return slurp
	// if err := cmd.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
}

func AddonShow(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "show", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func AddonUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAddon
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "addon", "update", "--name", input.ClusterName, "--resource-group", input.ResourceGroupName, "--addon", input.Addon)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func PodIdentityAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
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

func PodIdentityDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "delete", "--cluster-name", input.ClusterName, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func PodIdentityList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
		w.Write(output)
	}
}
func PodIdentityExceptionAdd(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodLabels, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func PodIdentityExceptionDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "delete", "--cluster-name", input.ClusterName, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func PodIdentityExceptionList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "exception", "list", "--cluster-name", input.ClusterName, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func PodIdentityExceptionUpdate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSPodIdentity
	util.Parser(w, req, &input)
	cmd := exec.Command("az", "aks", "pod-identity", "add", "--cluster-name", input.ClusterName, "--pod-labels", input.PodLabels, "--name", input.Name, "--namespace", input.Namespace, "--resource-group", input.ResourceGroupName)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}
func AppUp(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
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

func Browse(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
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

func CheckAcr(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
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

func GetUpgrades(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
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

func GetVersions(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
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

func Kanalyze(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
	args := []string{"aks", "kanalyze", "--name", input.Name, "-g", input.ResourceGroup}
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func NodepoolGetUpgrades(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, &input)
	args := []string{"aks", "nodepool", "get-upgrades", "--cluster-name", input.Name, "-g", input.ResourceGroup, "--nodepool-name", input.NodepoolName}

	if input.Subscription != "" {
		args = append(args, "--subscription", input.Subscription)
	}
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func InstallCLI(w http.ResponseWriter, req *http.Request) {
	var input util.AKSInstallCLI
	util.Parser(w, req, &input)
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
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func ConnectedDisableFeatures(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, input)
	args := []string{"connectedk8s", "disable-features", "--name", input.Name, "-g", input.ResourceGroup, "--features"}
	for i := range input.Features {
		f := input.Features[i]
		args = append(args, f)
	}
	cmd := exec.Command("az", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(output)
	}
}

func ConnectedList(w http.ResponseWriter, req *http.Request) {
	var input util.AKSAPIParameter
	util.Parser(w, req, input)
	args := []string{"connectedk8s", "list", "-g", input.ResourceGroup}
	cmd := exec.Command("az", args...)
	output, _ := cmd.Output()
	w.Write(output)
}

func ConfigurationCreate(w http.ResponseWriter, req *http.Request) {
	var input util.AKSk8sConfiguration
	util.Parser(w, req, input)
	args := []string{"k8sconfiguration", "create", "-g", input.ResourceGroup, "-c", input.ClusterName, "--cluster-type", input.ClusterType, "-n", input.Name, "-u", input.RepositoryURL, "--scope", input.Scope}
	cmd := exec.Command("az", args...)
	output, _ := cmd.CombinedOutput()
	w.Write(output)
}

func ConfigurationDelete(w http.ResponseWriter, req *http.Request) {
	var input util.AKSk8sConfiguration
	util.Parser(w, req, input)
	fmt.Println(input.ClusterType)
	// args := []string{"k8sconfiguration", "delete", "-g", input.ResourceGroup, "-c", input.ClusterName, "--cluster-type", "managedClusters", "-n", input.Name}
	// cmd := exec.Command("az", args...)
	cmd := exec.Command("az", "k8sconfiguration", "delete", "-g", input.ResourceGroup, "-c", input.ClusterName, "--cluster-type", input.ClusterType, "-n", input.Name)
	output, _ := cmd.Output()
	w.Write(output)
}
