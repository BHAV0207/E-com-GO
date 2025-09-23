package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/BHAV0207/E-com-GO/internal/models"
	"github.com/BHAV0207/E-com-GO/internal/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductHandler struct {
	Collection *mongo.Collection
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := services.InsertProduct(ctx, h.Collection, product)
	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted product with ID: %v", id)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := services.GetAllProducts(ctx, h.Collection)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Tell client: "I’m sending JSON"
	json.NewEncoder(w).Encode(products)                // Actually send the JSON response

	// json.NewEncoder(w).Encode(products)
	// json.NewEncoder(w) → creates a JSON encoder that writes directly to the response w.
	// .Encode(products) → takes your products (a slice or struct) and converts it into JSON format, then sends it into the response body.
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) extracts the URL parameters from the incoming request (r).
	// It returns a map[string]string.
	vars := mux.Vars(r)
	idHex := vars["id"]

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var updateFields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateFields); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	modifiedCount, err := services.UpdateProductByID(ctx, h.Collection, id, updateFields)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Updated %d product(s)", modifiedCount)

}

func (h *ProductHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idHex := vars["id"]

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deletedCount, err := services.DeleteProductByID(ctx, h.Collection, id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted %d product(s)", deletedCount)
}
