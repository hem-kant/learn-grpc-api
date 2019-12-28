package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "grpc/greet/greetpb"

import "context"

func main() {
	fmt.Println("Hello I'm a Client!!")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close() // closing the connection , but this will call after or the end of the main

	c := greetpb.NewGreetServiceClient(clientConn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do a Unary RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hem",
			LastName:  "Kant",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Greet RPCt: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
