package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var Platform *mongo.Client = nil

func InsertDocument(data map[string]interface{}) error {
	collection := Platform.Database("development").Collection("settings")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	fmt.Println("Document inserted successfully!")
	return nil
}

func InsertAudit(data map[string]interface{}) error {
	database := Platform.Database("development")
	audit := database.Collection("audit")
	_, err := audit.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}
