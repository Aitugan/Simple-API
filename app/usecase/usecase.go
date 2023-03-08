package usecase

import (
	bookUC "github.com/Aitugan/Techt/app/usecase/book"
	userUC "github.com/Aitugan/Techt/app/usecase/user"
)

type Usecase struct {
	UserUsecase userUC.UseCase
	BookUsecase bookUC.UseCase
}

func Get() *Usecase {

	userUsecase := userUC.New()
	bookUsecase := bookUC.New()

	return &Usecase{
		UserUsecase: userUsecase,
		BookUsecase: bookUsecase,
	}
}
