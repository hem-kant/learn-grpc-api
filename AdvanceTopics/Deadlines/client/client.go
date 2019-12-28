package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	greetDeadlinepb "grpc/AdvanceTopics/Deadlines/greetpb"
)

func main() {
	fmt.Println("Hello I'm a Client!!")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConn.Close() // closing the connection , but this will call after or the end of the main

	c := greetDeadlinepb.NewGreetDeadlineServiceClient(clientConn)
	doUnaryDeadline(c, 5*time.Second) // should complete
	doUnaryDeadline(c, 1*time.Second) // DeadLine Error
}

func doUnaryDeadline(c greetDeadlinepb.GreetDeadlineServiceClient, Timeout time.Duration) {
	fmt.Printf("Starting to do a Unary Deadline RPC..")
	req := &greetDeadlinepb.GreetDeadRequest{
		Greeting: &greetDeadlinepb.Greeting{
			FirstName: "Hem",
			LastName:  "Kant",
		},
	}
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	res, err := c.GreetDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit!! Deadline was exceeded")
			} else {
				fmt.Println("Unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("Error While calling Greet Deadline RPCt: %v", err)
		}

	}
	log.Printf("Response from GreetDeadline: %v", res.Result)
}
