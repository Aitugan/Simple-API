package helpers

import (
	"time"

	"github.com/Aitugan/Techt/entity"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWTToken(user *entity.User) (string, error) {
	claims := &Claims{
		ID:    user.Id.String(),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8760).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(
		"TODO: REPLACE THIS STRING WITH LEGIT SECRET CODE",
	))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
