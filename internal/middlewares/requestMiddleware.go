package middlewares

import (
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func RequestLogger(sugar zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		var d data.Metrics
		err := c.ShouldBindJSON(&d)
		if err == nil {
			if d.MType == "counter" {
				sugar.Infoln(
					"data", d,
					"value", *d.Delta,
				)
			}
			if d.MType == "gauge" {
				sugar.Infoln(
					"data", d,
					"value", *d.Value,
				)
			}
		}

		c.Next()

		latency := time.Since(t)

		sugar.Infoln(
			"uri", c.Request.RequestURI,
			"method", c.Request.Method,
			"duration", latency,
		)
	}
}
