package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/sourabhsikarwar/go_movie_api/db"
	"github.com/sourabhsikarwar/go_movie_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This file contains all the Database controllers
// These are useful for the controllers functions

// ------------- Modifications Endpoints --------------------
func InsertMovie(movie models.Movies) (string, error) {
	if movie.Movie == "" {
		return "", fmt.Errorf("no movie title provided")
	}
	inserted, err := db.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("Movie inserted with ID: %v", inserted.InsertedID)
	return message, nil
}

func UpdateMovie(movieId string) {
	if movieId == "" {
		fmt.Println("No movie ID provided")
		return
	}
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie updated", id, res.ModifiedCount)
}

func DeleteMovie(movieId string) {
	if movieId == "" {
		fmt.Println("No movie ID provided")
		return
	}
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deletedCount, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted", id, deletedCount)
}

func DeleteMovies(movieIds []string) {
	if len(movieIds) == 0 {
		fmt.Println("No movies to delete")
		return
	}
	var idList []primitive.ObjectID

	for _, id := range movieIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatal(err)
		}
		idList = append(idList, oid)
	}

	filter := bson.M{"_id": bson.M{"$in": idList}}
	res, err := db.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies deleted", res.DeletedCount)
}

func DeleteAllMovies() {
	res, err := db.Collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Movies deleted", res.DeletedCount)
}

// ----------------------------------

// ----------------- Reading Endpoints ---------------

func GetAllMovies() []primitive.M {
	// Here mongo db returns a cursor object from which we can loop and get individual object or data
	cursor, err := db.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	// Looping until the curson.Next has a valid value
	// This works like a while loop
	for cursor.Next(context.Background()) {
		var movie bson.M

		// passing reference of movie variable to store the error value for checking
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	// closing the context for the cursor
	defer cursor.Close(context.Background())
	return movies
}
