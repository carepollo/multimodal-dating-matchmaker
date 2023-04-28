package controllers

import "github.com/gofiber/fiber/v2"

func LoginWithFacebook(ctx *fiber.Ctx) error {
	ctx.SendString("lol")
	return nil
}
