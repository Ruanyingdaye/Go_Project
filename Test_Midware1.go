package main

import (
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize a zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Create a Gin router
	r := gin.New()

	// Use ginzap middleware for logging
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// Define a simple GET route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	r.Run(":8080")
}
