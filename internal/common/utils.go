package common

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RenderTemplate(c *gin.Context, tmpl string, p interface{}) {
	c.HTML(http.StatusOK, tmpl+".html", p)
}

func Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Err(err interface{}, message ...string) {
	if err != nil {
		log.Fatal(message[0], err)
	}
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load environment variables")
	}
}

// HashPassword Hash the password
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return bytes, err
}

// CheckPasswordHash compare from the input password and the hashed password
func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
