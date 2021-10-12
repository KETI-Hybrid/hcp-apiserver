package main

import (
	"Hybrid_Cluster/hcp-apiserver/converter"
	"Hybrid_Cluster/hybridctl/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	var input util.EksAPIParameter
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
	var input util.EksAPIParameter
	parser(w, req, &input)
	response, err := handler.AksStop(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
}

func aksRotateCerts(w http.ResponseWriter, req *http.Request) {
	var input util.EksAPIParameter
	parser(w, req, &input)
	response, err := handler.AksRotateCerts(input)
	checkErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	defer response.Body.Close()
	w.Write(bytes)
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

	http.ListenAndServe(":8080", nil)
}
