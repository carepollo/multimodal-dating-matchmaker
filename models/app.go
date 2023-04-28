package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB          *mongo.Database
	Ctx         context.Context
	Router      *fiber.App
	Environment Environment
}
