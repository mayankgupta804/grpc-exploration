package main

import (
	"context"
	pb "grpc-exploration/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Mayank",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
