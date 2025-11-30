package main

import (
	"log"

	"github.com/brojyf/NetworkDashboard/internal/Handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// TODO: Remove cors & Build Frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type"},
	}))
	router.GET("/api/query", Handler.Handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
