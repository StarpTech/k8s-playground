# K8s Playground

Install https://microk8s.io/

Docs: https://microk8s.io/docs/

## 2. Setup aliases

```bash
# Alias
sudo snap alias microk8s.kubectl kubectl
sudo snap alias microk8s.helm helm
# Export image from local registry
docker save k8-simple-golang > k8-simple-golang.tar
# Import image
microk8s.ctr image import k8-simple-golang.tar
# List all images
microk8s.ctr images ls
```

## 3. Enable K8s addons

```bash
microk8s.enable knative dashboard helm registry storage ingress
microk8s.helm init
microk8s.status
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

# Reference tutorials

- https://dev.to/ishankhare07/deploy-a-golang-api-on-kubernetes-with-nginx-ingress-33k6
- https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-resource-requests-and-limits

# K8s CodeSheet

https://kubernetes.io/de/docs/reference/kubectl/cheatsheet/
