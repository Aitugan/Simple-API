package handlers

import (
	"github.com/Aitugan/Techt/app/usecase"
	handlerIntf "github.com/Aitugan/Techt/internal/api/handler_interfaces"
)

type Handler struct {
	UserHandler handlerIntf.RestUserHandler
	BookHandler handlerIntf.RestBookHandler
}

func Get() *Handler {
	uc := usecase.Get()
	bookHandler := NewBookHandler(uc.UserUsecase, uc.BookUsecase)
	userHandler := NewUserHandler(uc.UserUsecase)

	return &Handler{
		BookHandler: bookHandler,
		UserHandler: userHandler,
	}
}
