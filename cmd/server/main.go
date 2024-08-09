package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	H "github.com/fjod/golang_advanced_course/internal/Handlers"
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
	router.POST("/update/:type/:name/:value", func(context *gin.Context) {
		H.Update(context, &storage.StorageOperations)
	})
	router.GET("/value/:type/:name", func(context *gin.Context) {
		H.Get(context, &storage.StorageOperations)
	})
	router.GET("", func(context *gin.Context) {
		H.Html(context, &storage.StorageOperations)
	})
	err = router.Run(server)
	if err != nil {
		fmt.Println("router dead", err)
		return
	}
}
