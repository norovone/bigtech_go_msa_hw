package main

import (
	"context"
	"log"
	"net"

	pb "github.com/norovone/bigtech_go_msa_hw/chat/api/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: server is used to implement pb.NotesServiceServer.
type server struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateDirectChat(ctx context.Context, req *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {
	return &pb.CreateDirectChatResponse{
		ChatId: "test chat id",
	}, nil
}

func (s *server) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	return &pb.GetChatResponse{
		Chat: &pb.Chat{
			Id:             "test id",
			Type:           pb.ChatType_CHAT_TYPE_CHANNEL,
			Title:          "test title",
			ParticipantIds: []string{"1", "2"},
		},
	}, nil
}

func (s *server) ListUserChats(ctx context.Context, req *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	return &pb.ListUserChatsResponse{
		Chats: []*pb.Chat{
			{
				Id:             "test id",
				Type:           pb.ChatType_CHAT_TYPE_CHANNEL,
				Title:          "test title",
				ParticipantIds: []string{"1", "2"},
			},
			{
				Id:             "test id 2",
				Type:           pb.ChatType_CHAT_TYPE_DIRECT,
				Title:          "test title 2",
				ParticipantIds: []string{"3", "4"},
			},
		},
	}, nil
}

func (s *server) ListChatMembers(ctx context.Context, req *pb.ListChatMembersRequest) (*pb.ListChatMembersResponse, error) {
	return &pb.ListChatMembersResponse{
		UserIds: []string{"test 1", "test 2"},
	}, nil
}

func (s *server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{
		Message: &pb.Message{
			Id:       "test id 1",
			ChatId:   "test chat id 1",
			SenderId: "test sender 1",
			Text:     "test text",
		},
	}, nil
}

func (s *server) ListMessages(ctx context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	return &pb.ListMessagesResponse{
		Messages: []*pb.Message{
			{
				Id:       "test id 1",
				ChatId:   "test chat id 1",
				SenderId: "test sender 1",
				Text:     "test text",
			},
			{
				Id:       "test id 2",
				ChatId:   "test chat id 2",
				SenderId: "test sender 2",
				Text:     "test text",
			},
		},
	}, nil
}

func (s *server) StreamMessages(req *pb.StreamMessagesRequest, stream pb.ChatService_StreamMessagesServer) error {
	return nil
}

func main() {
	implementation := NewServer()

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, implementation)

	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
