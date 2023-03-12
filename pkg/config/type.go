package config

type (
	Config struct {
		Port     string
		Database DatabaseConfig
	}
	DatabaseConfig struct {
		Master string
		Slave  string
	}
)
