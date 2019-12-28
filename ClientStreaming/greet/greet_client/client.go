package main

import (
	"context"
	"fmt"
	ClientStreamingpb "grpc/ClientStreaming/greet/greetpb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a Client!!")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close() // closing the connection , but this will call after or the end of the main

	c := ClientStreamingpb.NewLongGreetServiceClient(clientConn)

	doClientStreaming(c)
}

func doClientStreaming(c ClientStreamingpb.LongGreetServiceClient) {
	fmt.Printf("Starting to do a Client Streaming RPC..")
	requests := []*ClientStreamingpb.LongGreetRequest{
		&ClientStreamingpb.LongGreetRequest{
			Greeting: &ClientStreamingpb.Greeting{
				FirstName: "Hem",
			},
		},
		&ClientStreamingpb.LongGreetRequest{
			Greeting: &ClientStreamingpb.Greeting{
				FirstName: "Hem - 1",
			},
		},
		&ClientStreamingpb.LongGreetRequest{
			Greeting: &ClientStreamingpb.Greeting{
				FirstName: "Hem - 2",
			},
		},
		&ClientStreamingpb.LongGreetRequest{
			Greeting: &ClientStreamingpb.Greeting{
				FirstName: "Hem - 3",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}
	// Now, we iterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response : %v\n", res)
}
