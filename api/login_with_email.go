package api

import (
	"github.com/gofiber/fiber/v2"
)

func (api *API) LoginWithEmail(c *fiber.Ctx) error {
	c.JSON(struct {
		A string `json:"a"`
	}{
		A: "a",
	})

	return nil
}
