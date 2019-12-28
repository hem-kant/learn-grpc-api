package main

import (
	"fmt"
	"grpc/BiDi/bidipb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"io"
	"log"
	"net"
)

type server struct{}

func (*server) GreetEveryone(stream bidipb.GreetEveryoneService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming requests \n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %V", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&bidipb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %V", err)
			return err
		}
	}
}
func main() {
	fmt.Printf("BiDi Service Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
	)
	bidipb.RegisterGreetEveryoneServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
