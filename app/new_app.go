package app

import (
	"context"
	"log"
	"os"

	"github.com/carepollo/multimodal-dating-matchmaker/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	dbClient *mongo.Client
	ctx      = context.TODO()
)

func newRouter() *fiber.App {
	app := fiber.New()
	api.RegisterEndpoints(app)

	return app
}

func Run() {
	defer dbClient.Disconnect(ctx)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	connectOptions := options.Client().ApplyURI(os.Getenv("DB_STRING_CONNECTION"))
	dbClient, err = mongo.Connect(
		ctx,
		connectOptions,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	newRouter().Listen(":8080")
}
