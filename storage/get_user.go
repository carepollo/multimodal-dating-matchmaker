package storage

import (
	"errors"

	"github.com/carepollo/multimodal-dating-matchmaker/models"
	"go.mongodb.org/mongo-driver/bson"
)

// searchs on database users with such email and returns only the first one found
// as the email is supposed to be a unique value in database
func (db *Database) GetUserByEmail(email string) (*models.User, error) {
	query := bson.M{
		"email": email,
	}

	// execute the query
	results, err := get[models.User](db, "users", query)
	if err != nil {
		return nil, err
	}

	// check if there are results
	if len(results) == 0 {
		return nil, errors.New("not found")
	}

	user := results[0]
	return &user, nil
}
