package main

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: rand.Int31n(100),
		B: rand.Int31n(100),
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Sum: %v\n", res.Result)
}
