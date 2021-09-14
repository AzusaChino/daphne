package main

import (
	"context"
	"github.com/AzusaChino/daphne/common"
	pb "github.com/AzusaChino/daphne/vessel-service/proto/vessel"
	"github.com/micro/micro/v3/service"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

const serviceName = "daphne.service.vessel"

func main() {
	srv := service.New(
		service.Name(serviceName))
	srv.Init()

	uri := os.Getenv("DB_HOST")
	client, err := common.CreateMongoClient(context.Background(), uri, true)

	if err != nil {
		log.Panic(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	consignmentCollection := client.Database("daphne").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	h := &handler{repository}

	if err = pb.RegisterVesselServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}
	if err = srv.Run(); err != nil {
		log.Panic(err)
	}
}
