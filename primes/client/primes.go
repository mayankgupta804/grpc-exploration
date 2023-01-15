package main

import (
	"context"
	pb "grpc-exploration/primes/proto"
	"io"
	"log"
)

func getPrimes(c pb.PrimesServiceClient) {
	log.Println("getPrimes was invoked")

	req := &pb.PrimesRequest{
		Num: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Num)
	}
}
