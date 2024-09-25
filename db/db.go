package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbConnection = "mongodb+srv://Sourabh:sourabh1812@movie01.fwvyy.mongodb.net/?retryWrites=true&w=majority&appName=movie01"
const dbName = "movies"
const collectionName = "movie-list"

// Important part -> Creating an mongoDB collection instance

var collection *mongo.Collection

// Connect with mongoDb

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(dbConnection)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Collection Instance Creation
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")
}
