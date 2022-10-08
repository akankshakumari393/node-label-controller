# node-label-controller

This controller watches the Kubernetes nodes and attaches a label to the Kubernetes
Node object when the node uses a specific operating system based on a command line argument
that should decide which operating system is targeted to be labeled. 

## On Local Environment
Setup a kind cluster

### build the binary

```
go build
```

### Execute the binary

```
./node-label-controller "linux"
```

## On Kind Cluster

### Build the controller image and push to docker registry 
```
docker build -t akankshakumari393/node-label-controller:0.0.1 .
docker push akankshakumari393/node-label-controller:0.0.1
```

### create a namespace in which controller would run

```
kubectl create namespace controller
```

### Create ClusterRole to give permission for nodes 

```
kubectl create clusterrole cluster-role --verb=list,watch,update --resource=nodes
# kubectl create -f k8s-resources/clusterrole.yaml
```

### Create ClusterRoleBindings give access to `default` Service account in `controller` namespace 

```
kubectl create clusterrolebinding cluster-role-binding --clusterrole=cluster-role --user=system:serviceaccount:controller:default
# kubectl create -f k8s-resources/clusterrolebinding.yaml
```

### Create node label os controller as deployment
```
kubectl create -f k8s-resources/deployment.yaml -n controller
```

This would update the Node resource and add label `k8c.io/uses-linux=true` in above example if the OS matches the arg provided in the deployment, in our case `linux`