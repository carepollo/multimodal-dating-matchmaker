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
		ctx.SendStatus(fiber.ErrBadRequest.Code)
		return nil
	}

	// search user on DB to check for exitance
	user, err := api.DB.GetUserByEmail(body.Email)
	if err != nil {
		ctx.SendStatus(fiber.ErrNotFound.Code)
		return nil
	}

	// check given password if it is correct
	password := []byte(body.Password)
	passwordMatch := util.ComparePasswords(user.Password, password)
	if !passwordMatch {
		ctx.SendStatus(fiber.ErrNotFound.Code)
		return nil
	}

	// search user on
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	token.Claims.GetExpirationTime()
	value, err := token.SignedString(os.Getenv("JWT_KEY"))
	if err != nil {
		ctx.SendStatus(fiber.ErrInternalServerError.Code)
		return nil
	}

	ctx.JSON(struct {
		token string
	}{
		token: value,
	})
	return nil
}
