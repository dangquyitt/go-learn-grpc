package main

import (
	"context"
	"log"

	pb "github.com/dangquyitt/go-learn-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Geet(context.Background(), &pb.GreetRequest{
		FirstName: "Johnny",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
