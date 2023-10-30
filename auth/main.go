package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/db"
	"github.com/carepollo/multimodal-dating-matchmaker/auth/handlers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createDgraphClient() (*dgo.Dgraph, error) {
	conn, err := grpc.Dial("localhost:9121", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to DGraph!")

	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	), nil
}

func createMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func main() {
	app := fiber.New()

	client, err := createDgraphClient()
	if err != nil {
		log.Fatal(err)
	}
	db.Users = client

	otherclient, err := createMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	db.Platform = otherclient

	app.Post("/register", handlers.HandleCreateUser)
	app.Listen(":3000")
}
