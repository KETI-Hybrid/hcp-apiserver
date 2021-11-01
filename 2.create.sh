#!/bin/bash
# kubectl create ns hcp
kubectl create -f deploy/operator.yaml
# kubectl create -f deploy/pv.yaml
# kubectl create -f deploy/pvc.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/service.yaml
# kubectl create -f deploy/operator.yaml
