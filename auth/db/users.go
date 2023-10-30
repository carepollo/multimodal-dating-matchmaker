package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

var Users *dgo.Dgraph = nil

func CreateNode(data map[string]interface{}) error {
	ctx := context.Background()

	// Create a new transaction
	txn := Users.NewTxn()
	defer txn.Discard(ctx)

	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Perform a mutation to create data
	mutation := &api.Mutation{
		SetJson: jsonString,
	}
	res, err := txn.Mutate(ctx, mutation)
	if err != nil {
		return err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	fmt.Println(res.String())
	fmt.Println("Node inserted successfully!")
	return nil
}

func ReadNode(client *dgo.Dgraph) error {
	ctx := context.Background()

	// Create a new transaction
	txn := client.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	// Query for data
	query := `{
		all(func: has(name)) {
			uid
			name
			age
		}
	}`
	res, err := txn.Query(ctx, query)
	if err != nil {
		return err
	}

	fmt.Println(string(res.Json))
	return nil
}
