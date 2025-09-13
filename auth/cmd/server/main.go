package main

import pb "github.com/norovone/bigtech_go_msa_hw/auth"

// TODO: server is used to implement pb.NotesServiceServer.
type server struct {
	// UnimplementedNotesServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedNotesServiceServer
}
