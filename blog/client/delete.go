package main

import (
	"context"
	pb "grpc-exploration/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("----deleteBlog was invoked----")

	blogId := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), blogId)

	if err != nil {
		log.Fatalf("Error happended while deleting: %v\n", err)
	}

	log.Println("Blog was deleted!")
}
