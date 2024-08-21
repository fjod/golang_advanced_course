package Middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func RequestLogger(sugar zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		sugar.Infoln(
			"uri", c.Request.RequestURI,
			"method", c.Request.Method,
			"duration", latency,
		)
	}
}
