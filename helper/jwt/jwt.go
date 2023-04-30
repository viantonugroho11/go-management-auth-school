package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	// authEntity "go-management-auth-school/entity/auth"
)

func JwtGenerator(params jwt.StandardClaims, key string) string {
	//Generate Token JWT for auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, params)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func JwtParser(token string, key string) (jwt.MapClaims, error) {
	//Parse Token JWT for auth
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func JwtValidate(token string, key string) (bool, error) {
	//Validate Token JWT for auth
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return false, err
	}

	if tokenParsed.Valid {
		return true, nil
	}

	return false, nil
}

func JwtCheckExpiredAt(expired int64, timeAdd int) (int64, bool) {
	//Check Expired Token JWT for auth
	if expired < 0 {
		return 0, false
	}

	if expired < time.Now().Unix() {
		return expired, true
	}

	if expired < time.Now().Add(time.Minute * 5).Unix() {
		return time.Now().Add(time.Hour * time.Duration(timeAdd)).Unix(), false
	}


	return expired, false
}
