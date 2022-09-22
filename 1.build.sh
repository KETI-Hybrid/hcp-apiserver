#!/bin/bash

docker_id="ketidevit2"
controller_name="hcp-apiserver"

export GO111MODULE=on
go mod vendor

go build -o build/_output/bin/$controller_name -gcflags all=-trimpath=`pwd` -asmflags all=-trimpath=`pwd` -mod=vendor github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/main && \

docker build -t $docker_id/$controller_name:v0.0.2 build && \
docker push $docker_id/$controller_name:v0.0.2


#0.0.1 >> 운용중인 api server
#0.0.2 >> 테스트 (제병)