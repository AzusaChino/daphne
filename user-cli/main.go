package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/AzusaChino/daphne/user-service/proto/user"
	"github.com/micro/micro/v3/service"
)

func createUser(ctx context.Context, service interface{}, user *proto.User) error {
	client := proto.NewUserService("shippy.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// print the response
	fmt.Println("Response: ", rsp.User)

	return nil
}
