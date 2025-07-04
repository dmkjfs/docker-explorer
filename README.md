# Docker Explorer

<div>
  <a href="https://github.com/dmkjfs/docker-explorer/stargazers/"><img src="https://img.shields.io/github/stars/dmkjfs/docker-explorer?style=flat-square&label=stars&color=blue"></a>
  <a href="https://github.com/dmkjfs/docker-explorer/forks/"><img src="https://img.shields.io/github/forks/dmkjfs/docker-explorer?style=flat-square&label=forks&color=blue"></a>
  <a href="https://github.com/dmkjfs/docker-explorer/actions/workflows/test.yml/"><img src="https://img.shields.io/github/actions/workflow/status/dmkjfs/docker-explorer/test.yml?branch=main&style=flat-square&label=tests&color=blue"></a>
</div>


### ✍️ About

Use docker container as a file storage for your web app. Works via gRPC, easy in use.


### 🛠️ Quick start

Pull this service from DockerHub and run it, opening 50051 port outside of a container
```bash
docker pull dmkjfs/docker-explorer:latest
docker run -p 50051:50051 docker-explorer
```

Alternatively, for docker-compose add docker-explorer to services in your docker-compose.yml file
```yaml
services:
    explorer:
        image: dmkjfs/docker-explorer:latest
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

🗒️ Specification

See full specification for unary server requests&responses in [file_service.proto](https://github.com/dmkjfs/docker-explorer/blob/main/proto/file_service.proto)
