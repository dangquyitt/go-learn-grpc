package main

import (
	"context"
	"log"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Printf("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if !ok {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

		log.Printf("Error message from server: %s\n", e.Message())
		log.Printf("Error code from server: %v\n", e.Code())

		if e.Code() == codes.InvalidArgument {
			log.Println("We probaly sent a negative number")
		}

		return
	}

	log.Printf("Sqrt: %v\n", res.Result)
}
