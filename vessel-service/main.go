package main

import (
	"context"
	pb "github.com/AzusaChino/daphne/vessel-service/proto/vessel"
	"github.com/AzusaChino/ribes/db"
	"github.com/asim/go-micro/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

const serviceName = "daphne.service.vessel"

func main() {
	srv := micro.NewService(
		micro.Name(serviceName))

	srv.Init()

	uri := os.Getenv("DB_HOST")
	client, err := db.CreateMongoClient(context.Background(), uri, true)

	if err != nil {
		log.Panic(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	vesselCollection := client.Database("daphne").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	if err = pb.RegisterVesselServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}
	if err = srv.Run(); err != nil {
		log.Panic(err)
	}
}
