package handler

import (
	cobrautil "Hybrid_Cluster/hybridctl/util"
	resourcev1alpha1 "Hybrid_Cluster/pkg/apis/resource/v1alpha1"
	resourcev1alpha1clientset "Hybrid_Cluster/pkg/client/resource/v1alpha1/clientset/versioned"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	policy "Hybrid_Cluster/resource"

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

	algorithm, err := policy.GetAlgorithm()
	if err != nil {
		fmt.Println(err)
		return
	}

	if resource.TargetCluster == "undefined" {

		master_config, err := cobrautil.BuildConfigFromFlags("kube-master", "/root/.kube/config")
		if err != nil {
			fmt.Println(err)
			return
		}
		clienset, err := resourcev1alpha1clientset.NewForConfig(master_config)
		if err != nil {
			fmt.Println(err)
			return
		}
		// analytic Engine
		hcp_resource := resourcev1alpha1.HCPDeployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: real_resource.Name,
			},
			Spec: resourcev1alpha1.HCPDeploymentSpec{
				RealDeploymentSpec:     real_resource.Spec,
				RealDeploymentMetadata: real_resource.ObjectMeta,

				SchedulingStatus: "Requested",
				SchedulingType:   algorithm,
			},
		}

		r, err := clienset.HcpV1alpha1().HCPDeployments("hcp").Create(context.TODO(), &hcp_resource, metav1.CreateOptions{})
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("[1] Request scheduling to scheduler : %s \n", r.Name)
		}
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
			fmt.Printf("success to create deployment %s \n", r.Name)
		}
	}
}

func DeleteDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	target_cluster := vals.Get("cluster")
	namespace := vals.Get("namespace")
	name := vals.Get("name")
	config, err := cobrautil.BuildConfigFromFlags(target_cluster, "/root/.kube/config")
	if err != nil {
		fmt.Println(err)
	}
	clientset, _ := kubernetes.NewForConfig(config)

	// create resource
	err = clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("success to delete deployment %s \n", name)
	}
}
