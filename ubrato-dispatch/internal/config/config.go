package config

import (
	"os"
)

type Nats struct {
	Addr string
}

type Smtp struct {
	From  string
	Login string
	Pass  string
	Host  string
}

type Config struct {
	Nats Nats
	Smtp Smtp
}

func Load() *Config {
	return &Config{
		Nats: Nats{
			Addr: getEnv("NATS_ADDR"),
		},
		Smtp: Smtp{
			From:  getEnv("EMAIL_FROM"),
			Login: getEnv("EMAIL_LOGIN"),
			Pass:  getEnv("EMAIL_PASS"),
			Host:  getEnv("EMAIL_HOST"),
		},
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
