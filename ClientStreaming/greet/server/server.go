package main

import (
	"fmt"
	ClientStreamingpb "grpc/ClientStreaming/greet/greetpb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream ClientStreamingpb.LongGreetService_LongGreetServer) error {
	fmt.Println("LongGreet funcation was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&ClientStreamingpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}
func main() {
	fmt.Println("Hello World!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()                                          // creating a GRPC server
	ClientStreamingpb.RegisterLongGreetServiceServer(s, &server{}) // register the service that we defined at line 15

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
