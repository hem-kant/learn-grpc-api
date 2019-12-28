package main

import (
	"context"
	"fmt"
	ErrorHandlingCalculatorpb "grpc/AdvanceTopics/HandlingErrorAndCodes/calculatorpb"
	"math"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *ErrorHandlingCalculatorpb.SquareRootRequest) (*ErrorHandlingCalculatorpb.SquareRootResponse, error) {

	fmt.Printf("Recevied SquareRoot RPC: %v\n", req)
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recevied a negative number: %v", number),
		)
	}
	return &ErrorHandlingCalculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}
func main() {
	fmt.Printf("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer()
	ErrorHandlingCalculatorpb.RegisterErrorHandlingSquareRootServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
