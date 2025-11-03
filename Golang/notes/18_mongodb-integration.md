# MongoDB Integration in Go

**📅 Created:** Database Series
**🏷️ Topics:** MongoDB, mongo-driver, CRUD Operations, BSON, Database Integration
**🔗 Related:** [14_json.md](./14_json.md), [16_building-rest-apis.md](./16_building-rest-apis.md), [Concepts/25mongoapi](../Concepts/25mongoapi/)

---

## Overview

MongoDB is a popular NoSQL database that stores data in flexible, JSON-like documents. Go's official MongoDB driver (`mongo-driver`) provides a powerful, idiomatic way to interact with MongoDB. Understanding MongoDB integration is essential for building modern web applications with flexible data models.

**Why MongoDB with Go:**
- Schema-less flexibility
- Horizontal scalability
- JSON-like documents (natural fit with Go structs)
- Official, well-maintained driver
- Production-ready for high-traffic applications

---

## Installation

```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
```

**Prerequisites:**
- MongoDB installed locally or MongoDB Atlas account
- MongoDB running on `localhost:27017` (default)

---

## Connection Setup

### Basic Connection

```go
package main

import (
    "context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Connection string
    connectionString := "mongodb://localhost:27017"

    // Client options
    clientOptions := options.Client().ApplyURI(connectionString)

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    // Disconnect when done
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()
}
```

**Reasoning:** `context.TODO()` provides a context for operations. In production, use `context.WithTimeout` for timeouts. Always disconnect when done.

### Get Database and Collection

```go
// Get database
db := client.Database("myapp")

// Get collection
collection := db.Collection("users")
```

**Reasoning:** Database and collection are created automatically when you first insert data. No need to explicitly create them.

---

## Context Package

MongoDB operations require context for timeouts, cancellation, and deadlines:

```go
import (
    "context"
    "time"
)

// Simple context
ctx := context.Background()

// With timeout (recommended)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

**Reasoning:** Context carries deadlines and cancellation signals. Essential for preventing operations from hanging forever.

---

## Models and BSON Tags

### Define Models

```go
import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
    Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}
```

**Tag Explanation:**
- `bson:"_id,omitempty"` - MongoDB field name (BSON format)
- `json:"_id,omitempty"` - JSON field name (for API responses)
- `omitempty` - Omit field if zero value
- `primitive.ObjectID` - MongoDB's unique identifier type

**Reasoning:** BSON tags control MongoDB field names. `primitive.ObjectID` is MongoDB's 12-byte unique identifier.

---

## CRUD Operations

### Create (Insert)

#### Insert One Document

```go
import "go.mongodb.org/mongo-driver/bson/primitive"

func insertOne(collection *mongo.Collection) {
    movie := Movie{
        Movie:   "The Shawshank Redemption",
        Watched: false,
    }

    result, err := collection.InsertOne(context.Background(), movie)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted ID:", result.InsertedID)
}
```

#### Insert Many Documents

```go
func insertMany(collection *mongo.Collection) {
    movies := []interface{}{
        Movie{Movie: "The Godfather", Watched: true},
        Movie{Movie: "Pulp Fiction", Watched: false},
        Movie{Movie: "Inception", Watched: true},
    }

    result, err := collection.InsertMany(context.Background(), movies)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted IDs:", result.InsertedIDs)
}
```

**Reasoning:** `InsertOne` returns the generated ObjectID. `InsertMany` requires `[]interface{}` for type compatibility.

### Read (Find)

#### Find One Document

```go
import "go.mongodb.org/mongo-driver/bson"

func findOne(collection *mongo.Collection) {
    var movie Movie

    filter := bson.M{"movie": "Inception"}

    err := collection.FindOne(context.Background(), filter).Decode(&movie)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("No document found")
            return
        }
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", movie)
}
```

#### Find All Documents

```go
func findAll(collection *mongo.Collection) []primitive.M {
    // Empty filter gets all documents
    cursor, err := collection.Find(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }

    var movies []primitive.M

    // Iterate through cursor
    for cursor.Next(context.Background()) {
        var movie bson.M
        err := cursor.Decode(&movie)
        if err != nil {
            log.Fatal(err)
        }
        movies = append(movies, movie)
    }

    // Close cursor when done
    defer cursor.Close(context.Background())

    return movies
}
```

**Reasoning:** `FindOne` returns a single result. `Find` returns a cursor for iterating multiple results. Always close cursors.

#### Find with Filter

```go
// Find watched movies
filter := bson.M{"watched": true}
cursor, _ := collection.Find(context.Background(), filter)

// Multiple conditions (AND)
filter := bson.M{
    "watched": true,
    "rating":  bson.M{"$gte": 8},
}

// OR condition
filter := bson.M{
    "$or": []bson.M{
        {"watched": true},
        {"rating": bson.M{"$gte": 9}},
    },
}
```

**Reasoning:** `bson.M` is `map[string]interface{}` for building filters. Use MongoDB query operators like `$gte`, `$or`, `$in`.

### Update

#### Update One Document

```go
func updateOne(collection *mongo.Collection, movieID string) {
    // Convert string ID to ObjectID
    id, err := primitive.ObjectIDFromHex(movieID)
    if err != nil {
        log.Fatal(err)
    }

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"watched": true}}

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Modified %d document(s)\n", result.ModifiedCount)
}
```

#### Update Many Documents

```go
func updateMany(collection *mongo.Collection) {
    filter := bson.M{"watched": false}
    update := bson.M{"$set": bson.M{"watched": true}}

    result, err := collection.UpdateMany(context.Background(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Modified %d document(s)\n", result.ModifiedCount)
}
```

**Reasoning:** Always use update operators like `$set`, `$inc`, `$push`. Converting string to ObjectID is necessary for querying by ID.

### Delete

#### Delete One Document

```go
func deleteOne(collection *mongo.Collection, movieID string) {
    id, _ := primitive.ObjectIDFromHex(movieID)

    filter := bson.M{"_id": id}

    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Deleted %d document(s)\n", result.DeletedCount)
}
```

#### Delete Many Documents

```go
func deleteMany(collection *mongo.Collection) int64 {
    // Empty filter deletes all
    result, err := collection.DeleteMany(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }

    return result.DeletedCount
}
```

**Reasoning:** `DeleteMany` with empty filter removes all documents. Use with caution!

---

## Complete MongoDB API Example

### Project Structure

```
mongoapi/
├── main.go
├── model/
│   └── models.go
├── controller/
│   └── controllers.go
└── router/
    └── routers.go
```

### models.go

```go
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
    ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
    Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}
```

### controllers.go

```go
package controller

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "yourmodule/model"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// Initialize MongoDB connection
func init() {
    clientOptions := options.Client().ApplyURI(connectionString)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("MongoDB connection success")

    collection = client.Database(dbName).Collection(collectionName)
    fmt.Println("Collection reference ready")
}

// Helper: Insert one movie
func insertOneMovie(movie model.Netflix) {
    inserted, err := collection.InsertOne(context.Background(), movie)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted ID:", inserted.InsertedID)
}

// Helper: Update one movie
func updateOneMovie(movieID string) {
    id, _ := primitive.ObjectIDFromHex(movieID)

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"watched": true}}

    result, _ := collection.UpdateOne(context.TODO(), filter, update)
    fmt.Printf("Modified %d document(s)\n", result.ModifiedCount)
}

// Helper: Delete one movie
func deleteOneMovie(movieID string) {
    id, _ := primitive.ObjectIDFromHex(movieID)
    filter := bson.M{"_id": id}

    deleteCount, _ := collection.DeleteOne(context.Background(), filter)
    fmt.Println("Deleted count:", deleteCount.DeletedCount)
}

// Helper: Delete all movies
func deleteAllMovies() int64 {
    result, err := collection.DeleteMany(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }
    return result.DeletedCount
}

// Helper: Get all movies
func getAllMovies() []primitive.M {
    cursor, err := collection.Find(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }

    var movies []primitive.M

    for cursor.Next(context.Background()) {
        var movie bson.M
        cursor.Decode(&movie)
        movies = append(movies, movie)
    }

    defer cursor.Close(context.Background())
    return movies
}

// API Handlers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    allMovies := getAllMovies()
    json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var movie model.Netflix
    err := json.NewDecoder(r.Body).Decode(&movie)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    insertOneMovie(movie)
    json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    updateOneMovie(params["id"])
    json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    deleteOneMovie(params["id"])
    json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    count := deleteAllMovies()
    json.NewEncoder(w).Encode(count)
}
```

### routers.go

```go
package router

import (
    "yourmodule/controller"
    "github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
    router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
    router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
    router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
    router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

    return router
}
```

### main.go

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "yourmodule/router"
)

func main() {
    fmt.Println("MongoDB API")

    r := router.Router()

    fmt.Println("Server starting on port 4000...")
    log.Fatal(http.ListenAndServe(":4000", r))
}
```

**Reasoning:** The `init()` function in controller runs before main, establishing the MongoDB connection. This pattern separates concerns: models define structure, controllers handle logic, routers define endpoints.

---

## Query Operators

### Comparison

```go
// Equal
bson.M{"age": 25}

// Greater than
bson.M{"age": bson.M{"$gt": 18}}

// Greater than or equal
bson.M{"age": bson.M{"$gte": 18}}

// Less than
bson.M{"age": bson.M{"$lt": 65}}

// Less than or equal
bson.M{"age": bson.M{"$lte": 65}}

// Not equal
bson.M{"status": bson.M{"$ne": "deleted"}}

// In array
bson.M{"category": bson.M{"$in": []string{"electronics", "books"}}}
```

### Logical

```go
// AND (implicit)
bson.M{"age": bson.M{"$gte": 18}, "status": "active"}

// OR
bson.M{
    "$or": []bson.M{
        {"age": bson.M{"$lt": 18}},
        {"age": bson.M{"$gte": 65}},
    },
}

// NOT
bson.M{"age": bson.M{"$not": bson.M{"$gte": 18}}}
```

---

## Summary

**Key Takeaways:**
- Use `mongo.Connect` to establish connection
- `context.Background()` or `context.WithTimeout()` for operations
- `bson.M` for filters and updates
- `primitive.ObjectID` for MongoDB IDs
- Always close cursors after iteration
- Use `$set`, `$inc` operators for updates
- `init()` function for connection setup
- BSON and JSON tags for model mapping

**Quick Reference:**
```go
// Connect
client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
collection := client.Database("db").Collection("coll")

// Insert
collection.InsertOne(ctx, document)
collection.InsertMany(ctx, documents)

// Find
collection.FindOne(ctx, filter).Decode(&result)
cursor, _ := collection.Find(ctx, filter)

// Update
collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"field": value}})
collection.UpdateMany(ctx, filter, update)

// Delete
collection.DeleteOne(ctx, filter)
collection.DeleteMany(ctx, filter)

// Filters
bson.M{"field": value}
bson.M{"field": bson.M{"$gt": 10}}
bson.M{"$or": []bson.M{{...}, {...}}}
```

**When to Use This Note:**
- Building APIs with MongoDB
- NoSQL database integration
- Document-based data storage
- Flexible schema requirements
- Microservices with MongoDB

---

**📝 Last Updated:** Database Series
**➡️ Related:** See [16_building-rest-apis.md](./16_building-rest-apis.md) for API development, [14_json.md](./14_json.md) for JSON handling
**🔗 Example Code:** [Concepts/25mongoapi](../Concepts/25mongoapi/)
