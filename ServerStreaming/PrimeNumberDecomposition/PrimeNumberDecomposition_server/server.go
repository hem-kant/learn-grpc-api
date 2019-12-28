package main

import (
	"fmt"
	"grpc/ServerStreaming/PrimeNumberDecomposition/PrimeNumberDecompositionpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *PrimeNumberDecompositionpb.PrimeNumberDecompositionRequest, stream PrimeNumberDecompositionpb.PrimeNumberDecompositionService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Recevied PrimeNumberDecomposition RPC: %V\n", req)
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&PrimeNumberDecompositionpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to: %v", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Printf("PrimeNumberDecomposition Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer()
	PrimeNumberDecompositionpb.RegisterPrimeNumberDecompositionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
