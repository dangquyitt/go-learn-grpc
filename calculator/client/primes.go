package main

import (
	"context"
	"io"
	"log"
	"math/rand"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: rand.Int63n(100000),
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
