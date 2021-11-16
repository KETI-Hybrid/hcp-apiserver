#!/bin/bash
# kubectl delete ns "kube-federation-system" --context gke_keti-container_us-central1-a_cluster-1
# kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:cluster-1 --context gke_keti-container_us-central1-a_cluster-1 
# kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:cluster-1-hcp --context gke_keti-container_us-central1-a_cluster-1
# kubectl delete kubefedclusters cluster-1 -n kube-federation-system --context kube-master


kubectl delete ns "kube-federation-system" --context arn:aws:eks:us-east-2:741566967679:cluster/eks-cluster;
kubectl delete sa eks-cluster-hcp -n kube-federation-system --context arn:aws:eks:us-east-2:741566967679:cluster/eks-cluster;
kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:eks-cluster --context arn:aws:eks:us-east-2:741566967679:cluster/eks-cluster ;
kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:eks-cluster-hcp --context arn:aws:eks:us-east-2:741566967679:cluster/eks-cluster;
kubectl delete kubefedclusters.hcp.k8s.io arn:aws:eks:us-east-2:741566967679:cluster/eks-cluster -n kube-federation-system --context kube-master


# kubectl delete ns "kube-federation-system" --context gke_keti-container_us-central1-a_cluster-1;
# kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:cluster-1 --context gke_keti-container_us-central1-a_cluster-1 ;
# kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:cluster-1-hcp --context gke_keti-container_us-central1-a_cluster-1;
# kubectl delete kubefedclusters cluster-1 -n kube-federation-system --context kube-master

# kubectl delete ns "kube-federation-system" --context aks-master;
# kubectl delete sa aks-master-hcp -n kube-federation-system --context aks-master;
# kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:aks-master --context aks-master ;
# kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:aks-master-hcp --context aks-master;
# kubectl delete kubefedclusters aks-master -n kube-federation-system --context master

# kubectl delete ns "kube-federation-system" --context eks-master;
# kubectl delete sa eks-master-hcp -n kube-federation-system --context eks-master;
# kubectl delete clusterroles.rbac.authorization.k8s.io kubefed-controller-manager:eks-master --context eks-master ;
# kubectl delete clusterrolebindings.rbac.authorization.k8s.io kubefed-controller-manager:eks-master-hcp --context eks-master;
# kubectl delete kubefedclusters eks-master -n kube-federation-system --context master
