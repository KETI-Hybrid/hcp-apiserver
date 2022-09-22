package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/KETI-Hybrid/hcp-pkg/apis/resource/v1alpha1"
	"github.com/KETI-Hybrid/hcp-pkg/util/clusterManager"

	"github.com/gorilla/mux"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vpaclientset "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned"
	"k8s.io/klog"
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

	// RealResource 읽어오기
	var real_resource *appsv1.Deployment
	bytes, _ := json.Marshal(resource.RealResource)
	json.Unmarshal(bytes, &real_resource)

	cm, _ := clusterManager.NewClusterManager()
	// TargetCluster가 지정되지 않은 경우
	if resource.TargetCluster == "" {

		// HCPDeployment 생성하기
		hcp_resource := deploymentToHCPDeployment(real_resource)

		r, err := cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Create(context.TODO(), &hcp_resource, metav1.CreateOptions{})
		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Info("Request scheduling to scheduler : %s \n", r.Name)
		}
	} else {
		// TargetCluster가 지정된 경우
		target_clientset := cm.Cluster_kubeClients[resource.TargetCluster]
		// namespace
		namespace := real_resource.ObjectMeta.Namespace
		if namespace == "" {
			namespace = "default"
		}

		hcp_resource := deploymentToHCPDeployment(real_resource)
		hcp_resource.Spec.SchedulingResult.Targets = append(hcp_resource.Spec.SchedulingResult.Targets, v1alpha1.Target{
			Cluster:  resource.TargetCluster,
			Replicas: real_resource.Spec.Replicas,
		})
		_, err = cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Create(context.TODO(), &hcp_resource, metav1.CreateOptions{})
		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Info("Succeed to create hcpdeployment: %s \n", hcp_resource.Name)
		}
		// Kubernetes Deployment 생성
		r, err := target_clientset.AppsV1().Deployments(namespace).Create(context.TODO(), real_resource, metav1.CreateOptions{})

		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Info("Succeed to create deployment %s \n", r.Name)
		}
	}
}

func DeleteDeploymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]
	name := vars["name"]

	hcpdeployment, err := cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Get(context.TODO(), name, metav1.GetOptions{})

	if !hcpdeployment.Spec.SchedulingNeed && hcpdeployment.Spec.SchedulingComplete {
		targets := hcpdeployment.Spec.SchedulingResult.Targets
		for _, target := range targets {
			// TODO : cluster unregister한 경우
			cm, _ := clusterManager.NewClusterManager()
			clientset := cm.Cluster_kubeClients[target.Cluster]
			err = clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		}
	}

	if err != nil {
		klog.Error(err)
		return
	} else {
		klog.Info("Succeed to delete deployment %s \n", name)

		cm.HCPResource_Client.HcpV1alpha1().HCPHybridAutoScalers("hcp").Get(context.TODO(), name, metav1.GetOptions{})
		deleteHCPHAS(hcpdeployment, name)

		err = cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Info("Succeed to delete hcpdeployment %s \n", name)
		}
	}
}

func CreatePodHandler(w http.ResponseWriter, r *http.Request) {

	var resource Resource

	jsonDataFromHttp, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsonDataFromHttp, &resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// RealResource 읽어오기
	var real_resource *v1.Pod
	bytes, _ := json.Marshal(resource.RealResource)
	json.Unmarshal(bytes, &real_resource)

	// TargetCluster가 지정되지 않은 경우
	if resource.TargetCluster == "undefined" {
		// HCPDeployment 생성하기
		hcp_resource := v1alpha1.HCPPod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "HCPPod",
				APIVersion: "hcp.crd.com/v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: real_resource.Name,
			},
			Spec: v1alpha1.HCPPodSpec{
				RealPodSpec:     real_resource.Spec,
				RealPodMetadata: real_resource.ObjectMeta,

				// SchedulingStatus "Requested"
				SchedulingStatus: "Requested",
				// SchedulingType:   algorithm,
			},
		}

		r, err := cm.HCPResource_Client.HcpV1alpha1().HCPPods("hcp").Create(context.TODO(), &hcp_resource, metav1.CreateOptions{})
		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Infof("request scheduling to scheduler : %s \n", r.Name)
		}
	} else {
		// TargetCluster가 지정된 경우
		cm, _ := clusterManager.NewClusterManager()
		clientset := cm.Cluster_kubeClients[resource.TargetCluster]
		// namespace
		namespace := real_resource.ObjectMeta.Namespace
		if namespace == "" {
			namespace = "default"
		}

		// Kubernetes Deployment 생성
		r, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), real_resource, metav1.CreateOptions{})

		if err != nil {
			klog.Error(err)
			return
		} else {
			klog.Infof("Succeed to create pod %s \n", r.Name)
		}
	}
}

func CreateHCPHASHandler(w http.ResponseWriter, r *http.Request) {
	klog.Info("Called CreateHCPHASHandler")
	var resource Resource

	jsonDataFromHttp, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsonDataFromHttp, &resource)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// RealResource 읽어오기
	var real_resource *v1alpha1.HCPHybridAutoScaler
	bytes, _ := json.Marshal(resource.RealResource)
	json.Unmarshal(bytes, &real_resource)

	cm, _ := clusterManager.NewClusterManager()

	name := real_resource.Spec.ScalingOptions.HpaTemplate.Spec.ScaleTargetRef.Name
	klog.Infof("[1]Get HCPDeployment %s\n", name)
	_, err = cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Error(err)
		w.Write([]byte(err.Error()))
		return
	} else {
		// has 삭제
		klog.Infof("[2]Create HCPHAS %s\n", real_resource.ObjectMeta.Name)
		real_resource.Status.ResourceStatus = "CREATED"
		r, err := cm.HCPResource_Client.HcpV1alpha1().HCPHybridAutoScalers("hcp").Create(context.TODO(), real_resource, metav1.CreateOptions{})
		if err != nil {
			klog.Error(err)
			w.Write([]byte(err.Error()))
			return
		} else {
			str := "Succeed to create hcphas " + r.ObjectMeta.Name
			klog.Info(str)
			w.Write([]byte(str + "\n"))
		}

	}
}

func DeleteHCPHASHandler(w http.ResponseWriter, r *http.Request) {
	klog.Info("Called DeleteHCPHASHandler")
	vars := mux.Vars(r)
	name := vars["name"]
	cm, _ := clusterManager.NewClusterManager()
	var msg string

	klog.Infof("[1]Get HCPDeployment %s", name)
	hcpdeployment, err := cm.HCPResource_Client.HcpV1alpha1().HCPDeployments("hcp").Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	} else {
		msg = deleteHCPHAS(hcpdeployment, name)
		// targets := hcpdeployment.Spec.SchedulingResult.Targets
		// deployment := hcpdeployment.Spec.RealDeploymentMetadata
		// var namespace string
		// if deployment.Namespace == "" {
		// 	namespace = "default"
		// } else {
		// 	namespace = deployment.Namespace
		// }
		// for _, target := range targets {

		// 	target_clientset := cm.Cluster_kubeClients[target.Cluster]
		// 	target_config := cm.Cluster_configs[target.Cluster]
		// 	vpa_clientset, _ := vpaclientset.NewForConfig(target_config)

		// 	// hpa 삭제
		// 	klog.Infof("[2-1]Delete HPA %s in cluster %s\n", name, target.Cluster)
		// 	_, err := target_clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		// 	if !errors.IsNotFound(err) {
		// 		err = target_clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		// 		if err != nil {
		// 			msg += err.Error() + "\n"
		// 		} else {
		// 			str := "Succeed to delete hpa " + name
		// 			klog.Infof("Succeed to delete hpa %s \n", name)
		// 			msg += str + "\n"
		// 		}
		// 	} else {
		// 		klog.Error(err)
		// 		msg += err.Error() + "\n"
		// 	}

		// 	// vpa 삭제
		// 	klog.Infof("[2-1]Delete VPA %s in cluster %s\n", name, target.Cluster)
		// 	_, err = vpa_clientset.AutoscalingV1beta2().VerticalPodAutoscalers(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		// 	if !errors.IsNotFound(err) {
		// 		err = vpa_clientset.AutoscalingV1beta2().VerticalPodAutoscalers(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		// 		if err != nil {
		// 			msg += err.Error() + "\n"
		// 		} else {
		// 			str := "Succeed to delete vpa " + name
		// 			klog.Infof("Succeed to delete vpa %s \n", name)
		// 			msg += str + "\n"
		// 		}
		// 	} else {
		// 		klog.Error(err)
		// 		msg += err.Error() + "\n"
		// 	}

		// 	// has 삭제
		// 	klog.Infof("[2-2]Delete HCPHAS %s", name)
		// 	err = cm.HCPResource_Client.HcpV1alpha1().HCPHybridAutoScalers("hcp").Delete(context.TODO(), name, metav1.DeleteOptions{})
		// 	if err != nil {
		// 		msg += err.Error() + "\n"
		// 	} else {
		// 		str := "Succeed to delete hcphas " + name
		// 		klog.Infof("Succeed to delete hcphas %s \n", name)
		// 		msg += str + "\n"
		// 	}

		w.Write([]byte(msg))
		// }
	}
}

func deleteHCPHAS(hcpdeployment *v1alpha1.HCPDeployment, name string) string {

	var msg string

	targets := hcpdeployment.Spec.SchedulingResult.Targets
	deployment := hcpdeployment.Spec.RealDeploymentMetadata
	var namespace string
	if deployment.Namespace == "" {
		namespace = "default"
	} else {
		namespace = deployment.Namespace
	}
	for _, target := range targets {

		target_clientset := cm.Cluster_kubeClients[target.Cluster]
		target_config := cm.Cluster_configs[target.Cluster]
		vpa_clientset, _ := vpaclientset.NewForConfig(target_config)

		// hpa 삭제
		klog.Infof("[2-1]Delete HPA %s in cluster %s\n", name, target.Cluster)
		_, err := target_clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if !errors.IsNotFound(err) {
			err = target_clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
			if err != nil {
				msg += err.Error() + "\n"
			} else {
				str := "Succeed to delete hpa " + name
				klog.Infof("Succeed to delete hpa %s \n", name)
				msg += str + "\n"
			}
		} else {
			klog.Error(err)
			msg += err.Error() + "\n"
		}

		// vpa 삭제
		klog.Infof("[2-1]Delete VPA %s in cluster %s\n", name, target.Cluster)
		_, err = vpa_clientset.AutoscalingV1beta2().VerticalPodAutoscalers(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if !errors.IsNotFound(err) {
			err = vpa_clientset.AutoscalingV1beta2().VerticalPodAutoscalers(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
			if err != nil {
				msg += err.Error() + "\n"
			} else {
				str := "Succeed to delete vpa " + name
				klog.Infof("Succeed to delete vpa %s \n", name)
				msg += str + "\n"
			}
		} else {
			klog.Error(err)
			msg += err.Error() + "\n"
		}

		// has 삭제
		klog.Infof("[2-2]Delete HCPHAS %s", name)
		err = cm.HCPResource_Client.HcpV1alpha1().HCPHybridAutoScalers("hcp").Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			msg += err.Error() + "\n"
		} else {
			str := "Succeed to delete hcphas " + name
			klog.Infof("Succeed to delete hcphas %s \n", name)
			msg += str + "\n"
		}
	}
	return msg
}

func deploymentToHCPDeployment(real_resource *appsv1.Deployment) v1alpha1.HCPDeployment {
	// HCPDeployment 생성하기
	hcp_resource := v1alpha1.HCPDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HCPDeployment",
			APIVersion: "hcp.crd.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: real_resource.Name,
		},
		Spec: v1alpha1.HCPDeploymentSpec{
			RealDeploymentSpec:     real_resource.Spec,
			RealDeploymentMetadata: real_resource.ObjectMeta,

			// SchedulingStatus "Requested"
			SchedulingNeed:     true,
			SchedulingComplete: false,
			//SchedulingType:   algorithm[0],
		},
	}
	return hcp_resource
}
