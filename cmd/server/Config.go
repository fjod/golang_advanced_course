package main

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type Config struct {
	Address string `env:"ADDRESS"`
}

func GetConfigValues() (server string) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	s := flag.String("a", "localhost:8080", "отвечает за адрес эндпоинта HTTP-сервера (по умолчанию localhost:8080)")

	// разбор командной строки
	flag.Parse()

	if len(cfg.Address) != 0 {
		*s = cfg.Address
	}

	return *s
}
