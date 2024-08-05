package main

import (
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

	server, reportInterval, pollInterval := GetConfigValues()
	SendMetrics(server, reportInterval, pollInterval)
}
