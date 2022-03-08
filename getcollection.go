package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection("myUsers")
	return collection
}
