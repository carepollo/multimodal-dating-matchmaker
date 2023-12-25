package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/implementation"
	"github.com/carepollo/multimodal-dating-matchmaker/auth/models"
	"github.com/carepollo/multimodal-dating-matchmaker/protos"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var env = models.Environment{}

func setUpEnvironment() models.Environment {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	if err := vp.ReadInConfig(); err != nil {
		panic("could not load env vars: " + err.Error())
	}

	if err := vp.Unmarshal(&env); err != nil {
		panic("could not load env vars: " + err.Error())
	}

	return env
}

func createMongoClient(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(env.Datasources.MongoDB.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic("could not connect to mongodb instance: " + err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic("could not connect to mongodb instance: " + err.Error())
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func createNeo4jClient(ctx context.Context) (*neo4j.SessionWithContext, error) {
	uri := env.Datasources.Neo4j.Uri
	username := env.Datasources.Neo4j.Username
	password := env.Datasources.Neo4j.Password

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic("could not connect to neo4j instance: " + err.Error())
	}

	// Create a Neo4j session
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	fmt.Println("Connected to Neo4j!")
	return &session, nil
}

func main() {
	setUpEnvironment()

	port := "3000"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %v", port)
		return
	}

	ctx := context.TODO()

	mongoClient, err := createMongoClient(ctx)
	if err != nil {
		panic("couldn't connect to mongo database: " + err.Error())
	}

	neoClient, err := createNeo4jClient(ctx)
	if err != nil {
		panic("couldn't connect to critical dgraph database: " + err.Error())
	}

	// Defer closing the Neo4j session when the program finishes

	instance := grpc.NewServer()
	service := &implementation.AuthService{
		DocumentDB: mongoClient,
		GraphDB:    neoClient,
		Ctx:        ctx,
		Env:        env,
	}

	protos.RegisterAuthServiceServer(instance, service)
	log.Printf("server listening on port %v", port)

	// Use a signal channel to wait for program termination
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := instance.Serve(listen); err != nil {
			log.Fatalf("failed to start service: %v", err)
		}
	}()

	// Wait for program termination signal
	<-signalCh

	// Graceful shutdown
	log.Println("Shutting down server...")
	instance.GracefulStop()

	// Explicitly call the close method for Neo4j to ensure the session is closed
	service.CloseDocumentDB()
	service.CloseGraphDB()

	log.Println("Server gracefully stopped.")
}
