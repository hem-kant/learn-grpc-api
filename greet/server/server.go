package main

import "fmt"

import "net"

import "log"

import "google.golang.org/grpc"

import "grpc/greet/greetpb"

type server struct{}

func main() {
	fmt.Println("Hello World!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
