package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Default struct {
	Debug     bool      `env:"DEBUG" envDefault:"false"`
	Store     Store     `envPrefix:"STORE_"`
	Transport Transport `envPrefix:"TRANSPORT_"`
	Gateway   Gateway   `envPrefix:"GATEWAY_"`
	Auth      Auth      `envPrefix:"AUTH_"`
	Broker    Broker    `envPrefix:"BROKER_"`
}

func ReadConfig() (Default, error) {
	var cfg Default

	if err := env.ParseWithOptions(&cfg, env.Options{
		RequiredIfNoDef: true,
	}); err != nil {
		return Default{}, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}
