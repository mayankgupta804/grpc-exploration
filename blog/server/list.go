package main

import (
	"context"
	"fmt"
	pb "grpc-exploration/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		if err := cur.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil

}
