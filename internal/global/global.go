package global

import (
	"database/sql"

	"github.com/azusachino/daphne/pkg/conf"
	v3 "go.etcd.io/etcd/client/v3"
)

var (
	// DaphneConfig is the global config
	DaphneConfig *conf.DaphneConfig

    // Db Client
    DbClient *sql.DB

    // Etcd Client
	EtcdClient *v3.Client
)
