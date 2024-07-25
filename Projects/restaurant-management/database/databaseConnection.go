package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Collection {

	const connectionString = "mongodb://localhost:27017"
	const dbName = "restaurant"
	const collectionName = "food"

	// Reference of collection - Most important
	var mongoCollection *mongo.Collection

	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb collection success...")
	mongoCollection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection reference is ready...")

	return mongoCollection
}
