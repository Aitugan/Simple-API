package middleware_test

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"

	"github.com/Aitugan/Techt/pkg/middleware"
)

func TestValidateToken(t assert.TestingT) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaim{
		ID:    "ad0cf393-72c5-4825-bc24-8daf9cfa84ae",
		Email: "test@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	})

	signedToken, err := token.SignedString([]byte("secret-code"))
	assert.NoError(t, err, "There should be no error while generating JWT token")

	// Test valid token
	claims, err := middleware.ValidateToken(signedToken)
	assert.NoError(t, err, "There should be no error while validating JWT token")
	assert.NotNil(t, claims, "The validated JWT token should not be nil")
	assert.Equal(t, "123", claims.ID, "The ID in the JWT token should match the ID passed to the function")
	assert.Equal(t, "test@example.com", claims.Email, "The email in the JWT token should match the email passed to the function")

	// Test expired token
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaim{
		ID:    "123",
		Email: "test@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour * 1).Unix(),
		},
	})
	signedExpiredToken, err := expiredToken.SignedString([]byte("secret-code"))
	assert.NoError(t, err, "There should be no error while generating JWT token")

	_, err = middleware.ValidateToken(signedExpiredToken)
	assert.Error(t, err, "There should be an error while validating an expired JWT token")
	assert.True(t, errors.Is(err, errors.New("token expired")), "The error returned while validating an expired JWT token should be 'ErrExpiredToken'")
}
