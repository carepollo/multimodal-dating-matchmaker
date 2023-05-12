// the actions possible over a mongodb connection instace, using a custom wrapper to
// apply DRY principles over the handlers that consume this package
package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println("Connected to database successfully")
}

// disconnect from the database
func (db *Database) Disconnect() {
	db.client.Disconnect(db.context)
}

// generic implementation to make cleaner the insert process
func (db *Database) insert(collectionName string, body interface{}) error {
	collection := db.database.Collection(collectionName)
	_, err := collection.InsertOne(db.context, body)
	return err
}

// this method is to get data from any collection given, to use must cast the value into the desired
func get[T any](db *Database, collectionName string, query bson.M) ([]T, error) {
	result := []T{}
	collection := db.database.Collection(collectionName)
	cursor, err := collection.Find(db.context, query)

	if err != nil {
		return nil, err
	}

	for cursor.Next(db.context) {
		var record T
		err = cursor.Decode(&record)
		if err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// method to get a single value, search by id and returns nil if not found.
// Check for error when using it
func getById[T any](db *Database, collectionName string, id string) (T, error) {
	var result T
	collection := db.database.Collection(collectionName)
	query := bson.M{"_id": id}

	// if error arises return empty value of such type
	err := collection.FindOne(db.context, query).Decode(&result)
	if err != nil {
		var empty T
		return empty, err
	}

	return result, nil
}

func updateById(db *Database, collectionName string, id string) error {
	collection := db.database.Collection(collectionName)
	_, err := collection.UpdateByID(db.context, bson.M{"_id": id}, nil)
	return err
}

func deleteById(db *Database, collectionName string, id string) error {
	collection := db.database.Collection(collectionName)
	_, err := collection.DeleteOne(db.context, bson.M{"_id": id})
	return err
}
