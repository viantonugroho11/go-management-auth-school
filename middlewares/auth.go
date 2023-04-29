package middlewares

import (
	"errors"
	"go-management-auth-school/config"
	"go-management-auth-school/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	// config "go-management-auth-school/config"
)

// ValidateToken ...
func ValidateToken(config config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get token from header
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return response.RespondError(c, 401, ErrTokenRequired)
			}
			secretKey := []byte(config.JwtAuth.JwtSecretKey)
			tokenParsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
				}
				return secretKey, nil
			})
			if err != nil {
				return response.RespondError(c, 401, ErrUnauthorizedError)
			}
			if !tokenParsed.Valid {
				return response.RespondError(c, 401, ErrUnauthorizedError)
			}
			if claims, ok := tokenParsed.Claims.(*jwt.StandardClaims); ok && tokenParsed.Valid {
				if claims.ExpiresAt < time.Now().Unix() {
					return response.RespondError(c, 401, ErrTokenExpired)
				}
				
				// get data from database

			}
			return next(c)
		}
	}
}
