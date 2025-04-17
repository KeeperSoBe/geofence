package service

import (
	"context"
	"errors"
	"time"

	"github.com/keepersobe/geofence/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Defines the duration a request to the database can last before being cancelled.
const MongoTimeoutDuration = 30*time.Second

type GeoFenceService struct {
	DB *mongo.Collection
}

type GeoFenceServiceInterface interface {
	List() (results *[]model.Point, error error)
	Get(id string) (result *model.Point, error error)
	Create(model model.Point) (created bool, error error)
	Update(id string, model map[string]any) (updated bool, error error)
	Delete(id string) (deleted bool, error error)
}

// Creates a new instance of the GeoFenceService.
func NewGeoFenceService(db *mongo.Collection) GeoFenceServiceInterface {
	return &GeoFenceService{ DB: db }
}

// Returns all entities.
func (s *GeoFenceService) List () (*[]model.Point, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	var cursor, err = s.DB.Find(ctx, bson.D{{}})
	
	var results = []model.Point{}

	if errors.Is(err, mongo.ErrNoDocuments) {
		// No entities, return an empty array.
		return &results, nil
	} else if  err != nil {
		// Internal server error.
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		// Internal server error.
		return nil, err
	}

	return &results, nil
}

// Returns a single entity by its id.
func (s *GeoFenceService) Get (id string) (*model.Point, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	var result model.Point

	if err := s.DB.FindOne(ctx, bson.D{{ Key: "id", Value: id }}).Decode(&result); errors.Is(err, mongo.ErrNoDocuments) {
		// Entity not found.
		return nil, nil
		} else if err != nil {
		// Internal server error.
		return nil, err
	}
	
	return &result, nil
}

// Creates a new entity.
func (s *GeoFenceService) Create (model model.Point) (created bool, error error) {
	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()
	
	if _, err := s.DB.InsertOne(ctx, model); err != nil {
		// Internal server error.
		return false, err
	} else {
		return true, nil
	}
}


// Patch updates an entity by its id.
func (s *GeoFenceService) Update (id string, model map[string]any) (updated bool, error error) {
	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	if result, err := s.DB.UpdateOne(ctx, bson.D{{ Key: "id", Value: id }}, bson.D{{ Key: "$set", Value: model }}); errors.Is(err, mongo.ErrNoDocuments) {
		// Entity not found.
		return false, nil
	} else if err != nil {
		// Internal server error.
		return false, err
	} else {
		return result.ModifiedCount == 1, nil
	}
}


// Deletes an entity by its id.
func (s *GeoFenceService) Delete (id string) (deleted bool, error error) {
	var ctx, cancel = context.WithTimeout(context.Background(), MongoTimeoutDuration)

	defer cancel()

	if result, err := s.DB.DeleteOne(ctx, bson.D{{ Key: "id", Value: id }}); errors.Is(err, mongo.ErrNoDocuments) {
		// Entity not found.
		return false, nil
	} else if err != nil {
		// Internal server error.
		return false, err
	} else {
		return result.DeletedCount == 1, nil
	}
}