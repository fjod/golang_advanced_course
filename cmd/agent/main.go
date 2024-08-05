package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("agent")
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}
	server := flag.String("a", "localhost:8080", "отвечает за адрес эндпоинта HTTP-сервера (по умолчанию localhost:8080)")
	reportInterval := flag.Int("r", 10, "частота отправки метрик на сервер (по умолчанию 10 секунд)")
	pollInterval := flag.Int("p", 10, " частоту опроса метрик из пакета runtime (по умолчанию 2 секунды)")

	// разбор командной строки
	flag.Parse()
	SendMetrics(*server, *reportInterval, *pollInterval)
}
