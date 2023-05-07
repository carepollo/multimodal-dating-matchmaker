// this package covers all of the stuff related to endpoints of the REST APi,
// its handlers.
package api

import (
	"github.com/carepollo/multimodal-dating-matchmaker/storage"
	"github.com/carepollo/multimodal-dating-matchmaker/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// an entity to group the data that is going to shared across the API
type API struct {
	DB                *storage.Database // connection instance of mongodb
	Router            *fiber.App        // the endpoints and web server
	googleOauthConfig *oauth2.Config    // google object to authentication
	Cache             *storage.Cache    // cache database
}

// create a new instance of API
func New() *API {
	return &API{
		Router:            fiber.New(),
		DB:                storage.NewMongoDB(),
		googleOauthConfig: util.SetupGoogleConfig(),
		Cache:             storage.NewRedis(),
	}
}

// add all of the endpoints and append the handlers to the fiber router
func RegisterEndpoints(app *API) {
	router := app.Router

	auth := router.Group("/auth")
	auth.Post("/login-email", app.loginWithEmail)
	auth.Post("/login-google", app.loginWithGoogle)
	auth.Post("/login-facebook", app.loginWithFacebook)
	auth.Post("/register-email", app.registerWithEmail)
	auth.Post("/register-google", app.registerWithGoogle)
	auth.Post("/google-callback", app.callbackRegisterWithGoogle) // redirect for google authentication
	auth.Post("/register-facebook", app.registerWithFacebook)

	// user := router.Group("/user")
}
