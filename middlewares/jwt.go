package middlewares

import (
	"github.com/dgrijalva/jwt-go"
)




func JwtGenerator(username, firstname, lastname, key string) string {
    //Generate Token JWT for auth
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username":  username,
        "firstname": firstname,
        "lastname":  lastname,
    })

    tokenString, err := token.SignedString([]byte(key))
    if err != nil {
        return err.Error()
    }
    return tokenString
}