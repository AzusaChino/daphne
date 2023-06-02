package conf

import "time"

type DaphneConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Etcd     EtcdConfig     `yaml:"etcd"`
}

type PostgresConfig struct {
	User     string `yaml:"user" default:"postgres"`
	Database string `yaml:"database" default:"public"`
	Port     int    `yaml:"port" default:"5432"`
	SslMode  string `yaml:"sslMode" default:"disable"`
}

type EtcdConfig struct {
	Endpoints []string      `yaml:"endpoints"`
	Timeout   time.Duration `yaml:"timeout"`
}
