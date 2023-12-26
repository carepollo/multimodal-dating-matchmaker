package helpers

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateUser(ctx context.Context, client neo4j.SessionWithContext, data map[string]any) error {
	registered, err := client.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `CREATE (:User {
			email: $email,
			password: $password,
			age: $age,
			status: $status,
			_id: $_id
		  });`, data)

		if err != nil {
			return nil, err
		}

		// successful
		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	log.Printf("user created successfully: %v\n", registered)
	return nil
}
