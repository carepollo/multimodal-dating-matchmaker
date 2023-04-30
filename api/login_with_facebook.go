package api

import "github.com/gofiber/fiber/v2"

func (api *API) LoginWithFacebook(ctx *fiber.Ctx) error {
	ctx.SendString("lol")
	return nil
}
