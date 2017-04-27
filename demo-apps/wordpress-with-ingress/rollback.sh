#!/bin/bash

#To roll back
kubectl rollout history deployment wordpress

kubectl rollout undo deployment wordpress

kubectl get replicaset
