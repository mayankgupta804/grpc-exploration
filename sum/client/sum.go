package main

import (
	"context"
	pb "grpc-exploration/sum/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Sum)
}
