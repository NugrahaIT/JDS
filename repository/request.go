package repository

import "github.com/dgrijalva/jwt-go"

type NewRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
