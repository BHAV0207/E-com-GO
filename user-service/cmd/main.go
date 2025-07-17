package main

import (
	"github.com/gin-gonic/gin"
	// "log"
)

func main() {
	router := gin.Default()

	router.GET("/health" , func(c *gin.Context) {
		c.JSON(200 , gin.H{"message" : "User Service is healthy"})
	})

	// log.Println("starting server on port 8001")
	// err := 	router.Run(":8001")
	// if err != nil {
	// 	log.Fatal("Failed to run server: ", err)
	// }
}