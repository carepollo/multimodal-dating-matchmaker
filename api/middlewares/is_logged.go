package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsLogged(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return fiber.NewError(fiber.ErrBadRequest.Code, "request doesn't have authentication token")
	}

	result, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		key := os.Getenv("JWT_KEY")
		return []byte(key), nil
	})

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "couldn't parse given token")
	}
	if !result.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "the token expired or is invalid")
	}

	return ctx.Next()
}
