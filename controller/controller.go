package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keepersobe/geofence/model"
)

func createPoint() model.Point {
	return model.Point{
		ID: "mock-uuid",
		Type: "Point",
		Radius: 1,
		Coordinates: [2]float32{1, 2},
	}
}

var geoFences = []model.Point{
	createPoint(),
}

func FetchGeoFences(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, geoFences)
}