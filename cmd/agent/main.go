package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	}))

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}

	server, reportInterval, pollInterval := GetConfigValues()
	fmt.Println("agent запущен на ", server)
	go SendMetrics(server, reportInterval, pollInterval)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
