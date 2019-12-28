package main

import (
	"context"
	"fmt"
	"log"

	ErrorHandlingCalculatorpb "grpc/AdvanceTopics/HandlingErrorAndCodes/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Printf("Error Handling Square Root Client..")
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to listen: %V", err)
	}

	defer clientConn.Close()

	c := ErrorHandlingCalculatorpb.NewErrorHandlingSquareRootServiceClient(clientConn)

	doErrorUnaryHandling(c)
}

func doErrorUnaryHandling(c ErrorHandlingCalculatorpb.ErrorHandlingSquareRootServiceClient) {
	fmt.Printf("Error Handling Square Root RPC...")
	// correct call
	//number := int32(10)
	doErrorCalls(c, 10)

	// error call
	doErrorCalls(c, -2)
}

func doErrorCalls(c ErrorHandlingCalculatorpb.ErrorHandlingSquareRootServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &ErrorHandlingCalculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from grpc (user Error)
			fmt.Println("Error message from server: %V\n", n, respErr.Message())
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number")
			}
		} else {
			log.Fatalf("Error while calling Error Handling Square Root RPC: %V", respErr)
		}
	}
	fmt.Println("Result of square root of %V: %V", n, res.GetNumberRoot())

}
