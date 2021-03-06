# K8s Playground

Install https://microk8s.io/

Docs: https://microk8s.io/docs/

## 2. Setup aliases

```bash
# Alias
sudo snap alias microk8s.kubectl kubectl
sudo snap alias microk8s.helm helm
```

## 3. Enable K8s addons

```bash
microk8s.enable knative dashboard helm registry storage ingress
microk8s.helm init
microk8s.status
```

## Build docker image

```bash
docker build -t k8-simple-golang:1.0 .
# test it
docker run -p 8000:8000 k8-simple-golang:1.0
```

## Import locally build image into K8s registry

```bash
# Export image from local registry
docker save k8-simple-golang > k8-simple-golang.tar
# Import image
microk8s.ctr image import k8-simple-golang.tar
# List all images
microk8s.ctr images ls
```

## Apply configs

```bash
kubectl apply -f name.yaml
```

## K8s CodeSheet

https://kubernetes.io/de/docs/reference/kubectl/cheatsheet/

# Experiments

- Run nginx and route to two different HTTP services: One service is replicated with a factor of 3. https://dev.to/ishankhare07/deploy-a-golang-api-on-kubernetes-with-nginx-ingress-33k6. It will touch following resources and tools:
  - Deployment
  - Service
  - Ingress
  - Helm

## Links

- https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-resource-requests-and-limits
