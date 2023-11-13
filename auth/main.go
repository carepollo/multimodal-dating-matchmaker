package main

import (
	"log"
	"net"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/handlers"
	"github.com/carepollo/multimodal-dating-matchmaker/auth/protos"
	"google.golang.org/grpc"
)

func main() {
	port := ":3000"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %v", port)
		return
	}

	instance := grpc.NewServer()
	protos.RegisterAuthServiceServer(instance, &handlers.AuthService{})
	log.Printf("server listening on port %v", port)

	if err := instance.Serve(listen); err != nil {
		log.Fatalf("failed to start service: %v", err)
	}
}
