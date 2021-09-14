package main

import (
	"context"
	pb "github.com/AzusaChino/daphne/user-service/proto/user"
	"github.com/AzusaChino/ribes/db"
	"github.com/asim/go-micro/v3"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

const serviceUser = "daphne.service.user"

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

var postgres *sqlx.DB

func init() {
	postgres, err := db.NewConnection()
	if err != nil {
		log.Fatalf("could not conenct to DB: %v", err)
	}
	postgres.MustExec(schema)
}

func main() {
	defer func(postgres *sqlx.DB) {
		err := postgres.Close()
		if err != nil {

		}
	}(postgres)

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

	repo := NewPostgresRepository(postgres)
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name(serviceUser),
		micro.Version("latest"))
	srv.Init()

	h := &handler{repository: repo, tokenService: tokenService}
	if err = pb.RegisterUserServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}

	if err = srv.Run(); err != nil {
		log.Panic(err)
	}
}
