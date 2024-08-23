package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	H "github.com/fjod/golang_advanced_course/internal/Handlers"
	MW "github.com/fjod/golang_advanced_course/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

var storage = internal.NewStorage()
var sugar zap.SugaredLogger

func main() {

	CreateLogger()

	server := GetConfigValues()
	fmt.Println("server запущен на ", server)
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}
	router := gin.Default()

	// Сведения о запросах должны содержать URI, метод запроса и время, затраченное на его выполнение.
	router.Use(MW.RequestLogger(sugar))

	// Сведения об ответах должны содержать код статуса и размер содержимого ответа.
	router.Use(MW.ResponseLogger(sugar))

	router.POST("/update/:type/:name/:value", func(context *gin.Context) {
		H.Update(context, &storage.StorageOperations)
	})
	router.POST("/value/", func(context *gin.Context) {
		H.GetJSON(context, &storage.StorageOperations)
	})
	router.GET("/value/:type/:name", func(context *gin.Context) {
		H.Get(context, &storage.StorageOperations)
	})
	router.GET("", func(context *gin.Context) {
		H.HTML(context, &storage.StorageOperations)
	})
	err = router.Run(server)
	if err != nil {
		fmt.Println("router dead", err)
		return
	}
}

func CreateLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("zap logger sync error", err)
		}
	}(logger)

	sugar = *logger.Sugar()
}
