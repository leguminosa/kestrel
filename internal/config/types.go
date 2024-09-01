package config

type (
	Config struct {
		Server         ServerConfig
		DatabaseConfig DatabaseConfig
		Redis          RedisConfig
	}
	ServerConfig struct {
		Port string
	}
	DatabaseConfig struct {
		ConnectionString string
	}
	RedisConfig struct {
		Address string
	}
)
