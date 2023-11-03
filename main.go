package main

import (
	"github.com/golang/golang-login/initial"

	"github.com/golang/golang-login/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnvVar()
	initial.ConnectDB()
	migrate.Migrate()
}

func main() {
	router := gin.Default()
	router.Run()
}
