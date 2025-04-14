package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Point struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Radius uint `json:"radius"`
	Coordinates [2]float32 `json:"coordinates"`
}

func createPoint() Point {
	return Point{
		ID: "mock-uuid",
		Type: "Point",
		Radius: 1,
		Coordinates: [2]float32{1, 2},
	}
}

var geoFences = []Point{
	createPoint(),
}

func fetchGeoFences(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, geoFences)
}

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	  }

	router := gin.Default()

	router.GET("/", fetchGeoFences)


	var port = os.Getenv("PORT");

	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}