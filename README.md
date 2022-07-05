# Test running a service on GKE

Steps:

``` sh
gcloud auth login $GCP_USER
gcloud config set project $PROJECT_ID
gcloud container clusters get-credentials $GKE_CLUSTER_ID
gcloud auth configure-docker

cd ./test-app
nix-build container.nix
podman load < result
podman images # Choose test-app image
podman tag localhost/test-app:$TAG gcr.io/$PROJECT_ID/test-app:tag1
podman push gcr.io/$PROJECT_ID/test-app:tag1

cd ..
kubectl apply -f gke/
```
