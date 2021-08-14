package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	i "github.com/bonzayio/go-atc-server/internal"
)

// Starting the server
func main() {
	// This is the domain the server should accept connections for.
	// domains := []string{"bonzay.io", "www.bonzay.io"}
	handler := i.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server
	go func() {
		log.Printf("Listening on port %s", port)
		srv.ListenAndServe()
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
