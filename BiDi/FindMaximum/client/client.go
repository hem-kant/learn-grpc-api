package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"grpc/BiDi/FindMaximum/findmaximumpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("ComputeAverage Client..")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}

	defer clientConn.Close()

	c := findmaximumpb.NewFindMaximumServiceClient(clientConn)

	doFindMaxBidiStreaming(c)

}

func doFindMaxBidiStreaming(c findmaximumpb.FindMaximumServiceClient) {
	fmt.Printf("Starting to do a Find Maximum BiDi Client Sreaming RPC...")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream and calling Find Maximum: %v", err)
	}
	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{10, 20, 1, 40, 50}
		for _, numbers := range numbers {
			fmt.Printf("Sending Number: %v\n", numbers)
			stream.Send(&findmaximumpb.FindMaximumRequest{
				Number: numbers,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream and calling Find Maximum: %v", err)
				break
			}
			maximum := res.GetMaximum()
			fmt.Printf("Recevied a new Max. of : %v\n", maximum)

		}
		close(waitc)
	}()
	<-waitc

}
