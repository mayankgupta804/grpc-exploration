package main

import (
	pb "grpc-exploration/max/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.MaxService_MaxServer) error {
	log.Println("Max was invoked")

	var currentMax int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		num := req.GetNum()
		if num > currentMax {
			currentMax = num
			if err = stream.Send(&pb.MaxResponse{
				Num: currentMax,
			}); err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
