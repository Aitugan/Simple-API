package api

import (
	"github.com/Aitugan/Techt/internal/api/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	h := handlers.Get()

	router := mux.NewRouter()
	return router, nil
}
