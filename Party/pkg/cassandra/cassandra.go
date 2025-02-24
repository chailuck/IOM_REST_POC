package cassandra

import (
	"party/internal/config"

	"github.com/gocql/gocql"
)

func NewCassandraConnection(cfg config.CassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	return cluster.CreateSession()
}
