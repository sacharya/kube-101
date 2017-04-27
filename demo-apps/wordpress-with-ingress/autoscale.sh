#!/bin/bash

kubectl autoscale deployment wordpress --cpu-percent=50 --min=1 --max=3

kubectl get hpa

#kubectl delete hpa wordpress
