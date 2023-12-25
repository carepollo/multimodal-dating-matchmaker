package handlers

import (
	"context"
	"log"

	"github.com/carepollo/multimodal-dating-matchmaker/protos"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	protos.UnimplementedAuthServiceServer

	Ctx        context.Context
	GraphDB    *neo4j.SessionWithContext
	DocumentDB *mongo.Client
}

// they did it with a fucking interface instead of a struct, therefore I cannot access it with pointers
func (s *AuthService) GetGraphDB() neo4j.SessionWithContext {
	return *s.GraphDB
}

func (s *AuthService) CloseGraphDB() {
	session := *s.GraphDB // can't use directly because is an interface
	if err := session.Close(s.Ctx); err != nil {
		log.Printf("could not close connection to neo4j: %v", err)
	}
}

func (s *AuthService) CloseDocumentDB() {
	if err := s.DocumentDB.Disconnect(s.Ctx); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}
}
