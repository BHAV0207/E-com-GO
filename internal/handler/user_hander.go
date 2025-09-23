package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/BHAV0207/E-com-GO/internal/models"
	"github.com/BHAV0207/E-com-GO/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	Collection *mongo.Collection
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// user.Role = "user" default for now

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = h.Collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user registered successfully"))
}
