package main

import (
	"context"
	pb "grpc-exploration/max/proto"
	"io"
	"log"
	"time"
)

func getMax(c pb.MaxServiceClient) {
	log.Println("Max was invoked")

	reqs := []*pb.MaxRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 20},
	}

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Num)
		}
		close(waitc)
	}()
	<-waitc
}
