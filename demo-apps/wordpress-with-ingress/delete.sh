kubectl delete -f local-volumes.yaml
kubectl delete secret mysql-pass
kubectl delete -f mysql-deployment.yaml
kubectl delete -f wordpress-deployment.yaml

#kubectl delete deployment,service -l app=wordpress
#kubectl delete secret mysql-pass
#kubectl delete pvc -l app=wordpress
#kubectl delete pv local-pv-1 local-pv-2

# Delete ingress
kubectl delete -f nginx-ingress-controller.yaml 
kubectl delete -f myingress.yaml
