package main

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

// gke container images
type Images struct {
	SRC_IMAGE  string
	DEST_IMAGE string
	IMAGE_NAME string
}

func (i *Images) AddTag(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "add-tag", i.SRC_IMAGE, i.DEST_IMAGE)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		w.Write(data)
	}
}

func (i *Images) Delete(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "delete", i.IMAGE_NAME)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) Describe(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "describe", i.IMAGE_NAME)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) List(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "list")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) ListTags(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "list-tags", i.IMAGE_NAME)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) UnTags(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "untag", i.IMAGE_NAME)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

type Operations struct {
	PROJECT_ID   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ZONE         string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	OPERATION_ID string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	NAME         string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

// create NewClusterManagerClient
func NewClusterManagerClient() (*container.ClusterManagerClient, error) {
	ctx := context.TODO()
	c, err := container.NewClusterManagerClient(ctx, option.WithCredentialsFile("/root/hcp-key.json"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//defer c.Close()
	return c, err
}

// unmarshalling request to Operation struct
func SetOperationRequest(r *http.Request, input interface{}) {
	jsonDataFromHttp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err = json.Unmarshal(jsonDataFromHttp, &input); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}
}

func GetServerConfig(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	var req containerpb.GetServerConfigRequest
	SetOperationRequest(r, &req)
	fmt.Println(req)

	resp, err := c.GetServerConfig(context.TODO(), &req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

// https://pkg.go.dev/cloud.google.com/go/container/apiv1
func (op *Operations) GetOperation(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	SetOperationRequest(r, op)
	req := &containerpb.GetOperationRequest{
		ProjectId:   (*op).PROJECT_ID,
		Zone:        (*op).ZONE,
		OperationId: (*op).OPERATION_ID,
		Name:        (*op).NAME,
	}

	resp, err := c.GetOperation(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func (op *Operations) ListOperations(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	SetOperationRequest(r, op)
	req := &containerpb.ListOperationsRequest{
		ProjectId: (*op).PROJECT_ID,
		Zone:      (*op).ZONE,
	}

	resp, err := c.ListOperations(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func RollbackNodePoolUpgrade(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	var req *containerpb.RollbackNodePoolUpgradeRequest
	SetOperationRequest(r, &req)

	resp, err := c.RollbackNodePoolUpgrade(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func (op *Operations) Wait(w http.ResponseWriter, r *http.Request) {
	SetOperationRequest(r, op)
	fmt.Println((*op).OPERATION_ID)
	cmd := exec.Command("gcloud", "container", "operations", "wait", (*op).OPERATION_ID)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

// gcloud auth configure-docker
func ConfigureDocker(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "configure-docker")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

type Auth struct{}

func (a *Auth) List(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "list")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (a *Auth) Revoke(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "revoke")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

// gcloud docker
type Docker struct {
	AUTHORIZE_ONLY bool
	DOCKER_HOST    string
	SERVER         string
}

func (d *Docker) Docker(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, d)
	args := []string{"docker"}
	if d.AUTHORIZE_ONLY {
		args = append(args, "-a")
	}

	if d.DOCKER_HOST != "" {
		args = append(args, "--docker-host", d.DOCKER_HOST)
	}

	if d.SERVER != "" {
		args = append(args, "-s", d.SERVER)
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func main() {

	var i Images
	http.HandleFunc("/gke/container/images/addTag", i.AddTag)
	http.HandleFunc("/gke/container/images/delete", i.Delete)
	http.HandleFunc("/gke/container/images/describe", i.Describe)
	http.HandleFunc("/gke/container/images/list", i.List)
	http.HandleFunc("/gke/container/images/listTags", i.ListTags)
	http.HandleFunc("/gke/container/images/unTags", i.UnTags)

	var operations Operations
	http.HandleFunc("/gke/container/operations/describe", operations.GetOperation)
	http.HandleFunc("/gke/container/operations/list", operations.ListOperations)
	http.HandleFunc("/gke/container/operations/wait", operations.Wait)
	http.HandleFunc("/gke/container/getServerConfig", GetServerConfig)
	http.HandleFunc("/gke/container/rollbackNodePoolUpgrade", RollbackNodePoolUpgrade)

	var auth Auth
	http.HandleFunc("/gke/auth/configureDocker", ConfigureDocker)
	http.HandleFunc("/gke/auth/list", auth.List)
	http.HandleFunc("/gke/auth/revoke", auth.Revoke)

	var d Docker
	http.HandleFunc("/gke/docker", d.Docker)

	http.ListenAndServe(":3080", nil)
}
