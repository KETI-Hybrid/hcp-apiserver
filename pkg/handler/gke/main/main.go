package main

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func main() {
	/* container operation list
	client, err := container.NewClient(context.TODO(), "keti-container")
	if err != nil {
		fmt.Println(err)
	}

	var op []*container.Op
	op, err = client.Operations(context.TODO(), "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(op)
	for _, o := range op {
		fmt.Println(*o)
	}
	*/

	// container operation describe
	// Reference: https://pkg.go.dev/cloud.google.com/go/container/apiv1
	ctx := context.TODO()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	req := &containerpb.GetOperationRequest{
		ProjectId:   "keti-container",
		Zone:        "us-central1-a",
		OperationId: "operation-1648309236003-34160983",
		Name:        "operation-1648309236003-34160983",
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*resp)
}
