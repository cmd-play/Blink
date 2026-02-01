package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs HTTP requests in a formatted way
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate request duration
		duration := time.Since(startTime)

		// Get request details
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.RequestURI
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Format log message
		logMessage := fmt.Sprintf(
			"[%s] %s %s | Status: %d | Duration: %v | IP: %s | User-Agent: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			statusCode,
			duration,
			clientIP,
			userAgent,
		)

		// Log based on status code
		if statusCode >= 500 {
			log.Printf("❌ %s\n", logMessage)
		} else if statusCode >= 400 {
			log.Printf("⚠️  %s\n", logMessage)
		} else if statusCode >= 300 {
			log.Printf("ℹ️  %s\n", logMessage)
		} else {
			log.Printf("✅ %s\n", logMessage)
		}
	}
}
