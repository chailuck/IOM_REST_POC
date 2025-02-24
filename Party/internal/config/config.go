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
			Host:     "localhost",
			Port:     5432,
			Database: "partydb",
			Username: "party_user",
			Password: "party_password",
			SSLMode:  "disable",
		},
	}
}
