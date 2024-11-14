package config

type Config struct {
	PORT              string `env:"PORT" default:"3000"`
	BASE_URL          string `env:"BASE_URL" default:""`
	CONNECTION_STRING string `env:"CONNECTION_STRING" default:"postgresql://postgres:admin@localhost/postgres?sslmode=disable"`
}
