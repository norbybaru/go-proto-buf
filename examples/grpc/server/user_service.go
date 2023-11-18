package main

import (
	"context"

	userpb "github.com/norbybaru/go-proto-buf/gen/go/user/v1"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     "12345432",
			FullName: req.FirstName + " " + req.LastName,
		},
	}, nil
}
