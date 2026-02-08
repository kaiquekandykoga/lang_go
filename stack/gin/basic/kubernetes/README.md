# Kubernetes Deployment for Gin Basic App

## Prerequisites
- Rancher Desktop with Kubernetes enabled
- Docker images built locally

## Build and Deploy

### 1. Build the Docker image
```bash
# From the basic directory
nerdctl build -t lang-go-gin-basic:latest .
# Save image to tarball
nerdctl save -o lang-go-gin-basic.tar d43662a41cb9
# Load tarball
nerdctl load -i lang-go-gin-basic.tar --namespace k8s.io
# Verify image in k8s.io
nerdctl images --namespace k8s.io | grep lang-go-gin-basic
```

### 2. Deploy to Kubernetes
```bash
# Apply the manifests
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml

# Restart deployment
kubectl rollout restart deployment gin-basic
```

### 3. Verify deployment
```bash
# Check pods
kubectl get pods -l app=gin-basic

# Check service
kubectl get svc gin-basic

# Get the NodePort
kubectl get svc gin-basic -o jsonpath='{.spec.ports[0].nodePort}'
```

### 4. Access the application
```bash
# Get the NodePort (output will be a port number like 30000-32767)
export NODE_PORT=$(kubectl get svc gin-basic -o jsonpath='{.spec.ports[0].nodePort}')

# Test the status endpoint
curl http://localhost:$NODE_PORT/status

# Test the time endpoint
curl http://localhost:$NODE_PORT/time
```

## Clean up
```bash
kubectl delete -f kubernetes/deployment.yaml
kubectl delete -f kubernetes/service.yaml
```

## Notes
- The service uses `NodePort` type, suitable for Rancher Desktop local development
- Health checks are configured on the `/status` endpoint
- `imagePullPolicy: Never` ensures Kubernetes uses the local image built with nerdctl
