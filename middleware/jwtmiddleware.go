package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func IsAut() gin.HandlerFunc {
	return checkJWT()
}

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})
		if len(bearerToken) == 2 {

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var tm time.Time
				switch iat := claims["exp"].(type) {
				case float64:
					tm = time.Unix(int64(iat), 0)
				case json.Number:
					v, _ := iat.Int64()
					tm = time.Unix(v, 0)
				}
				c.JSON(http.StatusOK, gin.H{
					"username":   claims["username"],
					"is_valid":   token.Valid,
					"expired_at": tm.Format("2006-01-02"),
				})
			} else {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"errors": err,
					"status": "Invalid Token",
				})
				return
			}
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": err,
				"status": "Authorization token not provided",
			})
			return
		}
	}

}
