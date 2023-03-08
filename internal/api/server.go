package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

// NewServer creates a new Server instance
func NewServer() (*Server, error) {
	fmt.Println("Configuring server...")
	router, err := NewRouter()
	if err != nil {
		log.Panicln("Failed to start new router:", err)
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	return &Server{&server}, nil
}

// Start - starts server
func (srv *Server) Start() {
	fmt.Println("Starting server..")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Println("Listening on", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Println("Shutting down server.. Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Server gracefully stopped")
}
