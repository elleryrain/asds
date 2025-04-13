package config

type Gateway struct {
	Dadata Dadata `envPrefix:"DADATA_"`
}

type Dadata struct {
	APIKey string `env:"API_KEY,required"`
}
