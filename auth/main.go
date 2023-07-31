package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/handlers"
)

func main() {
	logger := log.New(os.Stdout, "auth", log.LstdFlags)
	mux := http.NewServeMux()

	ping := handlers.NewPing(logger)
	login := handlers.NewLogin(logger)

	mux.Handle("/", ping)
	mux.Handle("/login", login)

	server := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	sig := <-channel
	log.Println("Graceful shutdown:", sig)

	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeout)
}
