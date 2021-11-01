package handler

import (
	KubeFedCluster "Hybrid_Cluster/apis/kubefedcluster/v1alpha1"
	clusterRegister "Hybrid_Cluster/clientset/v1alpha1"
	mappingTable "Hybrid_Cluster/hcp-apiserver/pkg/converter"
	util "Hybrid_Cluster/hcp-apiserver/pkg/util"
	cobrautil "Hybrid_Cluster/hybridctl/util"
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Unjoin(info mappingTable.ClusterInfo) {

	master_config, _ := cobrautil.BuildConfigFromFlags("kube-master", "/root/.kube/config")
	clusterRegisterClientSet, err := clusterRegister.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
	}

	clusterRegisters, err := clusterRegisterClientSet.ClusterRegister(info.PlatformName).Get(info.ClusterName, metav1.GetOptions{})

	if info.PlatformName == "gke" {
		projectId := clusterRegisters.Spec.Projectid
		fProjectId := flag.String("projectId", projectId, "specify a project id to examine")
		flag.Parse()
		if *fProjectId == "" {
			log.Fatal("must specific -projectId")
		}

		kubeConfig, err := util.GetK8sClusterConfigs(context.TODO(), projectId)
		if err != nil {
			log.Println(err)
		}

		var join_cluster_client *kubernetes.Clientset
		var join_cluster_config *rest.Config
		for clusterName := range kubeConfig.Clusters {
			gkeClusterName := "gke" + "_" + clusterRegisters.Spec.Projectid + "_" + clusterRegisters.Spec.Region + "_" + info.ClusterName
			if clusterName == gkeClusterName {
				join_cluster_config, err = clientcmd.NewNonInteractiveClientConfig(*kubeConfig, gkeClusterName, &clientcmd.ConfigOverrides{CurrentContext: clusterName}, nil).ClientConfig()
				if err != nil {
					log.Println(err)
				}

				join_cluster_client, err = kubernetes.NewForConfig(join_cluster_config)
				if err != nil {
					log.Println(err)
				} else {
					unjoinCluster(info, join_cluster_client)
				}
			}
		}

	} else if info.PlatformName == "aks" {
		unjoin_cluster_config, _ := cobrautil.BuildConfigFromFlags(info.ClusterName, "/root/.kube/config")
		unjoin_cluster_client := kubernetes.NewForConfigOrDie(unjoin_cluster_config)
		unjoinCluster(info, unjoin_cluster_client)
	} else if info.PlatformName == "eks" {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(clusterRegisters.Spec.Region),
		}))
		eksSvc := eks.New(sess)

		input := &eks.DescribeClusterInput{
			// clusterName eks-master
			Name: aws.String(info.ClusterName),
		}
		result, err := eksSvc.DescribeCluster(input)
		if err != nil {
			fmt.Println(err)
		}

		unjoin_cluster_client, err := util.NewClientset(result.Cluster)
		if err != nil {
			fmt.Println(err)
		} else {
			unjoinCluster(info, unjoin_cluster_client)
		}
	}
}

func unjoinCluster(info mappingTable.ClusterInfo, cluster_client *kubernetes.Clientset) bool {

	master_config, _ := cobrautil.BuildConfigFromFlags("kube-master", "/root/.kube/config")
	apiextensionsClientSet, err := KubeFedCluster.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
	}

	var status = false
	kubefedList, err := apiextensionsClientSet.KubeFedCluster("kube-federation-system").List(metav1.ListOptions{})
	fmt.Printf("kubefedList.Items: %v\n", kubefedList.Items[0].ObjectMeta.Name)

	if err != nil {
		log.Println(err)
	} else {
		for index := range kubefedList.Items {
			if kubefedList.Items[index].ObjectMeta.Name == info.ClusterName {
				status = true
			}
		}
	}

	if status {
		// 1. Delete namespace kube-federation-sysem & service account & secret (in targetcluster)
		err_deletens := cluster_client.CoreV1().Namespaces().Delete(context.TODO(), "kube-federation-system", metav1.DeleteOptions{})
		if err_deletens == nil {
			fmt.Println("[Step 1] DELETE NS Resource in", info.ClusterName)
		} else {
			fmt.Println("Fail to DELETE NS Resource in", info.ClusterName)
			fmt.Println("err_deletens: ", err_deletens)
		}

		// 2. Delete Cluster Role (in targetcluster)
		err_deletecr := cluster_client.RbacV1().ClusterRoles().Delete(context.TODO(), "kubefed-controller-manager:"+info.ClusterName, metav1.DeleteOptions{})
		if err_deletecr == nil {
			fmt.Println("[Step 2] DELETE CR Resource in", info.ClusterName)
		} else {
			fmt.Println("Fail to DELETE CR Resource in", info.ClusterName)
			fmt.Println("err_deletecr: ", err_deletecr)
		}

		// 3. Delete Cluster Role Binding (in targetcluster)
		err_deletecrb := cluster_client.RbacV1().ClusterRoleBindings().Delete(context.TODO(), "kubefed-controller-manager:"+info.ClusterName+"-hcp", metav1.DeleteOptions{})
		if err_deletens == nil {
			fmt.Println("[Step 3] DELETE CRB Resource in", info.ClusterName)
		} else {
			fmt.Println("Fail to DELETE CRB Resource in", info.ClusterName)
			fmt.Println("err_deletecrv: ", err_deletecrb)
		}

		// 4. Delete secret
		saName := info.ClusterName + "-hcp"
		join_cluster_sa, err_sa1 := cluster_client.CoreV1().ServiceAccounts("kube-federation-system").Get(context.TODO(), saName, metav1.GetOptions{})
		if err_sa1 != nil {
			log.Println(err_sa1)
		}
		join_cluster_secret, err_sc := cluster_client.CoreV1().Secrets("kube-federation-system").Get(context.TODO(), join_cluster_sa.Secrets[0].Name, metav1.GetOptions{})
		if err_sc != nil {
			log.Println(err_sc)
		} else {
			err := cluster_client.CoreV1().Secrets("kube-federation-system").Delete(context.TODO(), join_cluster_secret.Name, metav1.DeleteOptions{})
			if err != nil {
				log.Println(err)
			}
			fmt.Println("[Step 4] DELETE Secret Resource [" + join_cluster_secret.Name + "] in master")
		}

		// 5. Delete Kubefedcluster
		_, err_nkfc := apiextensionsClientSet.KubeFedCluster("kube-federation-system").Delete(context.TODO(), info.ClusterName, metav1.DeleteOptions{})
		if err_nkfc != nil {
			log.Println(err_nkfc)
		} else {
			fmt.Println("[Step 5] Create KubefedCluster Resource [" + info.ClusterName + "] in hcp")
		}
	} else {
		return false
	}
	return true
}
