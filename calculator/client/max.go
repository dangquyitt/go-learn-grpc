package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/dangquyitt/go-learn-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling max: %v\n", err)
	}

	wait := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			req := &pb.MaxRequest{Number: rand.Int31n(10)}
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
			}

			log.Printf("Max: %v\n", msg.Result)
		}
		close(wait)
	}()

	<-wait
}
