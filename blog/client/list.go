package main

import (
	"context"
	pb "grpc-exploration/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("----listBlog was invoked----")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
