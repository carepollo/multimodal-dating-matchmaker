package api

import "github.com/gofiber/fiber/v2"

// handler to get profile data of user
func (api *API) getUserData(ctx *fiber.Ctx) error {
	return ctx.SendString("a")
}
