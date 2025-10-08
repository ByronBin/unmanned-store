package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unmanned-store/backend/pkg/logger"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		logger.Info("HTTP Request",
			"method", method,
			"path", path,
			"query", query,
			"status", statusCode,
			"latency", latency,
			"client_ip", clientIP,
		)
	}
}
