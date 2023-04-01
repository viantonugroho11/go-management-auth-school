package jwt

import (
	"github.com/dgrijalva/jwt-go"
	authEntity "go-management-auth-school/entity/auth"
)

func JwtGenerator(params authEntity.JwtCustomClaimsStudent, key string) string {
	//Generate Token JWT for auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": params,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}
