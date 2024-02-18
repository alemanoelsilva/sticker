package entity

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}
