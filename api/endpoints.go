package api

import "github.com/carepollo/multimodal-dating-matchmaker/api/middlewares"

// add all of the endpoints and append the handlers to the fiber router
func RegisterEndpoints(app *API) {
	router := app.Router
	router.Use(middlewares.Logger)

	auth := router.Group("/auth")
	auth.Post("/login", app.loginWithEmail)
	auth.Post("/register", app.registerWithEmail)

	user := router.Group("/user")
	user.Use(middlewares.IsLogged).Get("/profile/:id", app.getUserData)
}
