package storage

import "github.com/carepollo/multimodal-dating-matchmaker/models"

// insert a user in users collection
func (db *Database) CreateUser(data models.User) error {
	return db.insert(data, "users")
}
