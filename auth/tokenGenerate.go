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

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	
}



