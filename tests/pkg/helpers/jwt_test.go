package helpers_test

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Aitugan/Techt/entity"
	"github.com/Aitugan/Techt/pkg/helpers"
)

func TestGenerateJWTToken(t assert.TestingT) {
	user := &entity.User{
		Id:    uuid.MustParse("ad0cf393-72c5-4825-bc24-8daf9cfa84ae"),
		Email: "test@example.com",
	}

	tokenString, err := helpers.GenerateJWTToken(user)
	assert.NoError(t, err, "There should be no error while generating JWT token")
	assert.NotEmpty(t, tokenString, "Generated JWT token should not be empty")

	claims := &helpers.Claims{}
	_, parseErr := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("TODO: REPLACE THIS STRING WITH LEGIT SECRET CODE"), nil
	})

	assert.NoError(t, parseErr, "There should be no error while parsing JWT token")
	assert.Equal(t, user.Id.String(), claims.ID, "The user ID in the JWT token should match the user ID passed to the function")
	assert.Equal(t, user.Email, claims.Email, "The user email in the JWT token should match the user email passed to the function")
	expiryTime := time.Now().Add(time.Hour * 8760).Unix()
	assert.Equal(t, expiryTime, claims.ExpiresAt, "The expiry time in the JWT token should be 1 year from now")
}
