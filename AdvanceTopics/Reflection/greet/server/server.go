package main

import (
	"context"
	"fmt"
	"grpc/AdvanceTopics/Reflection/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	Fname := req.GetGreeting().GetFirstName()
	result := "Hello " + Fname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello, you can now test the Reflection!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()                            // creating a GRPC server
	greetpb.RegisterGreetServiceServer(s, &server{}) // register the service that we defined at line 15
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
