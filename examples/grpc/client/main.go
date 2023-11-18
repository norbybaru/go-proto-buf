package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/norbybaru/go-proto-buf/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("127.0.0.1:9292", opts...)

	if err != nil {
		log.Fatalf("fail dial: %v", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{FirstName: "Norby", LastName: "Baru"})

	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("%v\n", res)
}
