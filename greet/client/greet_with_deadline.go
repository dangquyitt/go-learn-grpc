package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dangquyitt/go-learn-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Johnny",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if !ok {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

		if e.Code() == codes.DeadlineExceeded {
			log.Println("Deadline exceeded!")
			return
		}

		log.Fatalf("Unexpected gRPC error: %v\n", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
