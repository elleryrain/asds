package config

type Transport struct {
	HTTP HTTP `envPrefix:"HTTP_"`
}

type HTTP struct {
	Port        int    `env:"PORT,required"`
	SwaggerPath string `env:"SWAGGER_BUNDLE_PATH"`
}
