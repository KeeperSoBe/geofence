package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Returns the current datetime in ISO 8601 format.
func createTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// Defines the shape of a Point entity.
type Point struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Radius uint `json:"radius"`
	Coordinates [2]float32 `json:"coordinates"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Defines the shape of a create point request body.
type CreatePointDto struct {
    Type  *string `json:"type" binding:"required"`
	Radius *uint `json:"radius" binding:"required"`
	Coordinates *[2]float32 `json:"coordinates" binding:"required"`
}

// Defines the shape of an update point request body.
type UpdatePointDto struct {
    Type  *string `json:"type,omitempty"`
	Radius *uint `json:"radius,omitempty"`
	Coordinates *[2]float32 `json:"coordinates,omitempty"`
}

// Creates and returns a new Point from a CreatePointDto.
func CreateNewPoint(createPointDto CreatePointDto) Point {
	var timestamp string = createTimestamp();

	return Point{
		ID: uuid.New().String(),
		Type: *createPointDto.Type,
		Radius: *createPointDto.Radius,
		Coordinates: *createPointDto.Coordinates,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
}

// Validates a create point dto, returns an array of validation errors if the dto is invalid.
func ValidateCreatePointDto(createPointDto *CreatePointDto) *[]string {
	var validationErrors = []string{}

	if *createPointDto.Type != "point" {
		validationErrors = append(validationErrors, "Invalid type")
	}

	if *createPointDto.Radius <= 0 {
		validationErrors = append(validationErrors, "Radius must be a positive value")
	}

	if len(validationErrors) > 0 {
		return &validationErrors
	} else {
		return nil
	}
}

// Validates an update point dto, returns an array of validation errors if the dto is invalid or the update model if it is.
func ValidateUpdatePointDto(updatePointDto *UpdatePointDto) (*map[string]any, *[]string) {
	var model = map[string]any{}
	var validationErrors = []string{}

	// Check if the type property was included in the update.
	if updatePointDto.Type != nil {
		var geoFenceType = *updatePointDto.Type;

		// If the type property fails validation add it to the validationErrors, otherwise add it to the update model.
		if geoFenceType != "point" {
			validationErrors = append(validationErrors, "Invalid type")
		} else {
			model["type"] = geoFenceType
		}
	}

	// Check if the radius property was included in the update.
	if updatePointDto.Radius != nil {
		var radius = *updatePointDto.Radius;

		// If the radius property fails validation add it to the validationErrors, otherwise add it to the update model.
		if radius <= 0 {
			validationErrors = append(validationErrors, "Radius must be a positive value")
		} else {
			model["radius"] = radius
		}
	}

	// Check if the coordinates property was included in the update.
	if updatePointDto.Coordinates != nil {
		var coordinates = *updatePointDto.Coordinates;

		var coordinatesOne float32 = coordinates[0]
		var coordinatesTwo float32 = coordinates[1]

		// If the coordinates property fails validation add it to the validationErrors, otherwise add it to the update model.
		if fmt.Sprintf("%T", coordinatesOne) != "float32" || fmt.Sprintf("%T", coordinatesTwo) != "float32" {
			validationErrors = append(validationErrors, "Invalid coordinates provided")
		} else {
			model["coordinates"] = [2]float32{coordinatesOne, coordinatesTwo}
		}
	}

	// If there are validation errors return them, otherwise return the update model.
	if len(validationErrors) > 0 {
		return nil, &validationErrors
	} else {
		model["updatedAt"] = createTimestamp()
		return &model, nil
	}
}