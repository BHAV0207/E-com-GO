package repository

import (
	"context"
	// context: Think of it like a timeout manager — used to cancel operations if they take too long.
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string , dbName string) *mongo.Database{
	// context: Think of it like a timeout manager — used to cancel operations if they take too long.
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second )

/*What is context.Background()?
context.Background() creates an empty context.
It is like the root context — a starting point when you don’t have any existing context to work with.

Analogy:
Imagine context.Background() as a blank whiteboard:
When you start a project, you begin with an empty whiteboard.
Later, you can add timeout rules or cancellation signals on top of this "blank" context (like writing on the whiteboard).
*/

	defer cancel()

	client , err := mongo.Connect(ctx , options.Client().ApplyURI(uri))
	
	if err != nil{
		log.Fatal("Failed to connect to mongo " , err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB not reachable:", err)
	}

	log.Println("Connected to MongoDB!")
	return client.Database(dbName)
} 

