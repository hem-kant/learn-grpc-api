package main

import (
	"context"
	"fmt"
	"grpc/AdvanceTopics/SSLEncryption/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	fmt.Println("Hello from SSL encrypted communication channel!!")

	lis, err := net.Listen("tcp", "localhost:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	certfile := "/Users/kanthem/go/src/grpc/AdvanceTopics/SSLEncryption/server.crt"
	keyfile := "/Users/kanthem/go/src/grpc/AdvanceTopics/SSLEncryption/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certfile, keyfile)
	if sslErr != nil {
		log.Fatalf("Failed loading certificates  %v", sslErr)
		return
	}
	opts := grpc.Creds(creds)

	s := grpc.NewServer(opts)                        // creating a GRPC server
	greetpb.RegisterGreetServiceServer(s, &server{}) // register the service that we defined at line 15

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
