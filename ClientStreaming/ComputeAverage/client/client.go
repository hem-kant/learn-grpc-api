package main

import (
	"context"
	"fmt"
	computeaveragepb "grpc/ClientStreaming/ComputeAverage/ComputeAveragepb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("ComputeAverage Client..")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}

	defer clientConn.Close()

	c := computeaveragepb.NewComputeAverageServiceClient(clientConn)

	doClientStreaming(c)

}
func doClientStreaming(c computeaveragepb.ComputeAverageServiceClient) {
	fmt.Printf("Starting to do a ComputeAverage Client Sreaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}
	numbers := []int32{10, 20, 30, 40, 50}
	for _, numbers := range numbers {
		fmt.Printf("Sending Number: %v\n", numbers)
		stream.Send(&computeaveragepb.ComputeAverageRequest{
			Number: numbers,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response: %v", err)
	}
	fmt.Printf("The Average is: %v\n", res.GetAverage())
}
