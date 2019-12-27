package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "grpc/greet/greetpb"

func main() {
	fmt.Println("Hello I'm a Client!!")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close() // closing the connection , but this will call after or the end of the main

	c := greetpb.NewGreetServiceClient(clientConn)
	fmt.Printf("Created client: %f", c)

}
