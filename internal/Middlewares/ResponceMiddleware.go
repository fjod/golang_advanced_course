package Middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ResponseLogger(sugar zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		sugar.Infoln(
			"status", c.Writer.Status(),
			"size", c.Writer.Size(),
		)
	}
}
