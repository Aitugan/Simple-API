package collection_usecase

import (
	"github.com/Aitugan/Techt/entity"
	"github.com/Aitugan/Techt/pkg/database"
	"github.com/Aitugan/Techt/pkg/helpers"
	"github.com/google/uuid"
)

type usecase struct {
}

func New() UseCase {
	return &usecase{}
}

func (uc *usecase) Register(user *entity.User) (string, error) {
	users := database.GetUsersDatabase()
	for _, userFromDB := range users {
		if userFromDB.Email == user.Email {
			return "", ErrUserExists
		}
	}

	user.Id = uuid.New()
	database.AddUserIntoDatabase(user)

	token, err := helpers.GenerateJWTToken(user)
	if err != nil {
		return "", ErrUnexpected
	}
	return token, nil
}

func (uc *usecase) Login(user *entity.User) (string, error) {
	users := database.GetUsersDatabase()
	for _, userFromDB := range users {
		if userFromDB.Email == user.Email {
			if userFromDB.Password != user.Password {
				return "", ErrUserPasswordIncorrect
			}
			token, err := helpers.GenerateJWTToken(userFromDB)
			if err != nil {
				return "", ErrUnexpected
			}
			return token, nil

		}
	}
	return "", ErrUserDoesntExist
}
