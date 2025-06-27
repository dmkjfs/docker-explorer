FROM alpine:latest

WORKDIR /

COPY ./docker_explorer /docker_explorer

EXPOSE 50051

ENTRYPOINT ["/docker_explorer"]
