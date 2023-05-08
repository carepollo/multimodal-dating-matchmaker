package api

import (
	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
)

// create a user in database with the given data model as long as it satisfies
// the structure of the User object.
func (api *API) registerWithEmail(ctx *fiber.Ctx) error {
	body := new(models.User)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	// check that email and password are not empty fields
	if util.ValidateEmail(body.Email) || util.ValidatePassword(body.Password) {
		ctx.SendStatus(fiber.ErrBadRequest.Code)
		return nil
	}

	// check that user doesn't exist already in database
	_, err := api.DB.GetUserByEmail(body.Email)
	if err == nil {
		ctx.SendStatus(fiber.ErrNotAcceptable.Code)
		return nil
	}

	// hash password
	body.Password = util.HashAndSalt(body.Password)
	if err := api.DB.CreateUser(*body); err != nil {
		ctx.SendStatus(fiber.ErrBadRequest.Code)
		return nil
	}

	// create user on DB
	body.Status = "Pending"
	if err := api.DB.CreateUser(*body); err != nil {
		ctx.SendStatus(fiber.ErrInternalServerError.Code)
		return nil
	}

	//TODO: send email to such email to validate address

	ctx.SendStatus(fiber.StatusOK)
	return nil
}
