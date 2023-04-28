package api

import (
	"github.com/carepollo/multimodal-dating-matchmaker/controllers"
	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/login-email", controllers.LoginWithEmail)
	auth.Post("/login-google", controllers.LoginWithGoogle)
	auth.Post("/login-facebook", controllers.LoginWithFacebook)

	auth.Post("/register-email", controllers.RegisterWithEmail)
	auth.Post("/register-google", controllers.RegisterWithGoogle)
	auth.Post("/register-facebook", controllers.RegisterWithFacebook)
}
