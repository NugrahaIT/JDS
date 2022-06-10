package handler

import (
	"fmt"
	"jds/repository"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Registrasi(c *gin.Context) {
	var paramdata repository.NewRequest

	err := c.ShouldBindJSON(&paramdata)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, Condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
			"status": "false",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": paramdata.Username,
		"role":     paramdata.Role,
		"password": genHurufAngka(6),
	})
}

func genHurufAngka(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
