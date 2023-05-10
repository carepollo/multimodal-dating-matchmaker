package api

import "github.com/gofiber/fiber/v2"

func (api *API) getUserData(ctx *fiber.Ctx) error {
	ctx.SendString("a")
	return nil
}
