package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Entry point of the program.
func main() {
	// Initialise dotenv.
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	  }

	router := gin.Default()

	router.GET("/", FetchGeoFences)
	router.GET("/:id", FetchGeoFences)
	router.POST("/", FetchGeoFences)
	router.PATCH("/:id", FetchGeoFences)
	router.DELETE("/:id", FetchGeoFences)

	var port = os.Getenv("PORT");

	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}