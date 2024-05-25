package main

import (
	"log"
	"net"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}