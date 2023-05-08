package storage

import (
	"errors"

	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *Database) GetUserByEmail(email string) (*models.User, error) {
	query := bson.M{
		"email": email,
	}
	results, err := db.get("users", query)
	if err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return nil, errors.New("not found")
	}

	user, ok := results[0].(models.User)
	if !ok {
		return nil, err
	}

	return &user, nil
}
