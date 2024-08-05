package main

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type Config struct {
	ReportInterval int    `env:"REPORT_INTERVAL"`
	Address        string `env:"ADDRESS"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

func GetConfigValues() (server string, report int, poll int) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	s := flag.String("a", "localhost:8080", "отвечает за адрес эндпоинта HTTP-сервера (по умолчанию localhost:8080)")
	r := flag.Int("r", 10, "частота отправки метрик на сервер (по умолчанию 10 секунд)")
	p := flag.Int("p", 2, " частоту опроса метрик из пакета runtime (по умолчанию 2 секунды)")

	// разбор командной строки
	flag.Parse()

	if cfg.ReportInterval != 0 {
		*r = cfg.ReportInterval
	}
	if cfg.PollInterval != 0 {
		*p = cfg.PollInterval
	}
	if len(cfg.Address) != 0 {
		*s = cfg.Address
	}

	return *s, *r, *p
}
