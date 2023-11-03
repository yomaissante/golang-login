package migrate

import (
	"github.com/golang/golang-login/initial"
	"github.com/golang/golang-login/model"
)

func Migrate() {
	initial.DB.AutoMigrate(&model.Userdata{})
}