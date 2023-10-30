package handlers

import (
	"log"
	"time"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/db"
	"github.com/gofiber/fiber/v2"
)

type data struct {
	User     map[string]interface{} `json:"user"`
	Settings map[string]interface{} `json:"settings"`
}

func addAuditData(data map[string]interface{}, collectionName string) map[string]interface{} {
	data["ip"] = "0.0.0.0"
	data["method"] = "InsertDocument"
	data["collection"] = collectionName
	data["action"] = "add"
	data["creationDate"] = time.Now()
	return data
}

// handle the logic to add data of user and user settings
func HandleCreateUser(c *fiber.Ctx) error {
	var payload = new(data)
	err := c.BodyParser(payload)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = db.CreateNode(payload.User); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = db.InsertDocument(payload.Settings); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = db.InsertAudit(addAuditData(payload.User, "users")); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = db.InsertAudit(addAuditData(payload.Settings, "settings")); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
