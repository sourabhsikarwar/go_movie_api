package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/sourabhsikarwar/go_movie_api/db"
	"github.com/sourabhsikarwar/go_movie_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ------------- Modifications Endpoints --------------------
func insertMovie(movie models.Movies) {
	if movie.Movie == "" {
		fmt.Println("No movie title provided")
		return
	}
	inserted, err := db.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie inserted with ID:", inserted.InsertedID)
}

func updateMovie(movieId string) {
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
	fmt.Println("Movie updated", res.ModifiedCount)
}

func deleteMovie(movieId string) {
	if movieId == "" {
		fmt.Println("No movie ID provided")
		return
	}
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	res, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted", res.DeletedCount)
}

func deleteMovies(movieIds []string) {
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

func deleteAllMovies() {
	res, err := db.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Movies deleted", res.DeletedCount)
}

// ----------------------------------

// ----------------- Reading Endpoints ---------------

func getAllMovies() []primitive.M {
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
