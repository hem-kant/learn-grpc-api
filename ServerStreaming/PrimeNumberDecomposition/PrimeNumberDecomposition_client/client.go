package main

import (
	"context"
	"fmt"
	"grpc/ServerStreaming/PrimeNumberDecomposition/PrimeNumberDecompositionpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Calculator Client..")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}

	defer clientConn.Close()

	c := PrimeNumberDecompositionpb.NewPrimeNumberDecompositionServiceClient(clientConn)

	doServerStreaming(c)
}

func doServerStreaming(c PrimeNumberDecompositionpb.PrimeNumberDecompositionServiceClient) {
	fmt.Printf("Starting to do a PrimeNumberDecomposition Server Streaming RPC...")
	req := &PrimeNumberDecompositionpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition: %V", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Something happened: %v", res)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
