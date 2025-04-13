package config

type Broker struct {
	JetStream JetStream `envPrefix:"JETSTREAM_"`
}

type JetStream struct {
	Addr string `env:"ADDR"`
}
