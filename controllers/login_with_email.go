package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func LoginWithEmail(c *fiber.Ctx) error {
	c.JSON(struct {
		A string `json:"a"`
	}{
		A: "a",
	})

	return nil
}
