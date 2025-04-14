package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/", fetchGeoFences)
	router.Run("localhost:8080")
}