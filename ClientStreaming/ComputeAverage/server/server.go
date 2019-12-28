package main

import (
	"fmt"
	computeaveragepb "grpc/ClientStreaming/ComputeAverage/ComputeAveragepb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"io"
	"log"
	"net"
)

type server struct{}

func (*server) ComputeAverage(stream computeaveragepb.ComputeAverageService_ComputeAverageServer) error {
	fmt.Printf("Recevied ComputeAverage RPC \n")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) //
			return stream.SendAndClose(&computeaveragepb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %V", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func main() {
	fmt.Printf("Compute Average Service Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
	)
	computeaveragepb.RegisterComputeAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
