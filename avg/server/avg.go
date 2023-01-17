package main

import (
	pb "grpc-exploration/avg/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.AvgService_AverageServer) error {
	log.Println("Average function was invoked")
	res := float64(0)
	numCount := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Num: res / float64(numCount),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += float64(req.Num)
		numCount += 1
	}
}
