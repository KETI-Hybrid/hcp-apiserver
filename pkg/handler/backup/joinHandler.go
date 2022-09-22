package backup

/*
import (
	cobrautil "Hybrid_Cloud/hybridctl/util"

	// KubeFedCluster "Hybrid_Cloud/pkg/apis/kubefedcluster/v1alpha1"
	// KubeFedCluster "Hybrid_Cloud/pkg/apis/kubefedcluster/v1alpha1"

	"Hybrid_Cloud/util/clusterManager"
	"context"
	"fmt"
	"log"
	"time"

	"k8s.io/client-go/kubernetes"
	kubefed "sigs.k8s.io/kubefed/pkg/kubefedctl"

	// clusterRegister "Hybrid_Cloud/pkg/client/clusterregister/v1alpha1/clientset/versioned/typed/clusterregister/v1alpha1"
	// "context"
	// "flag"
	// "log"
	// "os/exec"
	// "time"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/eks"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/kubefed/pkg/kubefedctl/options"

	// cobrautil "Hybrid_Cloud/hybridctl/util"
	// KubeFedCluster "Hybrid_Cloud/pkg/apis/kubefedcluster/v1alpha1"
	// "fmt"
	// mappingTable "github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/converter"
	// util "github.com/KETI-Hybrid/hcp-apiserver-v1/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// "k8s.io/client-go/tools/clientcmd"
)


func Join(info mappingTable.ClusterInfo) bool {

	master_config, _ := cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
	clusterRegisterClientSet, err := clusterRegister.NewForConfig(master_config)
	if err != nil {
		log.Println(err)
	}

	clusterRegisters, err := clusterRegisterClientSet.ClusterRegisters(info.PlatformName).Get(context.TODO(), info.ClusterName, metav1.GetOptions{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("---joinHandler start---")
	fmt.Println("--> join process start")

	if info.PlatformName == "gke" {
		projectId := clusterRegisters.Spec.ProjectId
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
			gkeClusterName := "gke" + "_" + clusterRegisters.Spec.ProjectId + "_" + clusterRegisters.Spec.Region + "_" + info.ClusterName
			if clusterName == gkeClusterName {
				join_cluster_config, err = clientcmd.NewNonInteractiveClientConfig(*kubeConfig, gkeClusterName, &clientcmd.ConfigOverrides{CurrentContext: clusterName}, nil).ClientConfig()
				if err != nil {
					log.Println(err)
				}

				join_cluster_client, err = kubernetes.NewForConfig(join_cluster_config)
				if err != nil {
					log.Println(err)
				}
			}
		}

		if JoinCluster(info, join_cluster_client, join_cluster_config.Host) {
			clusterRegisters.Status.Join = true
		}

	} else if info.PlatformName == "aks" {

		cmd := exec.Command("az", "aks", "get-credentials", "--resource-group", clusterRegisters.Spec.Resourcegroup, "--name", clusterRegisters.Spec.Name)
		_, err := cmd.Output()
		if err != nil {
			log.Println(err)
		}

		join_cluster_config, _ := cobrautil.BuildConfigFromFlags(info.ClusterName, "/root/.kube/config")
		join_cluster_client := kubernetes.NewForConfigOrDie(join_cluster_config)
		if JoinCluster(info, join_cluster_client, join_cluster_config.Host) {
			clusterRegisters.Status.Join = true
		}

	} else if info.PlatformName == "eks" {

		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(clusterRegisters.Spec.Region),
		}))
		eksSvc := eks.New(sess)

		input := &eks.DescribeClusterInput{
			Name: aws.String(info.ClusterName),
		}
		result, err := eksSvc.DescribeCluster(input)
		if err != nil {
			fmt.Println(err)
		}

		join_cluster_client, err := util.NewClientset(result.Cluster)
		if err != nil {
			fmt.Println(err)
		}
		if JoinCluster(info, join_cluster_client, *result.Cluster.Endpoint) {
			clusterRegisters.Status.Join = true
		}
	}

	fmt.Println("--> join Done!")
	fmt.Println("---joinHandler end---")
	return true
}


func JoinCluster(info ClusterInfo, join_cluster_client *kubernetes.Clientset, APIEndPoint string) bool {

	master_config, _ := cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
	master_client := kubernetes.NewForConfigOrDie(master_config)

	// 1. CREATE namespace "kube-federation-system"
	Namespace := corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-federation-system",
		},
	}

	ns, err_ns := join_cluster_client.CoreV1().Namespaces().Create(context.TODO(), &Namespace, metav1.CreateOptions{})

	if err_ns != nil {
		log.Println(err_ns)
		return false
	} else {
		fmt.Println("< Step 1 > Create Namespace Resource [" + ns.Name + "] in " + info.ClusterName)
	}

	// 2. CREATE service account
	ServiceAccount := corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      info.ClusterName + "-hcp",
			Namespace: "kube-federation-system",
		},
	}

	sa, err_sa := join_cluster_client.CoreV1().ServiceAccounts("kube-federation-system").Create(context.TODO(), &ServiceAccount, metav1.CreateOptions{})

	if err_sa != nil {
		log.Println(err_sa)
		return false
	} else {
		fmt.Println("< Step 2 > Create Namespace Resource [" + sa.Name + "] in " + info.ClusterName)
	}

	// 3. CREATE cluster role
	ClusterRole := rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "kubefed-controller-manager:" + info.ClusterName,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{rbacv1.APIGroupAll},
				Verbs:     []string{rbacv1.VerbAll},
				Resources: []string{rbacv1.ResourceAll},
			},
			{
				NonResourceURLs: []string{rbacv1.NonResourceAll},
				Verbs:           []string{"get"},
			},
		},
	}

	cr, err_cr := join_cluster_client.RbacV1().ClusterRoles().Create(context.TODO(), &ClusterRole, metav1.CreateOptions{})

	if err_cr != nil {
		log.Println(err_cr)
		return false
	} else {
		fmt.Println("< Step 3 > Create ClusterRole Resource [" + cr.Name + "] in " + info.ClusterName)
	}

	// 4. CREATE Cluster Role Binding
	ClusterRoleBinding := rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "kubefed-controller-manager:" + ServiceAccount.Name,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     ClusterRole.Name,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      ServiceAccount.Name,
				Namespace: ServiceAccount.Namespace,
			},
		},
	}
	_ = ClusterRoleBinding

	crb, err_crb := join_cluster_client.RbacV1().ClusterRoleBindings().Create(context.TODO(), &ClusterRoleBinding, metav1.CreateOptions{})

	if err_crb != nil {
		log.Println(err_crb)
		return false
	} else {
		fmt.Println("< Step 4 > Create ClusterRoleBinding Resource [" + crb.Name + "] in " + info.ClusterName)
	}

	time.Sleep(1 * time.Second)

	// 4. GET & CREATE secret (in hcp)
	join_cluster_sa, err_sa1 := join_cluster_client.CoreV1().ServiceAccounts("kube-federation-system").Get(context.TODO(), sa.Name, metav1.GetOptions{})
	if err_sa1 != nil {
		log.Println(err_sa1)
	}
	join_cluster_secret, err_sc := join_cluster_client.CoreV1().Secrets("kube-federation-system").Get(context.TODO(), join_cluster_sa.Secrets[0].Name, metav1.GetOptions{})
	if err_sc != nil {
		log.Println(err_sc)
		return false
	} else {
		fmt.Println("< Step 5-1 > Get Secret Resource [" + join_cluster_secret.Name + "] From " + info.ClusterName)
	}

	Secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: info.ClusterName + "-",
			Namespace:    "kube-federation-system",
		},
		Data: map[string][]byte{
			"token": join_cluster_secret.Data["token"],
		},
	}
	cluster_secret, err_secret := master_client.CoreV1().Secrets("kube-federation-system").Create(context.TODO(), Secret, metav1.CreateOptions{})

	if err_secret != nil {
		log.Println(err_secret)
		return false
	} else {
		fmt.Println("< Step 5-2 > Create Secret Resource [" + cluster_secret.Name + "] in " + "master")
	}

	cm, err := clusterManager.NewClusterManager()
	if err != nil {
		return false
	}
	scope, err := options.GetScopeFromKubeFedConfig(master_config, ns.Name)
	if err != nil {
		return false
	}
	newkubefedcluster, err := kubefed.JoinCluster(cm.Host_config, cm.Cluster_configs[info.ClusterName], ns.Name, master_config.Host, info.ClusterName, cluster_secret.Name, scope, false, false)
	if err != nil {
		log.Println(err)
		return false
	} else {
		fmt.Println("< Step 6 > Create KubefedCluster Resource [" + newkubefedcluster.Name + "] in hcp")
	}

		kubefedcluster := &fedv1b1.KubeFedCluster{
			TypeMeta: metav1.TypeMeta{
				Kind:       "kubefedcluster",
				APIVersion: "core.kubefed.io/v1beta1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      info.ClusterName,
				Namespace: "kube-federation-system",
			},
			Spec: fedv1b1.KubeFedClusterSpec{
				APIEndpoint: APIEndPoint,
				CABundle:    join_cluster_secret.Data["ca.crt"],
				SecretRef: fedv1b1.LocalSecretReference{
					Name: cluster_secret.Name,
				},
				// DisabledTLSValidations: disabledTLSValidations,
			},
		}

		apiextensionsClientSet, err := KubeFedCluster.NewForConfig(master_config)
		if err != nil {
			log.Println(err)
			return false
		}
		newkubefedcluster, err_nkfc := apiextensionsClientSet.KubeFedClusters("kube-federation-system").Create(context.TODO(), (*v1alpha1.KubeFedCluster)(kubefedcluster), metav1.CreateOptions{})

		if err_nkfc != nil {
			log.Println(err_nkfc)
			return false
		} else {
			fmt.Println("< Step 6 > Create KubefedCluster Resource [" + newkubefedcluster.Name + "] in hcp")
		}

	return true
}
*/
