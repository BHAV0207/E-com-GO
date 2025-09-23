package services

import (
	"context"

	"github.com/BHAV0207/E-com-GO/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertProduct(ctx context.Context, collection *mongo.Collection, product models.Product) (interface{}, error) {
	result, err := collection.InsertOne(ctx, product)
	return result.InsertedID, err
}

func GetAllProducts(ctx context.Context, collection *mongo.Collection) ([]models.Product, error) {

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func UpdateProductByID(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID, updateFields bson.M) (int64, err) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateFields}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func DeleteProductByID(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
