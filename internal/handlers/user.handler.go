package handlers

import (
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	e "go_social/database/entities"
	db "go_social/database/operation"
	"go_social/internal/common"
	"log"
	"net/http"
	"strings"
)

func GetAllUser(c *gin.Context) {
	var users []e.UserEntity
	err := db.FetchUser(&users)
	common.Err(err, "Can't fetch data")
	common.RenderTemplate(c, "user", gin.H{
		"users": users,
	})
}

func Login(c *gin.Context) {
	common.RenderTemplate(c, "login", gin.H{
		"title": "login",
	})
}

func UserLogin(c *gin.Context) {
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))

	user, err := db.FindOne(email)
	common.Err(err, "Can't find user")
	isCorrect := common.CheckPasswordHash(password, user.Password)
	if !isCorrect {
		log.Fatal("Password is not correct")
	}

	c.Redirect(http.StatusOK, "/")

}

func Signup(c *gin.Context) {
	common.RenderTemplate(c, "signup", gin.H{
		"title": "Sign up",
	})
}

func UserSignup(c *gin.Context) {
	//res := make(map[string]interface{})
	username := strings.TrimSpace(c.PostForm("username"))
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	confirmPassword := strings.TrimSpace(c.PostForm("confirmPassword"))
	if password != confirmPassword {
		common.Err(false, "Password not same with password confirm")
	}

	mailErr := checkmail.ValidateFormat(email)
	hashPassword, hashErr := common.HashPassword(password)
	common.Err(hashErr, "Can't hash the password")
	common.Err(mailErr, "Invalid email format!")

	newUser := &e.UserEntity{
		Username:        username,
		Email:           email,
		Password:        string(hashPassword),
		ConfirmPassword: confirmPassword,
	}
	_, err := db.CreateNewUser(newUser)
	common.Err(err, "Can't create new user")
	//common.Json(c, newUser)
	c.Redirect(http.StatusMovedPermanently, "/")
}
