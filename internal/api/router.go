package api

import (
	"github.com/Aitugan/Techt/internal/api/handlers"
	"github.com/Aitugan/Techt/pkg/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	h := handlers.Get()

	router := mux.NewRouter()
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.UserHandler.Register).Methods("POST")
	user.HandleFunc("/login", h.UserHandler.Login).Methods("POST")

	return router, nil
}
