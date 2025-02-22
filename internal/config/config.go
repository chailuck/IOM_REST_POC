package config

type Config struct {
	CassandraConfig CassandraConfig
	EncryptionKey   string
}

type CassandraConfig struct {
	Hosts    []string
	Keyspace string
	Username string
	Password string
}

func New() *Config {
	return &Config{
		CassandraConfig: CassandraConfig{
			Hosts:    []string{"172.16.2.114"},
			Keyspace: "iom",
			Username: "omx_sa",
			Password: "omx_sa",
		},
		EncryptionKey: "your-32-byte-encryption-key-here",
	}
}
