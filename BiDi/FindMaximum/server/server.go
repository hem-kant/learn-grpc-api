package main

import (
	"fmt"
	"grpc/BiDi/FindMaximum/findmaximumpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"io"
	"log"
	"net"
)

type server struct{}

func (*server) FindMaximum(stream findmaximumpb.FindMaximumService_FindMaximumServer) error {
	fmt.Printf("FindMaximum function was invoked with a streaming requests \n")
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %V", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&findmaximumpb.FindMaximumResponse{
				Maximum: maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %V", err)
				return err
			}
		}

	}
}
func main() {
	fmt.Printf("BiDi FindMaximum Service Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
	)
	findmaximumpb.RegisterFindMaximumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %V", err)
	}
}
