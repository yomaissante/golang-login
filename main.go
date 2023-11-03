package main

import (
	"github.com/golang/golang-login/initial"

	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

func init() {
	initial.LoadEnvVar()
}