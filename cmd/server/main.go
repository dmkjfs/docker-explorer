package main

import (
	"log"
	"net"

	"docker-explorer/api"
	pb "docker-explorer/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GetServer() (*grpc.Server, net.Listener) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	fileService := api.GetFileServiceServer()
	pb.RegisterFileServiceServer(server, fileService)
	reflection.Register(server)
	return server, listener
}

func main() {
	server, listener := GetServer()

	log.Println("gRPC server listening on :50051")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
