package config

import (
	"log"
	"os"
)

var Config *config

type config struct {
}

func InitConfig() {
	Config = &config{}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
