package main

import (
	"log"
	"user-service/internal/config"
	"user-service/internal/repository"

	"github.com/gin-gonic/gin"
)

// The main function is the entry point of the program.
func main() {

	cfg := config.LoadConfig()

	db := repository.ConnectMongo(cfg.MongoURI, cfg.DBName)

	// Create a new Gin router instance with default middleware (Logger and Recovery).
	// - Logger: Logs every request to the console.
	// - Recovery: Recovers the server from panics (so it doesn't crash).
	router := gin.Default()

	// Define a GET endpoint at the route "/health".
	// this function (a handler) will be executed.
	router.GET("/health", func(c *gin.Context) {
		// Respond with a JSON object and HTTP status code 200 (OK).
		// gin.H is a shortcut for map[string]interface{} to create JSON data.
		c.JSON(200, gin.H{"message": "User Service is healthy"})
	})

	log.Println("starting server on port 8001")

	// Start the HTTP server on port 8001.
	// router.Run(":8001") blocks execution until the server is stopped or an error occurs.
	// If it fails (e.g., port is already in use), it returns an error.
	err := router.Run(":8001")

	if err != nil {
		log.Fatal("Failed to run server: ", err)
	}

	_ = db
}

/*
router := gin.Default()
What is gin.Default()?
gin.Default() is a function provided by the Gin framework.
It creates a new router instance (of type *gin.Engine) that will handle all HTTP requests.
This router is responsible for:
Registering routes (e.g., GET /health).
Mapping incoming HTTP requests to the correct handler functions.
Running middleware (functions that process requests before they reach your handler).

What does the “default” mean?
It comes pre-configured with two middleware:
Logger: Automatically logs every request (method, status code, duration, etc.).
Recovery: Ensures that if your code panics, the server won’t crash but will return a 500 Internal Server Error.

What is router here?
router is a variable of type *gin.Engine.
It’s your main application router — the "traffic controller" for your HTTP server.


Equivalent Code (Internally):
router := gin.New()       // Create a new empty router.
router.Use(gin.Logger())  // Attach the Logger middleware.
router.Use(gin.Recovery())// Attach the Recovery middleware.
*/

/*
Why *gin.Context (Pointer)?
1. Efficiency (Performance Reason)
Passing a pointer avoids copying the entire gin.Context struct.
gin.Context is a large struct containing request and response data, headers, query parameters, etc.
If we pass it by value, Go will create a new copy of this struct each time a request is processed.
Using a pointer means we just pass a memory address, which is much faster and uses less memory.


2. Mutability (Changing Data Inside Context)
When you pass a pointer (*gin.Context), you can modify the original object.
For example, if you set the response status code or headers in c, you want those changes to reflect in the same context that Gin is using for that HTTP request.

Example:
func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "ok"})
}
Here, c.JSON() modifies the original gin.Context object so the HTTP response is sent properly.
If it was passed by value, you'd modify a copy, and the server wouldn't see those changes.


3. Avoiding Extra Memory Allocation
Gin reuses the gin.Context object across multiple requests using object pooling (to improve performance).
Passing a pointer means the same memory can be reused, which reduces garbage collection overhead.

*/
