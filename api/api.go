// this package covers all of the stuff related to endpoints of the REST APi,
// its handlers.
package api

import (
	"context"

	"github.com/carepollo/multimodal-dating-matchmaker/storage"
	"github.com/gofiber/fiber/v2"
)

// an entity to group the data that is going to shared across the API
type API struct {
	DB      *storage.Database // connection instance of mongodb
	Context context.Context   // context required for the DB
	Router  *fiber.App        // the endpoints and web server
}

// create a new instance of API
func New() *API {
	return &API{
		Context: context.TODO(),
		Router:  fiber.New(),
		DB:      storage.NewMongoDB(),
	}
}

// add all of the endpoints and append the handlers to the fiber router
func RegisterEndpoints(app *API) {
	router := app.Router

	auth := router.Group("/auth")
	auth.Post("/login-email", app.LoginWithEmail)
	auth.Post("/login-google", app.LoginWithGoogle)
	auth.Post("/login-facebook", app.LoginWithFacebook)
	auth.Post("/register-email", app.RegisterWithEmail)
	auth.Post("/register-google", app.RegisterWithGoogle)
	auth.Post("/register-facebook", app.RegisterWithFacebook)

	// user := router.Group("/user")
}
