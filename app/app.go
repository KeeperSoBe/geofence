package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/keepersobe/geofence/controller"
	"github.com/keepersobe/geofence/db"
)

type App struct {
	DB *mongo.Collection
	Router *gin.Engine
}

// Connects to the DB and sets a reference to the collection in the app.
func (app *App) SetupDB() {
	app.DB = db.ConnectDB()	
}

// Setup the Gin router and configure the routes.
func (app *App) Routes() {
	var controller controller.GeoFenceControllerInterface = controller.NewGeoFenceController(app.DB)

	router := gin.Default()

	router.GET("/", controller.List)
	router.GET("/:id", controller.Get)
	router.POST("/", controller.Create)
	router.PATCH("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)

	app.Router = router
}

// Start the application listening on the port.
func (app *App) Run() {
	var port string = os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	app.Router.Run(fmt.Sprintf("localhost:%s", port))
}