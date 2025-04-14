package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/keepersobe/geofence/controller"
	"github.com/keepersobe/geofence/db"
)

type App struct {
	Router *gin.Engine
}

// Setup the Gin router and configure the routes.
func (app *App) Routes() {
	db.SetupDB()

	router := gin.Default()

	router.GET("/", controller.FetchGeoFences)

	app.Router = router
}

// Start the application listening on the port.
func (app *App) Run() {
	var port = os.Getenv("PORT");

	if len(port) == 0 {
		port = "8080"
	}

	app.Router.Run(fmt.Sprintf("localhost:%s", port));
}