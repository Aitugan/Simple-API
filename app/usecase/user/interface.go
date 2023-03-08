package collection_usecase

import (
	"errors"

	"github.com/Aitugan/Techt/entity"
)

var ErrUnexpected = errors.New("Internal error")
var ErrUnauthorized = errors.New("ErrUnauthorized")
var ErrNotFound = errors.New("ErrNotFound")
var ErrUserExists = errors.New("ErrUserExists")
var ErrUserDoesntExist = errors.New("ErrUserDoesntExist")
var ErrUserPasswordIncorrect = errors.New("ErrUserPasswordIncorrect")

type UseCase interface {
	Register(user *entity.User) (string, error)
	Login(user *entity.User) (string, error)
}
