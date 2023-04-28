package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB      *mongo.Client   // connection instance of mongodb
	Context context.Context //context required for the DB
	Router  *fiber.App
}
