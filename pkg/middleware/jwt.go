package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/context"
)

func AuthMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}
		claim, err := ValidateToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect Token")
			return
		}
		context.Set(r, "id", claim.ID)
		context.Set(r, "email", claim.Email)
		next(w, r)
	}
}

type JWTClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(
				"TODO: REPLACE THIS STRING WITH LEGIT SECRET CODE",
			), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, err
}
