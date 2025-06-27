# Docker Explorer

<div>
  <a href="https://github.com/dmkjfs/docker-explorer/stargazers/"><img src="https://img.shields.io/github/stars/dmkjfs/docker-explorer?style=flat-square&label=stars&color=blue"></a>
  <a href="https://github.com/dmkjfs/docker-explorer/forks/"><img src="https://img.shields.io/github/forks/dmkjfs/docker-explorer?style=flat-square&label=forks&color=blue"></a>
  <a href="https://github.com/dmkjfs/docker-explorer/actions/workflows/check.yaml/"><img src="https://img.shields.io/github/actions/workflow/status/dmkjfs/docker-explorer/check.yaml?branch=dev&style=flat-square&label=linter&color=blue"></a>
  <img src="https://img.shields.io/codeclimate/maintainability/dmkjfs/docker-explorer?style=flat-square&label=maintainability&color=blue">
  <img src="https://img.shields.io/codeclimate/coverage/dmkjfs/docker-explorer?style=flat-square&label=coverage&color=blue">
</div>


### ✍️ About

Use docker container as a file storage for your web app. Works via gRPC, easy in use.


### 🛠️ Quick start

Pull this service from DockerHub and run it, opening 50051 port outside of a container
```bash
docker pull
docker run -p 50051:50051 docker-explorer
```

Alternatively, for docker-compose add docker-explorer to services in your docker-compose.yml file
```yaml
services:
    explorer:
        image: docker-explorer:latest
        container_name: explorer

        volumes:
            # mapping volumes for backing
            # up your files locally
            - ./local_storage:./

        networks:
            # using a network to make explorer
            # accessible from other services
            # without opening the port locally
            - some_network 
```
