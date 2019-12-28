package main

import (
	"fmt"
	"grpc/ProjectBloggingServiceUsingMongoAndGrpc/blogpb"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // If we crash the go code, we get the file name and line number
	fmt.Println("Blog Service Started!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)                   // creating a GRPC server
	blogpb.RegisterBlogServiceServer(s, &server{}) // register the service that we defined at line 15
	// stop server proprely
	go func() {
		fmt.Println("Starting Server!!")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	//  wait for channel C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is recevied
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("closing the listner")
	lis.Close()

	fmt.Println("END of the program")
}
