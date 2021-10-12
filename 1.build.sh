#!/bin/bash
docker_id="ketidevit"
controller_name="hcp-apiserver"

export GO111MODULE=on
go mod vendor

go build -o build/_output/bin/$controller_name -gcflags all=-trimpath=`pwd` -asmflags all=-trimpath=`pwd` -mod=vendor Hybrid_Cluster/$controller_name && \

docker build -t $docker_id/$controller_name:v0.0.1 build && \
docker push $docker_id/$controller_name:v0.0.1
