#!/bin/bash

# kubectl delete ns hcp
kubectl delete -f deploy/operator.yaml
# kubectl delete -f deploy/pv.yaml
kubectl delete -f deploy/pvc.yaml
kubectl delete -f deploy/role_binding.yaml
kubectl delete -f deploy/service_account.yaml
kubectl delete -f deploy/service.yaml