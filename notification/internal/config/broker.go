package config

type Broker struct {
	Nats Nats `envPrefix:"NATS_"`
}

type Nats struct {
	Address string `env:"ADDR,required"`
}
