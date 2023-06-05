package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/azusachino/daphne/internal/global"
	"github.com/azusachino/daphne/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	v3 "go.etcd.io/etcd/client/v3"
)

const appName = "daphne"

func init() {
	// load dotenv
	godotenv.Load()

	// init config
	initConfig()
	// init db
	initDb()
	// init etcd client
	initEtcd()
}

func main() {
	// defers
	defer global.EtcdClient.Close()
	defer global.DbClient.Close()

	app := gin.Default()
	router.InitRouter(app)

	// set the schema to daphne
	// _, err := global.DbClient.Exec("set search_path to daphne;")
	// if err != nil {
	// 	panic(err)
	// }

	// query := `select tt.name, tt.consume_type, td.target from tb_dispatch td left join tb_topic tt on td.topic_id = tt.id;`
	// rows, err := global.DbClient.Query(query)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var consumeType model.ConsumeType
	// 	var target string
	// 	err = rows.Scan(&name, &consumeType, &target)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(name, consumeType, target)
	// }
	watchCh := global.EtcdClient.Watch(context.Background(), "daphne/DEFAULT", v3.WithPrefix())
	for watchResp := range watchCh {
		for _, event := range watchResp.Events {
			switch event.Type {
			case v3.EventTypePut:
				fmt.Println("Put")
			case v3.EventTypeDelete:
				fmt.Println("Delete")
			}
			fmt.Printf("Type: %s Key:%s Value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}

	app.Run()
}

func initConfig() {
	// load config from yaml file
	var err error
	vp := viper.New()
	vp.AddConfigPath("configs")
	vp.SetConfigName(appName)
	vp.SetConfigType("yaml")

	err = vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = vp.Unmarshal(&global.DaphneConfig)
	if err != nil {
		panic(err)
	}
}

func initDb() {
	var err error
	pgHost := os.Getenv(global.PG_HOST)
	pgPass := os.Getenv(global.PG_PASS)
	pgCfg := global.DaphneConfig.Postgres
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		pgCfg.User, pgPass, pgHost, pgCfg.Port, pgCfg.Database, pgCfg.SslMode)
	global.DbClient, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func initEtcd() {
	var err error
	etcdUserPasss := os.Getenv(global.ETCD_PASS)
	etcdCfg := v3.Config{
		Endpoints:   global.DaphneConfig.Etcd.Endpoints,
		DialTimeout: global.DaphneConfig.Etcd.Timeout,
		Username:    global.DaphneConfig.Etcd.User,
		Password:    etcdUserPasss,
	}

	global.EtcdClient, err = v3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
}
