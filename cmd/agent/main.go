package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	time.Sleep(2 * time.Second)
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}

	server, reportInterval, pollInterval := GetConfigValues()
	fmt.Println("agent запущен на ", server)
	SendMetrics(server, reportInterval, pollInterval)
}
