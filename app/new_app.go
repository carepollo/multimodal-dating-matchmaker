package app

import (
	"context"
	"log"
	"os"

	"github.com/carepollo/multimodal-dating-matchmaker/api"
	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Run() {
	server := models.App{
		Context: context.TODO(),
		Router:  fiber.New(),
	}

	defer server.DB.Disconnect(server.Context)

	//loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	//connecting to database
	connectOptions := options.Client().ApplyURI(os.Getenv("DB_STRING_CONNECTION"))
	server.DB, err = mongo.Connect(
		server.Context,
		connectOptions,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	//starting web server
	api.RegisterEndpoints(server.Router)
	server.Router.Listen(":8080")
}
