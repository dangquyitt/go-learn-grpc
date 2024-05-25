package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dangquyitt/go-learn-grpc/greet/proto"
)

func (s *Server) Geet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.FirstName),
	}, nil
}
