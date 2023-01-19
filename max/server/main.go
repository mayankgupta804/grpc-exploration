package main

import (
	pb "grpc-exploration/max/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.MaxServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on: %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterMaxServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
