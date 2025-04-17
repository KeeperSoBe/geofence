package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/keepersobe/geofence/model"
	"github.com/keepersobe/geofence/service"
)

const BadRequestErrorMessage string = "Bad request error"
const NotFoundErrorMessage string = "Not found error"
const InternalServerErrorMessage string = "Internal server error"

type GeoFenceControllerInterface interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type GeoFenceController struct {
	Service service.GeoFenceServiceInterface
}

// Creates a new instance of the GeoFenceController.
func NewGeoFenceController(db *mongo.Collection) GeoFenceControllerInterface {
	return &GeoFenceController{
		Service: service.NewGeoFenceService(db),
	}
}

// Defines the shape of an error response payload.
type ErrorResponse struct {
	statusCode uint
	message string
}

// Error handler for bad request errors, responds to the request with a generic bad request message and code.
func handleBadRequestError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		statusCode: http.StatusBadRequest,
		message: BadRequestErrorMessage,
	})
}

// Error handler for not found errors, responds to the request with a generic not found message and code.
func handleNotFoundError(c *gin.Context) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		statusCode: http.StatusNotFound,
		message: NotFoundErrorMessage,
	})
}

// Error handler for internal server errors, responds to the request, logs out the error and exits the program.
func handleInternalServerError (err error, c *gin.Context) {
	c.String(http.StatusInternalServerError, InternalServerErrorMessage)
	log.Fatal(err)
}

// Returns all GeoFences entities.
func (g *GeoFenceController) List (c *gin.Context) {
	if results, err := g.Service.List(); err != nil {
		handleInternalServerError(err, c)
		return
	} else {
		c.JSON(http.StatusOK, results)
		return
	}
}

// Gets a single GeoFence entity by its id.
func (g *GeoFenceController) Get (c *gin.Context) {
	// Get the ID param from the request.
	var id, _ = c.Params.Get("id")

	// Validate the ID is a valid UUID.
	if err := uuid.Validate(id); err != nil {
		handleBadRequestError(c)
		return
	}

	if result, err := g.Service.Get(id); err != nil {
		handleInternalServerError(err, c)
	} else if result == nil {
		handleNotFoundError(c)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// Creates a new GeoFence entity.
func (g *GeoFenceController) Create (c *gin.Context) {
    var createPointDto = &model.CreatePointDto{}

	// Bind the JSON request.
	if err := c.ShouldBindJSON(createPointDto); err != nil {
		handleBadRequestError(c)
		return
	}

	// Validate the payload.
	if validationErrors := model.ValidateCreatePointDto(createPointDto); validationErrors != nil {
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	var point = model.CreateNewPoint(*createPointDto)

	// Create the entity.
	if _, err := g.Service.Create(point); err != nil {
		handleInternalServerError(err, c)
	} else {
		c.JSON(http.StatusCreated, point)
	}
}

// Patch updates a GeoFence entity by its id.
func (g *GeoFenceController) Update (c *gin.Context) {
	// Get the ID param from the request.
	var id, _ = c.Params.Get("id")

	// Validate the ID is a valid UUID.
	if err := uuid.Validate(id); err != nil {
		handleBadRequestError(c)
		return
	}

	var updatePointDto = model.UpdatePointDto{}
	
	// Bind the JSON request.
	if err := c.ShouldBindJSON(&updatePointDto); err != nil {
		handleBadRequestError(c)
		return
	}

	// Validate the payload
	var model, validationErrors = model.ValidateUpdatePointDto(&updatePointDto)

	// Respond to the request with any validation errors, if there are any.
	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	if updated, err := g.Service.Update(id, *model); err != nil {
		handleInternalServerError(err, c)
	} else if !updated {
		handleNotFoundError(c)
	} else {
		c.JSON(http.StatusOK, model)
	}
}


// Deletes a GeoFence entity by its id.
func (g *GeoFenceController) Delete (c *gin.Context) {
	// Get the ID param from the request.
	var id, _ = c.Params.Get("id")

	// Validate the ID is a valid UUID.
	if err := uuid.Validate(id); err != nil {
		handleBadRequestError(c)
		return
	}

	if result, err := g.Service.Delete(id); err != nil {
		handleInternalServerError(err, c)
	} else {
		if result {
			c.JSON(http.StatusOK, map[string]string{ "deletedAt": time.Now().Format(time.RFC3339) })
		} else {
			handleNotFoundError(c)
		}
	}	
}