package api

import "github.com/gofiber/fiber/v2"

func RegisterEndpoints(app *fiber.App) {
	registerAuthRoutes(app)
	//registeruserRouter(app)
}
