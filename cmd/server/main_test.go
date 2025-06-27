package main

import (
	"context"
	"log"
	"testing"

	pb "docker-explorer/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func GetClient(t *testing.T) (pb.FileServiceClient, func()) {
	server, listener := GetServer()
	go server.Serve(listener)

	conn, err := grpc.Dial(
		listener.Addr().String(),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}

	cleanup := func() {
		conn.Close()
		server.Stop()
		listener.Close()
	}

	return pb.NewFileServiceClient(conn), cleanup
}

func TestWrite(t *testing.T) {
	log.Println("start")
	client, cleanup := GetClient(t)
	log.Println("got client")
	log.Println("defer")

	response, err := client.Write(
		context.Background(),
		&pb.WriteRequest{Content: []byte("helloworld"), Path: "file.txt"},
	)
	log.Println("after write")
	assert.NoError(t, err)
	assert.Equal(t, true, response.Success)
	log.Println("assert")
	cleanup()
}

func TestRead(t *testing.T) {
	client, cleanup := GetClient(t)
	defer cleanup()

	response, err := client.Read(
		context.Background(),
		&pb.ReadRequest{Path: "file.txt"},
	)
	assert.NoError(t, err)
	assert.Equal(t, "helloworld", string(response.Content))
}

func TestDelete(t *testing.T) {
	client, cleanup := GetClient(t)
	defer cleanup()

	_, err := client.Delete(
		context.Background(),
		&pb.DeleteRequest{Path: "file.txt"},
	)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	client, cleanup := GetClient(t)
	defer cleanup()

	response, err := client.List(
		context.Background(),
		&pb.ListRequest{Path: "."},
	)
	assert.NoError(t, err)
	assert.IsType(t, &pb.ListResponse{}, response)
}
