package configuration

import (
	"os"
)

type Configuration struct {
	Port       string
	DbUrl      string
	JaegerUrl  string
	SomeAppId  string
	SomeAppKey string
}

func Load() (cfg *Configuration, err error) {

	cfg.Port = os.Getenv("PORT")
	cfg.DbUrl = os.Getenv("DB_URL")
	cfg.JaegerUrl = os.Getenv("JAEGER_URL")
	cfg.SomeAppId = os.Getenv("SOME_APP_ID")
	cfg.SomeAppKey = os.Getenv("SOME_APP_KEY")

	return cfg, nil
}
