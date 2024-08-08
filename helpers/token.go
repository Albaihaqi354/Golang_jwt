package helpers

import (
	"fmt"
	"golang_jwt/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySignininKey = []byte("secretkey")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySignininKey)

	return ss, err
}

func ValidationToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySignininKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
