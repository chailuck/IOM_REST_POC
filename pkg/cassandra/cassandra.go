package cassandra

import (
	"github.com/gocql/gocql"
	"gitlab.com/ft25/iom/engine/am/internal/config"
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
