package main

import (
	"fmt"
	"log"
	"net"

	userpb "github.com/norbybaru/go-proto-buf/gen/go/user/v1"
	wearablepb "github.com/norbybaru/go-proto-buf/gen/go/wearable/v1"
	"google.golang.org/grpc"
)

func main() {
	address := "127.0.0.1:9292"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &UserServer{})
	wearablepb.RegisterWearableServiceServer(grpcServer, &WearableServer{})

	fmt.Println("gRPC listening on", address)
	grpcServer.Serve(lis)
}
