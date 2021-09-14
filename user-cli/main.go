package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"log"

	pb "github.com/AzusaChino/daphne/user-service/proto/user"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("daphne.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// print the response
	fmt.Println("Response: ", rsp.User)

	return nil
}

func main() {
	ctx := context.Background()
	srv := micro.NewService()
	srv.Init()
	u := &pb.User{
		Name:     "abc",
		Email:    "abc@protomail.com",
		Company:  "ABC.INC",
		Password: "abc1234",
	}
	if err := createUser(ctx, srv, u); err != nil {
		log.Fatal(err)
	}
}
