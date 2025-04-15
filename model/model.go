package model

type Point struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Radius uint `json:"radius"`
	Coordinates [2]float32 `json:"coordinates"`
}