package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/golang-login/initial"

	"github.com/google/uuid"

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

func getAccountInfo(username string) (*model.Userdata, error){
	var user model.Userdata

	result := initial.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, errors.New("username / password incorrect")
	}

	return &user, nil
}

func validatePassword(loginPassword string, accountPassword string) (error){
	err := bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(loginPassword))

	if err != nil {
		return err
	}

	return nil
}

func login(c *gin.Context) {
	var login model.Login

	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messsage": "username / password required."})
	}

	username := login.Username
	
	accountInfo, err := getAccountInfo(username)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"messsage": "username / password incorrect."})
		return
	}

	accError := validatePassword(login.Password, accountInfo.Password)

	if accError != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"messsage": "username / password incorrect."})
		return
	}

	accountInfo.Is_login = true

	initial.DB.Model(&accountInfo).Where("username", login.Username).Update("is_login", accountInfo.Is_login)

	c.IndentedJSON(http.StatusAccepted, gin.H{"messsage": "Welcome!."})
}

func logout(c *gin.Context) {
	
}

func changePassword(c *gin.Context) {
	
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

	newUser.UserID = uuid.New().String()

	newUser.Is_login = false

	newUser.Password = string(hashedPass)

	initial.DB.Create(newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{"messsage": "User Created."})
}

func main() {
	router := gin.Default()
	router.POST("/logout", logout)
	router.POST("/login", login)
	router.POST("/change-password", changePassword)
	router.POST("/register", registerUser)
	router.Run()
}
