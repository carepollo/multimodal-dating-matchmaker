package api

import (
	"fmt"

	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
)

func (api *API) registerWithGoogle(ctx *fiber.Ctx) error {
	state := fmt.Sprintf("%v", util.RandomInt(100))
	url := api.googleOauthConfig.AuthCodeURL(state)
	ctx.Redirect(url, fiber.StatusTemporaryRedirect)
	return nil
}

func (api *API) callbackRegisterWithGoogle(ctx *fiber.Ctx) error {

	return nil
}
