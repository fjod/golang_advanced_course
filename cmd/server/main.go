package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	"github.com/gin-gonic/gin"
	"os"
)

var storage = internal.NewStorage()

func main() {
	server := GetConfigValues()
	fmt.Println("server запущен на ", server)
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}
	router := gin.Default()
	router.POST("/update/:type/:name/:value", update)
	router.GET("/value/:type/:name", get)
	router.GET("", html)
	err = router.Run(server)
	if err != nil {
		fmt.Println("router dead", err)
		return
	}
}
