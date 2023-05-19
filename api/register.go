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
		return fiber.NewError(fiber.StatusBadRequest, "the received object doesn't match the required fields")
	}

	// check that email and password are not empty fields
	if !util.ValidateEmail(body.Email) || !util.ValidatePassword(body.Password) {
		return fiber.NewError(fiber.ErrBadRequest.Code, "email or password doesn't match the requirements")
	}

	// check that user doesn't exist already in database
	exists, _ := api.DB.GetUserByEmail(body.Email)
	if exists != nil {
		return fiber.NewError(fiber.ErrNotAcceptable.Code, "this email is already registered")
	}

	// hash password and create user on DB
	body.Password = util.HashAndSalt(body.Password)
	body.Status = models.PENDING
	if err := api.DB.CreateUser(*body); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "couldn't create user")
	}

	//TODO: send email to such email to validate address

	return ctx.SendStatus(fiber.StatusOK)
}
