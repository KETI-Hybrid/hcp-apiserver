package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/container"
)

func main() {

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

}
