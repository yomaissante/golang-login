package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/golang-login/initial"

	"github.com/golang/golang-login/model"

	"github.com/golang/golang-login/migrate"

	"github.com/gin-gonic/gin"

	"errors"
)

func init() {
	initial.LoadEnvVar()
	initial.ConnectDB()
	migrate.Migrate()
}

type Login struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

func getAccountInfo(username string) (*model.Userdata, error){
	var user model.Userdata

	result := initial.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, errors.New("username / password incorrect")
	}

	return &user, nil
}

func login(c *gin.Context) {
	var login Login

	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messsage": "username / password required."})
	}

	username := login.Username
	
	accountInfo, err := getAccountInfo(username)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"messsage": "username / password incorrect."})
		return
	}

	accError := bcrypt.CompareHashAndPassword([]byte(accountInfo.Password), []byte(login.Password))

	if accError != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"messsage": "username / password incorrect."})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"messsage": "Welcome!."})
}

func registerUser(c *gin.Context) {
	var newUser model.Userdata

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	passBeforeHash := []byte(newUser.Password)

	hashedPass, err := bcrypt.GenerateFromPassword(passBeforeHash, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

	newUser.Password = string(hashedPass)

	initial.DB.Create(newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{"messsage": "User Created."})
}

func main() {
	router := gin.Default()
	router.POST("/login", login)
	router.POST("/register", registerUser)
	router.Run()
}
