package main

import (
	"context"
	"fmt"
	"grpc/AdvanceTopics/SSLEncryption/greet/greetpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Hello I'm a SSL Client!!")
	certfile := "/Users/kanthem/go/src/grpc/AdvanceTopics/SSLEncryption/ca.crt" // enable client to CA certifcate
	creds, sslErr := credentials.NewClientTLSFromFile(certfile, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certification: %v", sslErr)
		return
	}
	opts := grpc.WithTransportCredentials(creds)
	clientConn, err := grpc.Dial("localhost:50051", opts)
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
