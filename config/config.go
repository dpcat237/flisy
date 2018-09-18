package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var cfg Config

type Config struct {
	// For demo propose database parameters aren't required
	DbHost     string `envconfig:"DB_HOST"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbName     string `envconfig:"DB_NAME"`
	Port       string `envconfig:"ROUTER_PORT" default:"443"`
	CertFile   string `envconfig:"ROUTER_CERT_FILE" default:"config/ssl/server.rsa.crt"` // it's inside of the project only for the demo
	KeyFile    string `envconfig:"ROUTER_KEY_FILE" default:"config/ssl/server.rsa.key"`  // it's inside of the project only for the demo
}

func LoadConfigData() Config {
	return RefreshConfig()
}

func RefreshConfig() Config {
	if err := envconfig.Process("FLISY", &cfg); err != nil {
		panic(fmt.Sprintf("Failed reading environment variables: %s", err))
	}
	return cfg
}
