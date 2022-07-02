# stunning-memory

A repository where you can find sample rest-servers in the following languages :

1. Java (spring-boot)     -> `git checkout java`
2. Golang (gorilla mux)   -> `git checkout golang`
3. NodeJS (express)       -> `git checkout nodejs`
4. Python (flask)         -> `git checkout python`
5. Ruby (Sinatra)         -> `git checkout ruby`
6. C++                    -> `git checkout c++`
7. ReactJS                -> `git checkout reactjs`

To practice Docker, Kubernetes, Helm, checkout into the respective branch.</br>
The branch-name is exact name of your language in lowercase.
For example, if I select `ReactJS`, I will type:

```bash
# git checkout <LANGUAGE_IN_LOWERCASE>

git checkout reactjs
```

Now, delete the following files:

1. docker-compose.yml
2. Dockerfile
3. `kubernetes` directory

Once, you are done now, here are your tasks:

- [ ] Dockerize your application.
  - [ ] Write Dockerfile for your application
  - [ ] Create the Docker image.
  - [ ] Use docker-image to create a container for your application.
  - [ ] Hit the endpoints within application to test your application.
  - [ ] Tag your image properly.
  - [ ] Push your Docker image to your Docker Hub repository.
- [ ] Use Docker for development.
  - [ ] Create a docker-compose file for your application.
- [ ] Orchestrate your application manually.
  - [ ] Write a pod manifest for your application.
    - [ ] Create the pod using your manifest.
    - [ ] Use `kubectl port-forward` to test your application.
  - [ ] Write a deployment manifest for your application.
    - [ ] Create the deployment using your manifest.
    - [ ] Use `kubectl port-forward` to test your application.
  - [ ] Write a service manifest for your application - `ClusterIP`.
    - [ ] Create the service using your manifest.
    - [ ] Use `kubectl port-forward` to test your application.
  - [ ] Write a service manifest for your application - `NodeIP`.
    - [ ] Create the service using your manifest.
    - [ ] Test your application.
- [ ] Helmify your application.
  - [ ] Create Helm chart for your application.
  - [ ] Generate Kubernetes manifests for your application using Helm charts.
  - [ ] Orchestrate your application using helm-generate manifests.



