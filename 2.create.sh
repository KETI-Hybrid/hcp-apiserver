#!/bin/bash
#kubectl create ns openmcp
kubectl create -f deploy
cd /go/src/Hybrid_Cluster/clientset/clusterRegister/v1alpha1
kubectl create -f clusterregister.yaml 
#kubectl create -f deploy/operator.yaml

