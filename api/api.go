// this package covers all of the stuff related to endpoints of the REST APi,
// its handlers.
package api

import (
	"github.com/carepollo/multimodal-dating-matchmaker/storage"
	"github.com/gofiber/fiber/v2"
)

// an entity to group the data that is going to shared across the API
type API struct {
	DB     *storage.Database // connection instance of mongodb
	Router *fiber.App        // the endpoints and web server
	Cache  *storage.Cache    // cache database
}

// create a new instance of API
func New() *API {
	return &API{
		Router: fiber.New(),
		DB:     storage.NewMongoDB(),
		Cache:  storage.NewRedis(),
	}
}
