package controllers

import (
	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
)

func RegisterWithEmail(ctx *fiber.Ctx) error {
	body := new(models.User)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	body.Password = util.HashAndSalt(body.Password)

	ctx.JSON(body)
	return nil
}
