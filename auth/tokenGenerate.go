package auth

import (
	"os"

	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("SECRET_KEY"))

func generateJWT(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(30 * time.Minute)

	claims["authorized"] = true

	claims["user"] = username

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
