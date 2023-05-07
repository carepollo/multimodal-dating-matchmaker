package api

import (
	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
)

func (api *API) registerWithEmail(ctx *fiber.Ctx) error {
	body := new(models.User)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	body.Password = util.HashAndSalt(body.Password)
	if err := api.DB.CreateUser(*body); err != nil {
		return fiber.ErrBadRequest
	}

	if err := api.DB.CreateUser(*body); err != nil {
		return fiber.ErrInternalServerError
	}

	ctx.JSON(body) // might change in the future
	return nil
}
