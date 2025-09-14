package main

import (
	"context"
	"log"
	"net"

	pb "github.com/norovone/bigtech_go_msa_hw/users/api/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: server is used to implement pb.NotesServiceServer.
type server struct {
	pb.UnimplementedUserServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.UserProfile, error) {
	return &pb.UserProfile{
		UserId:    "Test_user_id",
		Nickname:  "Test_nickname",
		Bio:       "Test_bio",
		AvatarUrl: "test_avatar_url",
	}, nil
}

func (s *server) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserProfile, error) {
	return &pb.UserProfile{
		UserId:    "Test_user_id",
		Nickname:  "Test_nickname",
		Bio:       "Test_bio",
		AvatarUrl: "test_avatar_url",
	}, nil
}

func (s *server) GetProfileByID(ctx context.Context, req *pb.GetProfileByIDRequest) (*pb.UserProfile, error) {
	return &pb.UserProfile{
		UserId:    "Test_user_id",
		Nickname:  "Test_nickname",
		Bio:       "Test_bio",
		AvatarUrl: "test_avatar_url",
	}, nil
}

func (s *server) GetProfileByNickname(ctx context.Context, req *pb.GetProfileByNicknameRequest) (*pb.UserProfile, error) {
	return &pb.UserProfile{
		UserId:    "Test_user_id",
		Nickname:  "Test_nickname",
		Bio:       "Test_bio",
		AvatarUrl: "test_avatar_url",
	}, nil
}

func (s *server) SearchByNickname(ctx context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error) {
	return &pb.SearchByNicknameResponse{
		Results: []*pb.UserProfile{
			{
				UserId:    "Test_user_id",
				Nickname:  "Test_nickname",
				Bio:       "Test_bio",
				AvatarUrl: "test_avatar_url",
			},
		},
	}, nil
}

func main() {
	implementation := NewServer()

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
