package handler

import (
	cobrautil "Hybrid_Cluster/hybridctl/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Resource struct {
	TargetCluster string
	RealResource  interface{}
}

func CreateDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	jsonDataFromHttp, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsonDataFromHttp, &resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// get real_resource
	var real_resource *appsv1.Deployment
	bytes, _ := json.Marshal(resource.RealResource)
	json.Unmarshal(bytes, &real_resource)

	if resource.TargetCluster == "undefined" {
		// analytic Engine
		fmt.Println("scoring target_cluster")
	} else {
		config, err := cobrautil.BuildConfigFromFlags(resource.TargetCluster, "/root/.kube/config")
		if err != nil {
			fmt.Println(err)
		}
		clientset, _ := kubernetes.NewForConfig(config)

		// namespace
		namespace := real_resource.ObjectMeta.Namespace
		if namespace == "" {
			namespace = "default"
		}

		// create resource
		r, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), real_resource, metav1.CreateOptions{})

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("success to create hcp_resource %s \n", r.Name)
		}
	}
}
