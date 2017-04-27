#!/bin/bash

kubectl get replicaset

kubectl apply -f wordpress-deployment-update.yaml

kubectl get replicaset

