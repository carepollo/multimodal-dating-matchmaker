package api

import (
	"os"
	"time"

	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// returns a token
func (api *API) loginWithEmail(ctx *fiber.Ctx) error {
	body := new(models.User)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "object doesn't have required properties")
	}

	// search user on DB to check for existance
	user, err := api.DB.GetUserByEmail(body.Email)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "user not found")
	}

	// check that user has confirmed the email address
	if user.Status == models.PENDING {
		return fiber.NewError(fiber.StatusForbidden, "account not confirmed")
	}

	// check given password if it is correct
	password := []byte(body.Password)
	passwordMatch := util.ComparePasswords(user.Password, password)
	if !passwordMatch {
		return fiber.NewError(fiber.ErrNotFound.Code, "user not found")
	}

	// create token for user access, should expire within 14 days
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 14).Unix(),
	})
	key := []byte(os.Getenv("JWT_KEY"))
	value, err := token.SignedString(key)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "couldn't generate token: "+err.Error())
	}

	// return user id and generated token
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": value,
		"id":    user.ID,
	})
}
