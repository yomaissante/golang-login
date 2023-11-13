package auth

import (
	"fmt"

	"os"

	"strconv"

	"strings"

	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secret =  []byte(os.Getenv("SECRET_KEY"))

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(30 * time.Minute)

	claims["authorized"] = true

	claims["user"] = "username"

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	 }
	
	 return tokenString, nil
}



