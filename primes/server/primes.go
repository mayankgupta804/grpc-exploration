package main

import (
	pb "grpc-exploration/primes/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.PrimesService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	var c int64 = 2
	for in.Num > 1 {
		if in.Num%c == 0 {
			stream.Send(&pb.PrimesResponse{Num: c})
			in.Num = in.Num / c
		} else {
			c = c + 1
		}
	}
	return nil
}
