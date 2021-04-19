This POC made for testing gRPC load balancing with linkerd service mesh and Ambassador as ingress controller
on kubernetes.
Tested on MacOS Catalina 10.15.7 with Docker Desktop kubernetes service


The workflow
------------

![workflow](https://github.com/mpetke/poc-grpc-aes-linkerd/blob/master/diagram.png?raw=true)

When we open the http://localhost/client site, the gprc-client service serves us a JS, which makes a gRPC-web request to our grpc-echo service.
The grpc-echo service will answer with the same message, and add it's hostname.
(This case we can follow the load balancing method)

The grpc-requester service makes a constant load on the grpc-echo service, which procedure imitates an interprocess communication.

We can check the communication's state on Linkerd. The simplest way is to using the Linkerd cli, which will create a port-forwarding for us

```
linkerd dashboard
```

Install Ambassador
------------------

The easiest way to install is using Helm 3

```
# Add the Repo:
helm repo add datawire https://www.getambassador.io

# Install:
kubectl create namespace ambassador && \
helm install ambassador --namespace ambassador datawire/ambassador && \
kubectl -n ambassador wait --for condition=available --timeout=90s deploy -lproduct=aes

# Apply services
kubectl apply -f ./k8s/ambassador/ --namespace=ambassador
```

Install linkerd
---------------

On Mac you can install the cli tool from brew

```
brew install linkerd
```

Or simply use curl

```
curl -sL run.linkerd.io/install | sh
```

Apply on kubernetes

```
# We can do preflight check
linkerd check --pre
# Control plane install (latest)
linkerd install | kubectl apply -f -
# Check the state
linkerd check
# viz extension
linkerd viz install | kubectl apply -f -
```

Add microservices
-----------------

Create namespace and apply services

```
kubectl create namespace microservices
kubectl apply -f ./k8s/microservices/ --namespace=microservices
```

Build containers
----------------

All of the files/containers are available from my dockerhub, but you can build them yourself if you want

To generate the grpc-web client's protobuf message class and client stub js files

```
cd client-grpc
docker run \
    -v "$PWD:/protofile" \
    -e "protofile=src/echo.proto" \
    -e "output=/protofile/" \
    juanjodiaz/grpc-web-generator

npm install
npx webpack client.js
```
