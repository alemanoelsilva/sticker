package entity

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}
