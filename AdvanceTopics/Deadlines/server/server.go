package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	greetDeadlinepb "grpc/AdvanceTopics/Deadlines/greetpb"
)

type server struct{}

func (*server) GreetDeadline(ctx context.Context, req *greetDeadlinepb.GreetDeadRequest) (*greetDeadlinepb.GreetDeadResponse, error) {

	fmt.Printf("Recevied GreetDeadline RPC: %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Printf("The Client cancelled the request:")
			return nil, status.Error(codes.Canceled, "Client cancelled the reques")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetDeadlinepb.GreetDeadResponse{
		Result: result,
	}
	return res, nil
}
func main() {
	fmt.Printf("GreetDeadline Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer()
	greetDeadlinepb.RegisterGreetDeadlineServiceServer(s, &server{}) // register the deadlineservice

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
