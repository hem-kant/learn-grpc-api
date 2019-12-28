package main

import "fmt"

import "net"

import "log"

import "google.golang.org/grpc"

import "grpc/greet/greetpb"

import "context"

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
	fmt.Println("Hello World!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()                            // creating a GRPC server
	greetpb.RegisterGreetServiceServer(s, &server{}) // register the service that we defined at line 15

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
