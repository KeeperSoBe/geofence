package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Validate the database env vars and return them.
func getDbVars() (
	uri string,
	dbName string,
	dbUser string,
	dbPass string,
	) {
	uri = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")

	var valid = true

	if len(uri) == 0 {
		log.Fatal("You must set your 'DB_HOST' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
		valid = false
	} else if len(dbName) == 0 {
		log.Fatal("You must set your 'DB_NAME' environment variable")
		valid = false
	} else if len(dbUser) == 0 {
		log.Fatal("You must set your 'DB_USER' environment variable")
		valid = false
	} else if len(dbPass) == 0 {
		log.Fatal("You must set your 'DB_PASS' environment variable")
		valid = false
	}

	if !valid {
		os.Exit(1)
	}

	return uri, dbName, dbUser, dbPass
}

// Setup and confirm the database connection.
func SetupDB() {
	uri,
	dbName,
	dbUser,
	dbPass := getDbVars();

	credential := options.Credential{
		Username: dbUser,
		Password: dbPass,
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1.
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Set the Mongo credentials.
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetAuth(credential)
	// Create a new client and connect to the server.
	client, err := mongo.Connect(opts)
	
	if err != nil {
		panic(err)
	}
	
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection.
	var result bson.M

	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}