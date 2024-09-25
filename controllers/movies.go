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
