package main

import (
	"net/http"

	"github.com/golang/golang-login/initial"

	"github.com/golang/golang-login/model"

	"github.com/golang/golang-login/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnvVar()
	initial.ConnectDB()
	migrate.Migrate()
}

func registerUser(c *gin.Context) {
	var newUser model.Userdata

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	initial.DB.Create(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.POST("/register", registerUser)
	router.Run()
}
