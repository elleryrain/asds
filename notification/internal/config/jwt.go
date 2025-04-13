package config

type Auth struct {
	JWT JWT `envPrefix:"JWT_"`
}

type JWT struct {
	Secret string `env:"SECRET,required"`
}
