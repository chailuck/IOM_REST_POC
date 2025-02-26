package config

type Config struct {
	PostgresConfig PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	SSLMode  string
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
	}
}
