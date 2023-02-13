package main

import (
	"context"
	pb "grpc-exploration/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("----updateBlog was invoked----")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Mayank",
		Title:    "A new title",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happended while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
