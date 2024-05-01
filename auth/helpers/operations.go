package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/models"
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

func GetUserByEmailAndPassword(
	ctx context.Context,
	client neo4j.SessionWithContext,
	email string,
	password string,
) (models.User, error) {
	query := `
		MATCH (u:User {email: $email, password: $password})
		RETURN u
	`

	result, err := client.Run(ctx, query, map[string]interface{}{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return models.User{}, err
	}

	for result.Next(ctx) {
		record := result.Record()
		userNode := record.Keys
		fmt.Println(userNode)
	}

	if err = result.Err(); err != nil {
		return models.User{}, nil
	}

	return models.User{}, nil
}
