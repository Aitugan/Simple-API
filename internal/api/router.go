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

	book := router.PathPrefix("/book").Subrouter()

	book.HandleFunc("/my", middleware.AuthMW(h.BookHandler.GetUserBooks)).Methods("GET")
	book.HandleFunc("", middleware.AuthMW(h.BookHandler.GetAllBooks)).Methods("GET")
	book.HandleFunc("/{id}", middleware.AuthMW(h.BookHandler.GetBookById)).Methods("GET")
	book.HandleFunc("", middleware.AuthMW(h.BookHandler.CreateBook)).Methods("POST")
	book.HandleFunc("/{id}", middleware.AuthMW(h.BookHandler.UpdateBook)).Methods("PUT")
	book.HandleFunc("/{id}", middleware.AuthMW(h.BookHandler.DeleteBook)).Methods("DELETE")

	return router, nil
}
