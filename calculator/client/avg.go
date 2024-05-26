package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	reqs := []*pb.AvgRequest{
		{Number: rand.Int31n(10)},
		{Number: rand.Int31n(10)},
		{Number: rand.Int31n(10)},
		{Number: rand.Int31n(10)},
		{Number: rand.Int31n(10)},
	}

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}
