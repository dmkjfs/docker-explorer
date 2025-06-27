package api

import (
	"context"
	"os"
	"testing"

	"docker-explorer/proto"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	server := GetFileServiceServer()
	request := &proto.WriteRequest{
		Path:    "test.txt",
		Content: []byte("helloworld"),
	}
	response, err := server.Write(context.Background(), request)
	assert.NoError(t, err)
	assert.True(t, response.GetSuccess())
}

func TestRead(t *testing.T) {
	server := GetFileServiceServer()
	request := &proto.ReadRequest{
		Path: "test.txt",
	}
	response, err := server.Read(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, string(response.GetContent()), "helloworld")

	request = &proto.ReadRequest{
		Path: "nonexistent.txt",
	}
	_, err = server.Read(context.Background(), request)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	server := GetFileServiceServer()
	request := &proto.DeleteRequest{
		Path: "test.txt",
	}
	response, err := server.Delete(context.Background(), request)
	assert.NoError(t, err)
	assert.True(t, response.GetSuccess())
}

func TestList(t *testing.T) {
	server := GetFileServiceServer()
	expectedFiles := []string{"file1.txt", "file2.txt", "file3.txt"}

	for _, fname := range expectedFiles {
		err := os.WriteFile(fname, []byte("helloworld"), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", fname, err)
		}
	}

	req := &proto.ListRequest{
		Path: ".",
	}
	resp, err := server.List(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, len(resp.GetItems())-2, len(expectedFiles))
}
