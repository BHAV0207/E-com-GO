package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
So when we say:
“mongo.Client is a struct provided by the driver, which represents a pool of connections to a MongoDB deployment.”

We mean:
It’s a data structure (struct) from the MongoDB Go Driver.
Inside, it manages a pool of sockets to MongoDB (instead of you managing them manually).
It knows how to talk to any type of MongoDB server setup.
You, as a developer, just use this Client to run queries — the driver handles connection reuse, retries, load-balancing, etc.

💡 Analogy:
Think of mongo.Client like a manager at a restaurant:
You don’t talk to each waiter (connection) directly.
You talk to the manager (mongo.Client).
The manager decides which waiter (connection from the pool) should serve your table (query).
You don’t care about the internal details — the manager just makes sure you get served efficiently.
*/
var Client *mongo.Client

func Connect(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}
	Client = client
	return client
}
