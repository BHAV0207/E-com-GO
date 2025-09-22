package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BHAV0207/E-com-GO/internal/database"
	"github.com/BHAV0207/E-com-GO/internal/models"
	"github.com/BHAV0207/E-com-GO/internal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	uri := "mongodb+srv://jainbhav0207:XosZWJgwpDfcoJ7M@cluster0.g5yofar.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	client := database.Connect(uri)
	defer client.Disconnect(context.Background())

	db := client.Database("ecommerce")
	productsCol := db.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product := models.Product{
		Name:        "Sample Product",
		Description: "A great product!",
		Price:       19.99,
		InStock:     100,
	}

	_, _ = services.InsertProduct(ctx, productsCol, product)

	//  fetching all products

	products, err := services.GetAllProducts(ctx, productsCol)
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return
	}

	fmt.Println("All products:")
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	// reteriving id
	productID, _ := primitive.ObjectIDFromHex("your_product_id_here")

	// update
	updateFields := bson.M{"price": 24.99, "in_stock": 50}
	updated, err := services.UpdateProductByID(ctx, productsCol, productID, updateFields)
	if err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Printf("Updated %d product(s)\n", updated)
	}

	// deleting
	deleted, err := services.DeleteProductByID(ctx, productsCol, productID)
	if err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Printf("Deleted %d product(s)\n", deleted)
	}
}
