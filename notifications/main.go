package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/carepollo/multimodal-dating-matchmaker/notifications/implementation"
	"github.com/carepollo/multimodal-dating-matchmaker/protos"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	//loading env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := "50051"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	server := &implementation.NotificationsService{}

	protos.RegisterNotificationsServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("server listening on port %v", port)

	// graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start service: %v", err)
		}
	}()

	// Wait for program termination signal
	<-signalCh

	// Graceful shutdown
	log.Println("Shutting down server...")
	s.GracefulStop()
	log.Println("Server gracefully stopped.")
}
