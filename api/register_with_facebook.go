package api

import "github.com/gofiber/fiber/v2"

func (api *API) RegisterWithFacebook(ctx *fiber.Ctx) error {
	ctx.SendString("lol")
	return nil
}
