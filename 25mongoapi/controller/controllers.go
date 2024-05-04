package cotroller

// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
import (
	"context"
	"fmt"
	"log"
	"mongoapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchlist"

// Reference of collection - Most important
var collection *mongo.Collection

// connect with mongo - initialization method
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	// https://pkg.go.dev/context
	//  whenever we are making call outside our machine, then we need to provide the context for Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb collection success...")
	collection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection reference is ready...")
}

// mongodb helper methods - file

// TODO: insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted the 1 record of movie with id [", inserted.InsertedID, "]")
}

// TODO: update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updatedOne, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[", updatedOne.ModifiedCount, "] record updated with id", movieId)
}
