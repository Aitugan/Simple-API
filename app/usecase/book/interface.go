package collection_usecase

import (
	"errors"

	"github.com/Aitugan/Techt/entity"
	"github.com/google/uuid"
)

var ErrUnexpected = errors.New("Internal error")
var ErrUnauthorized = errors.New("ErrUnauthorized")
var ErrNotFound = errors.New("ErrNotFound")
var ErrForbiddenSelfRequest = errors.New("Self request is forbidden")

type UseCase interface {
	GetAllBooks() ([]*entity.Book, error)
	GetBookById(id uuid.UUID) (*entity.Book, error)
	CreateBook(book *entity.Book) error
	UpdateBook(id uuid.UUID, book *entity.BookUpdate) error
	DeleteBook(id uuid.UUID) error
	GetUserBooks(userId uuid.UUID) ([]*entity.Book, error)
}
