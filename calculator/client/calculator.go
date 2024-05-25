package main

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Printf("doCalculate was invoked")
	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		A: rand.Int31n(100),
		B: rand.Int31n(100),
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Sum: %v\n", res.Sum)
}
