package main

import (
	"log"

	"github.com/brojyf/NetworkDashboard/internal/Handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/query", Handler.Handler)

	// Serve static
	router.Static("/assets", "./static/assets")
	router.StaticFile("/", "./static/index.html")

	// Fallback for SPA routing
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
