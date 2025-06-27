package api

import (
	"context"
	"log"

	"docker-explorer/internal/service"
	pb "docker-explorer/proto"
)

type FileServiceServer struct {
	pb.UnimplementedFileServiceServer
	FileService *service.FileService
}

func GetFileServiceServer() *FileServiceServer {
	return &FileServiceServer{
		FileService: &service.FileService{},
	}
}

func (s *FileServiceServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	log.Println("Read request")
	path := req.GetPath()
	data, err := s.FileService.Read(path)
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{Content: data}, nil
}

func (s *FileServiceServer) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	log.Println("Write request")
	path := req.GetPath()
	err := s.FileService.Write(path, req.GetContent())
	if err != nil {
		return nil, err
	}
	return &pb.WriteResponse{Success: true}, nil
}

func (s *FileServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Println("Delete request")
	path := req.GetPath()
	err := s.FileService.Delete(path)
	if err != nil {
		return &pb.DeleteResponse{Success: false}, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (s *FileServiceServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	log.Println("List request")
	items, err := s.FileService.List(req.GetPath())
	if err != nil {
		return nil, err
	}
	return &pb.ListResponse{Items: items}, nil
}
