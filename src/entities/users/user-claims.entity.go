package users_entity

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}
