package main

import (
	"context"
	"log"

	pb "github.com/norovone/bigtech_go_msa_hw/auth/api/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient("localhost:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewAuthServiceClient(conn)

	{
		resp, err := cli.Register(context.Background(), &pb.RegisterRequest{})

		if err != nil {
			log.Fatalf("registration error: %v", err)
		}

		log.Println(resp)
	}

	{
		md := metadata.New(map[string]string{"x-header": "xxxxx"})
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		resp, err := cli.Login(ctx, &pb.LoginRequest{})

		if err != nil {
			log.Fatalf("login error: %v", err)
		}

		log.Println(resp)
	}

	{
		resp, err := cli.Refresh(context.Background(), &pb.RefreshRequest{})

		if err != nil {
			log.Fatalf("refresh token error: %v", err)
		}

		log.Println(resp)
	}
}
