package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	bookUC "github.com/Aitugan/Techt/app/usecase/book"
	userUC "github.com/Aitugan/Techt/app/usecase/user"
	"github.com/Aitugan/Techt/entity"
	handlerIntf "github.com/Aitugan/Techt/internal/api/handler_interfaces"
	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	userUsecase userUC.UseCase
	bookUsecase bookUC.UseCase
}

func NewBookHandler(userUsecase userUC.UseCase, bookUsecase bookUC.UseCase) handlerIntf.RestBookHandler {
	return &handlerBook{
		userUsecase: userUsecase,
		bookUsecase: bookUsecase,
	}
}

func (h *handlerBook) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookUsecase.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *handlerBook) GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect Id")
		return
	}
	book, err := h.bookUsecase.GetBookById(bookId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *handlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book *entity.Book
	json.NewDecoder(r.Body).Decode(&book)
	err := h.bookUsecase.CreateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *handlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect Id")
		return
	}

	var book *entity.BookUpdate
	json.NewDecoder(r.Body).Decode(&book)
	err = h.bookUsecase.UpdateBook(bookId, book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *handlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect Id")
		return
	}

	err = h.bookUsecase.DeleteBook(bookId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func (h *handlerBook) GetUserBooks(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(fmt.Sprintf("%v", context.Get(r, "id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect Id1")
		return
	}

	book, err := h.bookUsecase.GetUserBooks(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
