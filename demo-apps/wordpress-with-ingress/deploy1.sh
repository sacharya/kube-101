# Create local volumes for MySQL
kubectl create -f local-volumes.yaml

# Create secret for WordPress
kubectl create secret generic mysql-pass --from-file=password.txt

# Create MySQL deployment
kubectl create -f ./mysql-deployment.yaml

# Create WordPress deployment
kubectl create -f ./wordpress-deployment.yaml

#Ingress
kubectl create -f nginx-ingress-controller.yaml

kubectl create -f myingress1.yaml

kubectl get pods -o wide

kubectl get ing

kubectl get svc

#curl -I -H "Host: wordpress.furiouscat.com" nginx-ingress-ip
