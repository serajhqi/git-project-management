package config

type Config struct {
	PORT        string `env:"PORT" default:"3000"`
	BASE_URL    string `env:"BASE_URL" default:""`
	PG_HOST     string `env:"PG_HOST" default:"localhost:5432"`
	PG_USER     string `env:"PG_USER" default:"postgres"`
	PG_PASSWORD string `env:"PG_PASSWORD" default:"admin"`
	PG_DATABASE string `env:"PG_DATABASE" default:"postgres"`
}
