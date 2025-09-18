package storage // package name for storage-related code

import ( // import block for required packages
	"context" // standard library: provides context for timeouts, cancellation, deadlines
	"log"     // standard library: logging utilities
	"time"    // standard library: time durations and timestamps

	"go.mongodb.org/mongo-driver/mongo"         // MongoDB Go driver core types (Client, Database, Collection)
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDB Go driver options builder
)

type MongoStore struct { // MongoStore holds handles to MongoDB resources for reuse
	client     *mongo.Client     // top-level client; manages pooled connections to MongoDB
	Database   *mongo.Database   // handle to a specific database
	Collection *mongo.Collection // handle to a specific collection within the database
}

// Client *mongo.Client — the top-level client object; manages connections/pooling.
// Database *mongo.Database — a handle to a specific Mongo database.
// Collection *mongo.Collection — a handle to a specific collection (table-like).

func NewMongoStore(uri, dbName, collName string) *MongoStore { // constructor to create and initialize MongoStore
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // create a context with 10s timeout for connect
	defer cancel()                                                           // ensure context is canceled to free resources

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)) // connect to MongoDB using the provided URI
	if err != nil {                                                   // handle connection error
		log.Fatal(err) // fatal log and exit; stops the program if connection fails
	}
	db := client.Database(dbName)   // get a handle to the specified database
	coll := db.Collection(collName) // get a handle to the specified collection within the database

	return &MongoStore{ // return a pointer to the initialized MongoStore
		client:     client, // set client to the connected *mongo.Client
		Database:   db,     // set Database handle
		Collection: coll,   // set Collection handle
	}
}
