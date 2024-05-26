package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/dangquyitt/go-learn-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	var resBuilder strings.Builder

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: resBuilder.String(),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		resBuilder.WriteString(fmt.Sprintf("Hello %s\n", req.FirstName))
	}
}
