package handlers

import (
	"encoding/json"
	"net/http"

	bookUC "github.com/Aitugan/Techt/app/usecase/book"
	userUC "github.com/Aitugan/Techt/app/usecase/user"
	"github.com/Aitugan/Techt/entity"
	handlerIntf "github.com/Aitugan/Techt/internal/api/handler_interfaces"
)

type handlerUser struct {
	userUsecase userUC.UseCase
	bookUsecase bookUC.UseCase
}

func NewUserHandler(userUsecase userUC.UseCase) handlerIntf.RestUserHandler {
	return &handlerUser{userUsecase: userUsecase}
}

func (h *handlerUser) Register(w http.ResponseWriter, r *http.Request) {
	var userBody *entity.User
	json.NewDecoder(r.Body).Decode(&userBody)
	user, err := h.userUsecase.Register(userBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *handlerUser) Login(w http.ResponseWriter, r *http.Request) {
	var userBody *entity.User
	json.NewDecoder(r.Body).Decode(&userBody)
	user, err := h.userUsecase.Login(userBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Some error occured:" + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
