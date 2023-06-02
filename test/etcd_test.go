package test

import (
	"context"
	"testing"
	"time"

	v3 "go.etcd.io/etcd/client/v3"
)

func TestEtcd(t *testing.T) {
	ctx := context.TODO()
	key := "daphne/DEFAULT"
	client := initEtcd()
	pr, err := client.Put(ctx, key, "test")
	if err != nil {
		t.Error(err)
	}
	t.Logf("pr is: %v\n", pr)
}

func initEtcd() *v3.Client {
	cfg := v3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := v3.New(cfg)
	if err != nil {
		panic(err)
	}
	return client

}
