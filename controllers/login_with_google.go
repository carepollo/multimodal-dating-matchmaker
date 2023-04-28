package controllers

import "github.com/gofiber/fiber/v2"

func LoginWithGoogle(ctx *fiber.Ctx) error {
	ctx.SendString("lol")
	return nil
}
