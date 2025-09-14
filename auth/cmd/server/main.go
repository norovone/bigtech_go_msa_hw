package main

import (
	"context"
	"log"
	"net"

	pb "github.com/norovone/bigtech_go_msa_hw/auth/api/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: server is used to implement pb.NotesServiceServer.
type server struct {
	pb.UnimplementedAuthServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{UserId: "test_user_id_111"}, nil
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		AccessToken:  "access_token_test",
		RefreshToken: "refresh_token_test",
		UserId:       "test_user_id_111",
	}, nil
}

func (s *server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return &pb.RefreshResponse{
		AccessToken:  "access_token_test",
		RefreshToken: "refresh_token_test",
		UserId:       "test_user_id_111",
	}, nil
}

func main() {
	implementation := NewServer()

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
