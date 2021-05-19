package mongodb

import (
	"context"
	"errors"

	"albums/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MovieModel represent a mgo database session with a movie data model.
type AlbumStoryModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the movies table.
func (m *AlbumStoryModel) All() ([]models.AlbumStory, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.AlbumStory{}

	// Find all movies
	movieCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = movieCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

// FindByID will be used to find a new movie registry by id
func (m *AlbumStoryModel) FindByID(id string) (*models.AlbumStory, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find movie by id
	var album = models.AlbumStory{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&album)
	if err != nil {
		// Checks if the movie was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &album, nil
}

// Insert will be used to insert a new movie registry
func (m *AlbumStoryModel) Insert(album models.AlbumStory) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), album)
}

// Delete will be used to delete a movie registry
func (m *AlbumStoryModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}