package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/keepersobe/geofence/model"
)

const MissingIdQueryParamMessage string = "Request is missing ID query parameter"
const NotFoundError string = "Not found error"
const InternalServerError string = "Internal server error"
const MongoTimeoutDuration = 30*time.Second

type GeoFenceControllerInterface interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	// Update(*gin.Context)
	Delete(*gin.Context)
}

type GeoFenceController struct {
	DB *mongo.Collection
}

// Creates a new instance of the GeoFenceController.
func NewGeoFenceController(db *mongo.Collection) GeoFenceControllerInterface {
	return &GeoFenceController{ DB: db }
}


// Simple error handler for internal server errors, logs out the error and responds to the request.
func handleError (err error, c *gin.Context) {
	log.Fatal(err)
	c.String(http.StatusInternalServerError, InternalServerError)
}

// Returns all GeoFences records.
func (g *GeoFenceController) List (c *gin.Context) {
	var db = g.DB

	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	var cursor, err = db.Find(ctx, bson.D{})

	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusOK, [0]model.Point{})
		return
	} else if err != nil {
		handleError(err, c)
		return
	}

	var results []model.Point

	if err = cursor.All(ctx, &results); err != nil {
		handleError(err, c)
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusOK, [0]model.Point{})
		return
	}

	c.JSON(http.StatusOK, &results)
}

// Gets a single GeoFence record by its id.
func (g *GeoFenceController) Get (c *gin.Context) {
	var id, _ = c.Params.Get("id")

	if err := uuid.Validate(id); err != nil {
		c.String(http.StatusBadRequest, MissingIdQueryParamMessage)
		return
	}

	var db = g.DB

	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	var result model.Point

	err := db.FindOne(ctx, bson.D{{ Key: "id", Value: id }}).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		c.String(http.StatusNotFound, NotFoundError)
		return
	} else if err != nil {
		handleError(err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}

type CreatePointDto struct {
    Type  *string `json:"type" binding:"required"`
	Radius *uint `json:"radius" binding:"required"`
	Coordinates *[2]float32 `json:"coordinates" binding:"required"`
}



// func (g *GeoFenceController) Create (c *gin.Context) {
//     point := &CreatePointDto{}

// 	if err := c.ShouldBindJSON(point); err != nil {
// 		c.JSON(400, err.Error())
// 		return
// 	}

// 	if *point.Type != "point" {
// 		c.JSON(400, "Invalid type")
// 		return
// 	}

// 	if *point.Radius <= 0 {
// 		c.JSON(400, "Radius must be a positive value")
// 		return
// 	}

// 	// fmt.Printf("%+v\n", &point)
// 	c.JSON(200, point)



// 	// uuid := uuid.New()
// 	// uuid.String();
// }

func createNewPoint(createPointDto CreatePointDto) model.Point {
	return model.Point{
		ID: uuid.New().String(),
		Type: *createPointDto.Type,
		Radius: *createPointDto.Radius,
		Coordinates: *createPointDto.Coordinates,
	}
}

// Creates a new GeoFence record.
func (g *GeoFenceController) Create (c *gin.Context) {
    createPointDto := &CreatePointDto{}

	if err := c.ShouldBindJSON(createPointDto); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if *createPointDto.Type != "point" {
		c.JSON(400, "Invalid type")
		return
	}

	if *createPointDto.Radius <= 0 {
		c.JSON(400, "Radius must be a positive value")
		return
	}


	var db = g.DB

	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	var point = createNewPoint(*createPointDto)

	_, err := db.InsertOne(ctx, point)
	
	if err != nil {
		handleError(err, c)
		return
	}

	
	c.JSON(http.StatusCreated, point)
}

// Updates a GeoFence record by its id.
// func (g *GeoFenceController) Update (c *gin.Context) {}

// Deletes a GeoFence record by its id.
func (g *GeoFenceController) Delete (c *gin.Context) {
	var id, _ = c.Params.Get("id")

	if err := uuid.Validate(id); err != nil {
		c.String(http.StatusBadRequest, MissingIdQueryParamMessage)
		return
	}

	var db = g.DB

	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	result, err:= db.DeleteOne(ctx, bson.D{{ Key: "id", Value: id }})

	if errors.Is(err, mongo.ErrNoDocuments) {
		c.String(http.StatusNotFound, NotFoundError)
		return
	} else if result.DeletedCount == 0 {
		handleError(err, c)
		return
	} else if err != nil {
		handleError(err, c)
		return
	}
    
	// iso format
	// c.String(http.StatusOK, time.Now().Format(time.RFC3339))
	c.String(http.StatusOK, strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
}