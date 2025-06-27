build:
	go mod download
	CGO_ENABLED=0 go build -o docker_explorer ./cmd/server
	docker build -t docker-explorer .

run:
	docker run -p 50051:50051 docker-explorer

gen:
	rm -rf proto/*.go
	protoc --go_out=. --go-grpc_out=. proto/file_service.proto

test:
	go test -v ./...
