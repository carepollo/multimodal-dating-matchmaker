package api

import "github.com/gofiber/fiber/v2"

func (api *API) RegisterWithGoogle(ctx *fiber.Ctx) error {
	ctx.SendString("lol")
	return nil
}
