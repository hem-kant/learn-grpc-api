package main

import (
	"context"
	"fmt"
	"grpc/BiDi/bidipb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("ComputeAverage Client..")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}

	defer clientConn.Close()

	c := bidipb.NewGreetEveryoneServiceClient(clientConn)

	doBiDiStreaming(c)

}
func doBiDiStreaming(c bidipb.GreetEveryoneServiceClient) {
	fmt.Printf("Starting to do a BiDI Client Sreaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating a stream: %V", err)
		return
	}
	requests := []*bidipb.GreetEveryoneRequest{
		&bidipb.GreetEveryoneRequest{
			Greeting: &bidipb.Greeting{
				FirstName: "Hem",
			},
		},
		&bidipb.GreetEveryoneRequest{
			Greeting: &bidipb.Greeting{
				FirstName: "Hem -1 ",
			},
		},
		&bidipb.GreetEveryoneRequest{
			Greeting: &bidipb.Greeting{
				FirstName: "Hem - 2 ",
			},
		},
		&bidipb.GreetEveryoneRequest{
			Greeting: &bidipb.Greeting{
				FirstName: "Hem - 3 ",
			},
		},
	}
	waitc := make(chan struct{})
	// we send a bunch of mnessages to the client (go routine)
	go func() {
		// to send the bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n ", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we recevie a bunch of messages from the client (go routine)
	go func() {
		// to recevie the bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receving a stream: %V", err)
				break
			}
			fmt.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()
	//block until everything is done
	<-waitc
}
