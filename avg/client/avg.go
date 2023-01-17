package main

import (
	"context"
	pb "grpc-exploration/avg/proto"
	"log"
	"time"
)

func getAvg(c pb.AvgServiceClient) {
	log.Println("getAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Num)
}
