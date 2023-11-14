package auth

import (
	"os"

	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(30 * time.Minute)

	claims["authorized"] = true

	claims["username"] = username

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
