package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/keepersobe/geofence/app"
)

// Entry point of the program.
func main() {
	// Initialise dotenv.
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	  }


	var application app.App

	application.SetupDB()
	application.Routes()
	application.Run()
}