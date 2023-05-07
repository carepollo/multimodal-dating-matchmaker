package app

import (
	"os"

	"github.com/carepollo/multimodal-dating-matchmaker/api"

	"github.com/joho/godotenv"
)

var server = api.New()

func Run() {
	defer server.DB.Disconnect()

	//loading environment variables
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	//connecting to database
	server.DB.Connect(os.Getenv("MONGODB_CONNECTION"))

	//starting web server
	api.RegisterEndpoints(server)
	err = server.Router.Listen(":8080")
	if err != nil {
		panic(err.Error())
	}
}
