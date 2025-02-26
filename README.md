# DevOps Challenge: Fix &amp; Deploy Go App with Redis

##  Overview:
- Your task is to troubleshoot, fix, and deploy a Go web application that uses Redis for caching. There are some issues in the Dockerfile or the Go code. After fixing these issues, you'll deploy the app to a Kubernetes cluster with Redis.


## Part 1: Fix the Dockerfile and Go Application

- Fixed bug in `CMD ["bin/myapp"]` to `CMD ["app/myapp"]`
- Changed PORT to 8080 to ensure clarity
- Set up a redis container using docker-compose and passed its environment variables
- Optimized image size using multistage builds by using `debian:bookworm-slim`, effectively reducing image size from 1.03GB to 83.2MB (~92% size reduction!!!)
- Could have used smaller images like `alpine` (final size: 16.2MB!!), but this go application depends on `glibc` to work


## Part 2: Deploy to Kubernetes
- Create Kubernetes YAML files to deploy Go (as **stateless** workload) and Redis (as **stateful** workload).
- Use **separate namespaces: "app" & "db"**.
- Redis should have persistent storage & network ID.
- Manage variables using Kubernetes native capabilities.
- One pod per each workload is fine, no need for over provisioning.
- Expose the Go app using nodeport or loadbalancer service.
- Redis should be exposed with the appropriate service type that is suitable for internal communications.


# Proof:

<div style="text-align: center;">
  
  ![docker-compose](images/docker-compose-build.png)

  <p style="font-style: italic;">docker-compose up --build</p>
</div>

<div style="text-align: center;">
  
  ![docker-logs](images/docker-logs.png)
  
  <p style="font-style: italic;">docker compose logs app</p>
</div>

<div style="text-align: center;">
  
  ![docker-ps-and-curl](images/docker-ps-and-curl.png)
  
  <p style="font-style: italic;">docker ps & curl http://localhost:8080</p>
</div>

<div style="text-align: center;">

  ![docker-images](images/docker-images.png)
  
  <p style="font-style: italic;">docker images
  (notice the reduction in image size!!)</p>
</div>


# This challenge will help you:

- Troubleshoot Dockerfiles and Go applications.
- Build and run Docker containers locally.
- Deploy stateful & stateless workload to Kubernetes.
- Manage configurations and variables in Kubernetes environments.
- Use storage persistence in Kubernetes environments.
- Manage traffic across namespaces & persistent network IDs.
- Optimize Docker images size. 


