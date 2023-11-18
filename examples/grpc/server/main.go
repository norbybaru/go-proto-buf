package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userpb "github.com/norbybaru/go-proto-buf/gen/go/user/v1"
	"google.golang.org/grpc"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
}

func (u *UserService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     "12345432",
			FullName: req.FirstName + " " + req.LastName,
		},
	}, nil
}

func main() {
	address := "127.0.0.1:9292"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &UserService{})

	fmt.Println("gRPC listening on", address)
	grpcServer.Serve(lis)
}
