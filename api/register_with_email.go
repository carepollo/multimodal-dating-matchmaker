package api

import (
	"net/http"

	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
)

func (api *API) RegisterWithEmail(ctx *fiber.Ctx) error {
	body := new(models.User)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	body.Password = util.HashAndSalt(body.Password)
	if err := api.DB.CreateUser(*body); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return err
	}

	ctx.JSON(body)
	return nil
}
