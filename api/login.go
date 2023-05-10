package api

import (
	"os"

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

	// search user on DB to check for exitance
	user, err := api.DB.GetUserByEmail(body.Email)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "user not registered")
	}

	// check given password if it is correct
	password := []byte(body.Password)
	passwordMatch := util.ComparePasswords(user.Password, password)
	if !passwordMatch {
		return fiber.NewError(fiber.ErrNotFound.Code, "user not registered")
	}

	// create token for user access
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	key := os.Getenv("JWT_KEY")
	value, err := token.SignedString(key)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "couldn't generate token")
	}

	result := struct {
		Token string `json:"token"`
	}{
		Token: value,
	}
	return ctx.JSON(result)
}
