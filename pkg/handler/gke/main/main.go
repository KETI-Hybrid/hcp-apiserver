package main

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

type Operations struct {
	ProjectId   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Zone        string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	OperationId string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Name        string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
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
func SetOperationRequest(req *http.Request, op *Operations) {
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err = json.Unmarshal(jsonDataFromHttp, &op); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*op)
	}
}

// https://pkg.go.dev/cloud.google.com/go/container/apiv1
func (op *Operations) GetOperation(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	SetOperationRequest(r, op)
	req := &containerpb.GetOperationRequest{
		ProjectId:   (*op).ProjectId,
		Zone:        (*op).Zone,
		OperationId: (*op).OperationId,
		Name:        (*op).Name,
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
		ProjectId: (*op).ProjectId,
		Zone:      (*op).Zone,
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

func main() {
	var operations Operations
	http.HandleFunc("/gke/container/operations/describe", operations.GetOperation)
	http.HandleFunc("/gke/container/operations/list", operations.ListOperations)
	http.ListenAndServe(":3080", nil)
}
