package main

import (
	"context"
	"fmt"

	pb "github.com/AzusaChino/daphne/user-service/proto/user"
	"github.com/micro/micro/v3/service"
)

func createUser(ctx context.Context, service *service.Service, user *pb.User) error {
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
	srv := service.New()
	srv.Init()
	u := &pb.User{
		Name:     "",
		Email:    "",
		Company:  "",
		Password: "",
	}
	_ = createUser(ctx, srv, u)
}
