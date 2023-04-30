// the actions possible over a mongodb connection instace, using a custom wrapper to
// apply DRY principles over the handlers that consume this package
package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// entity to represent a mongodb connection that is easier to use covering all the common actions to perform
type Database struct {
	client   *mongo.Client   // the mongodb connection client
	database *mongo.Database // the database where is the data of the app is stored
	context  context.Context // context for the db actions
}

// create a new instance of database
func NewMongoDB() *Database {
	return &Database{
		context: context.TODO(),
	}
}

// connect to mongodb atlas using MongoDB atlas cluster instance
func (db *Database) Connect(connectionString string) {
	var err error
	connectOptions := options.Client().ApplyURI(connectionString)
	db.client, err = mongo.Connect(
		db.context,
		connectOptions,
	)
	if err != nil {
		panic(err.Error())
	}

	db.database = db.client.Database("multimodal-dating-matchmaker")
	fmt.Println("Connected to database succesfully")
}

// disconnect from the database
func (db *Database) Disconnect() {
	db.client.Disconnect(db.context)
}

// generic implementation to make cleaner the insert process
func (db *Database) insert(body interface{}, collectionTarget string) error {
	collection := db.database.Collection(collectionTarget)
	_, err := collection.InsertOne(db.context, body)
	if err != nil {
		return err
	}
	return nil
}

// func (db *Database) get()
// func (db *Database) getById()
// func (db *Database) delete() error
