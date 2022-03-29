package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	/*
		req := &containerpb.GetOperationRequest{
			ProjectId:   "keti-container",
			Zone:        "us-central1-a",
			OperationId: "operation-1648309236003-34160983",
			Name:        "operation-1648309236003-34160983",
		}
	*/

	resp, err := c.GetOperation(context.TODO(), req)
	defer c.Close()

	if err != nil {
		log.Println(err)
		bytes, err2 := json.Marshal(err)
		if err2 != nil {
			log.Println(err2)
		}
		w.Write(bytes)
	} else {
		bytes, err := json.Marshal(&resp)
		if err != nil {
			log.Println(err)
		}
		w.Write(bytes)
	}
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

	fmt.Println(resp)
	if err != nil {
		log.Println(err)
		bytes, err2 := json.Marshal(err)
		if err2 != nil {
			log.Println(err2)
		}
		w.Write(bytes)
	} else {
		bytes, err := json.Marshal(&resp)
		if err != nil {
			log.Println(err)
		}
		w.Write(bytes)
	}
}

func main() {
	var operations Operations
	http.HandleFunc("/gke/container/operations/describe", operations.GetOperation)
	http.HandleFunc("/gke/container/operations/list", operations.ListOperations)
	http.ListenAndServe(":3080", nil)
}
