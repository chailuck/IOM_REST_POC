package config

type Config struct {
	PostgresConfig  PostgresConfig
	CassandraConfig CassandraConfig
	EncryptionKey   string
}

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	SSLMode  string
}

type CassandraConfig struct {
	Hosts    []string
	Keyspace string
	Username string
	Password string
}

func New() *Config {
	return &Config{
		PostgresConfig: PostgresConfig{
			Host:     "192.168.10.109",
			Port:     5432,
			Database: "new_om_poc",
			Username: "user",
			Password: "password",
			SSLMode:  "disable",
		},
		CassandraConfig: CassandraConfig{
			Hosts:    []string{"172.16.2.114"},
			Keyspace: "iom",
			Username: "omx_sa",
			Password: "omx_sa",
		},
		EncryptionKey: "your-32-byte-encryption-key-here",
	}
}
