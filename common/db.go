package common

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(ctx context.Context, uri string, retry bool) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil && retry {
		for cnt, err := 0, conn.Ping(ctx, nil); err != nil; cnt++ {
			if cnt >= 3 {
				return nil, err
			}
			time.Sleep(2 * time.Second)
			conn, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		}
	}
	return conn, err
}
