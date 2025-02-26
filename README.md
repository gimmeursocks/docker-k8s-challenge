# DevOps Challenge: Fix &amp; Deploy Go App with Redis

![Deploy Status](https://github.com/gimmeursocks/docker-k8s-challenge/actions/workflows/ci-cd.yaml/badge.svg)

## Overview

- My task is to troubleshoot, fix, and deploy a Go web application that uses Redis for caching. There are some issues in the Dockerfile or the Go code. After fixing these issues, the app is to be deployed to a Kubernetes cluster with Redis.

## Part 1: Fix the Dockerfile and Go Application

- Fixed bug in `CMD ["bin/myapp"]` to `CMD ["app/myapp"]`
- Changed PORT to 8080 to ensure clarity
- Set up a Redis container using docker-compose and passed its environment variables
- Optimized image size using multistage builds by using `debian:bookworm-slim`, effectively reducing image size from 1.03GB to 83.2MB (~92% size reduction!!!)
- Could have used smaller images like `alpine` (final size: 16.2MB!!), but this Go application depends on `glibc` to work

### Proof

<div style="text-align: center;">

  ![docker-compose](images/docker-compose-build.png)

  <p>docker-compose up --build</p>

  ![docker-logs](images/docker-logs.png)
  
  <p>docker compose logs app</p>
  
  ![docker-ps-and-curl](images/docker-ps-and-curl.png)
  
  <p>docker ps & curl http://localhost:8080</p>

  ![docker-images](images/docker-images.png)
  
  <p>docker images
  (notice the reduction in image size!!)</p>
</div>

## Part 2: Deploy to Kubernetes

- Created Kubernetes YAML files and deployed Go (as **stateless** workload) and Redis (as **stateful** workload)
- Used separate namespaces: `app` & `db`
- Set up PVC for Redis persistent storage
- Set up clusterIP as none to make network ID stable
- Used configmap to manage variables
- Utilized one pod per each workload
- Uploaded local image to docker hub publicly `gimmeursocks/docker-k8s-challenge_app`
- Used nodeport due to local deployment using minikube
- Exposed Redis internally with stateful state with static hostname

### Proof

<div style="text-align: center;">
  
  ![k8s-namespaces](images/k8s-namespaces.png)

  <p>kubectl apply -f namespaces.yaml</p>

  ![k8s-stateless](images/k8s-stateless.png)

  <p>kubectl apply -f go-deployment.yaml<br/>kubectl apply -f go-service-nodeport.yaml</p>

  ![k8s-stateful](images/k8s-stateful.png)

  <p>kubectl apply -f redis-pvc.yaml<br/>kubectl apply -f redis-statefulset.yaml<br/>kubectl apply -f redis-service.yaml</p>

  ![k8s-verify-namespaces_app](images/k8s-verify-app.png)
  ![k8s-verify-namespaces_db](images/k8s-verify-db.png)

  <p>kubectl get all -n app<br/>kubectl get all -n db</p>

  ![docker-upload-image](images/docker-upload-image.png)

  <p>docker tag & docker push</p>

  ![docker-upload-image](images/docker-upload-image.png)

  <p>docker tag & docker push</p>

  ![k8s-curl](images/k8s-curl.png)

  <p>curl result</p>

  ![k8s-stateful-pvc](images/k8s-stateful-pvc.png)

  <p>Redis persistent storage</p>

  ![k8s-redis-pod-describe-1](images/k8s-redis-pod-1.png)
  ![k8s-redis-pod-describe-2](images/k8s-redis-pod-2.png)

  <p>kubctl describe pod redis-0 -n db</p>

  ![k8s-configmap](images/k8s-configmap.png)

  <p>Managing variables</p>

  ![k8s-one-per-one](images/k8s-one-per-one.png)

  <p>One pod per each workload</p>

  ![k8s-nodeport](images/k8s-svc-db.png)

  <p>Port exposure using nodeport</p>

  ![k8s-redis-internal](images/k8s-redis-stable-hostname.png)

  <p>Internal communication between Redis</p>

  ![k8s-conclusion](images/k8s-pods-and-services.png)

  <p>Pods and Services</p>
</div>

## Bonus

- Created Github Action workflow to automatically build and test the deployment of k8s
- Auto builds using docker compose and then pushes to docker hub
- It tests the deployment of k8s by installing and running minikube in the workflow

<div style="text-align: center;">
  
  ![k8s-namespaces](images/github-actions-testing.png)

  <p>Testing service URLs in the new deployment</p>

</div>