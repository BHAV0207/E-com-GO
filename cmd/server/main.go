package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/BHAV0207/E-com-GO/internal/database"
	"github.com/BHAV0207/E-com-GO/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	uri := "mongodb+srv://jainbhav0207:XosZWJgwpDfcoJ7M@cluster0.g5yofar.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	client := database.Connect(uri)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database("EcomServices")
	productsCol := db.Collection("products")

	productHandler := &handler.ProductHandler{Collection: productsCol}
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteById).Methods("DELETE")

	// Start server
	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
