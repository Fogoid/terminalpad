package mongodb

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbCollection *mongo.Collection
var mutex *sync.Mutex = &sync.Mutex{}

func GetDbCollection() *mongo.Collection {

	if dbCollection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if dbCollection == nil {
			return startDB()
		}
	}

	return dbCollection
}

func startDB() *mongo.Collection {
	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DATABASE")
	collection := os.Getenv("MONGODB_COLLECTION")

	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	dbCollection = dbClient.Database(database).Collection(collection)
    return dbCollection
}

func Disconnect() {
    err := dbCollection.Database().Client().Disconnect(context.TODO())
    if err != nil {
        log.Fatalf("Error closing database connection: %v", err)
    }
}
