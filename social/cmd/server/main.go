package main

import (
	"context"
	"log"
	"net"

	pb "github.com/norovone/bigtech_go_msa_hw/social/api/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO: server is used to implement pb.NotesServiceServer.
type server struct {
	pb.UnimplementedSocialServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.FriendRequest, error) {
	return &pb.FriendRequest{
		RequestId: "test_request_id",
		Status:    pb.FriendRequest_PENDING,
	}, nil
}

func (s *server) ListRequests(ctx context.Context, req *pb.ListRequestsRequest) (*pb.ListRequestsResponse, error) {
	return &pb.ListRequestsResponse{
		Requests: []*pb.FriendRequest{
			{
				RequestId: "test_request_id",
				Status:    pb.FriendRequest_PENDING,
			},
		},
	}, nil
}

func (s *server) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.FriendRequest, error) {
	return &pb.FriendRequest{
		RequestId: "test_request_id",
		Status:    pb.FriendRequest_ACCEPTED,
	}, nil
}

func (s *server) DeclineFriendRequest(ctx context.Context, req *pb.DeclineFriendRequestRequest) (*pb.FriendRequest, error) {
	return &pb.FriendRequest{
		RequestId: "test_request_id",
		Status:    pb.FriendRequest_DECLINED,
	}, nil
}

func (s *server) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *server) ListFriends(ctx context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	return &pb.ListFriendsResponse{
		FriendUserIds: []string{"test_request_id"},
	}, nil
}

func main() {
	implementation := NewServer()

	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterSocialServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
