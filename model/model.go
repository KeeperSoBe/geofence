package model

import "github.com/gin-gonic/gin"

type GeoFenceController interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type Point struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Radius uint `json:"radius"`
	Coordinates [2]float32 `json:"coordinates"`
}