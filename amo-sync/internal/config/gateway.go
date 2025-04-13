package config

type Gateway struct {
	AmoCRM AmoCRM `envPrefix:"AMOCRM_"`
}

type AmoCRM struct {
	Token string `env:"TOKEN"`
}
