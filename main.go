package main

import (
	"jds/handler"
	"jds/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	v1 := router.Group("/v1")

	// Post Register
	v1.POST("/registrasi", handler.Registrasi)
	v1.POST("/login", handler.Login)
	v1.GET("/check", middleware.IsAut())

	router.Run(":8080")

}
