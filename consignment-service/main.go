package main

import (
	"context"
	pb "github.com/AzusaChino/daphne/consignment-service/proto/consignment"
	vesselPb "github.com/AzusaChino/daphne/vessel-service/proto/vessel"
	"github.com/AzusaChino/ribes/db"
	"github.com/asim/go-micro/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

const serviceConsignment = "daphne.service.user"
const serviceVessel = "daphne.service.client"

func main() {
	srv := micro.NewService(
		micro.Name(serviceConsignment))
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

	consignmentCollection := client.Database("daphne").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselPb.NewVesselService(serviceVessel, srv.Client())
	h := &handler{repository: repository, vesselClient: vesselClient}

	if err = pb.RegisterShippingServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}

	if err = srv.Run(); err != nil {
		log.Panic(err)
	}
}
