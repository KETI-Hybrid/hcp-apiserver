package main

import (
	"flag"
	"hcp-apiserver/pkg/apis/azure"
	"hcp-apiserver/pkg/apis/eks"
	"hcp-apiserver/pkg/apis/google"
	"hcp-apiserver/pkg/apis/naver"
	"hcp-apiserver/pkg/client/kubernetes"
	docazure "hcp-apiserver/pkg/docs/azure"
	doceks "hcp-apiserver/pkg/docs/eks"
	docgoogle "hcp-apiserver/pkg/docs/google"
	docnaver "hcp-apiserver/pkg/docs/naver"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

func main() {
	var runType string

	flag.StringVar(&runType, "t", "real", "")

	kubernetes.InitHCPClient()
	router := httprouter.New()
	if strings.Compare(runType, "real") == 0 {
		azure.InitAKSEndPoint(router)
		eks.InitEKSEndPoint(router)
		google.InitGKEEndPoint(router)
		naver.InitNKSEndPoint(router)
	} else {
		docazure.InitAKSEndPoint(router)
		doceks.InitEKSEndPoint(router)
		docgoogle.InitGKEEndPoint(router)
		docnaver.InitNKSEndPoint(router)
	}

	klog.Fatal(http.ListenAndServe(":30850", router))
}

