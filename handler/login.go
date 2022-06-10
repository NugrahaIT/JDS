package handler

import (
	"fmt"
	"jds/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"nugraha": "df722a",
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var creds repository.Credentials
	err := c.ShouldBindJSON(&creds)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, Condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
			"status": "Bad Request",
		})
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": ok,
			"status": "Unauthorized",
		})
		return
	}

	expirationTime := time.Now().AddDate(0, 0, 7)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
			"status": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": creds.Username,
		"token":    tokenString,
	})

}
