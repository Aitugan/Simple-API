package collection_usecase

import (
	"github.com/Aitugan/Techt/entity"
	"github.com/Aitugan/Techt/pkg/database"
	"github.com/google/uuid"
)

type usecase struct {
}

func New() UseCase {
	return &usecase{}
}

func (uc *usecase) GetAllBooks() ([]*entity.Book, error) {
	books := database.GetBooksDatabase()
	return books, nil
}

func (uc *usecase) GetBookById(id uuid.UUID) (*entity.Book, error) {
	books := database.GetBooksDatabase()
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}
	return nil, ErrNotFound
}

func (uc *usecase) CreateBook(book *entity.Book) error {
	book.Id = uuid.New()
	database.AddBookIntoDatabase(book)
	return nil
}

func (uc *usecase) UpdateBook(id uuid.UUID, book *entity.BookUpdate) error {
	books := database.GetBooksDatabase()
	for _, bookToUpdate := range books {
		if bookToUpdate.Id == id {
			if book.Author != nil {
				bookToUpdate.Author = *book.Author
			}
			if book.PublicationDate != nil {
				bookToUpdate.PublicationDate = *book.PublicationDate
			}
			if book.Title != nil {
				bookToUpdate.Title = *book.Title
			}
			database.RemoveBookFromDatabase(id)
			database.AddBookIntoDatabase(bookToUpdate)
			return nil
		}
	}
	return ErrNotFound
}

func (uc *usecase) DeleteBook(id uuid.UUID) error {
	return database.RemoveBookFromDatabase(id)
}

func (uc *usecase) GetUserBooks(userId uuid.UUID) ([]*entity.Book, error) {
	books := database.GetBooksDatabase()
	users := database.GetUsersDatabase()
	userBooks := []*entity.Book{}
	userName := ""
	for _, user := range users {
		if user.Id == userId {
			userName = user.Name
		}
	}

	for _, book := range books {
		if book.Author == userName {
			userBooks = append(userBooks, book)
		}
	}
	return userBooks, nil
}
