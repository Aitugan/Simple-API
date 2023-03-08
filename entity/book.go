package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	Id              uuid.UUID `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Author          string    `json:"author,omitempty"`
	PublicationDate string    `json:"publicationDate,omitempty"`
}

type BookUpdate struct {
	Title           *string `json:"title,omitempty"`
	Author          *string `json:"author,omitempty"`
	PublicationDate *string `json:"publicationDate,omitempty"`
}
