package middlewares

import (
	"errors"
	"go-management-auth-school/config"
	"go-management-auth-school/response"
	"strings"

	// "time"

	verifyTokenController "go-management-auth-school/controller/verify_token"
	helperJwt "go-management-auth-school/helper/jwt"
	verifyTokenRepo "go-management-auth-school/service/verify_token"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	// verifyTokenEntity "go-management-auth-school/entity/verify_token"
	// config "go-management-auth-school/config"
)

// ValidateToken ...
func ValidateToken(config config.Config, verifyTokenRepo verifyTokenRepo.VerifyTokenRepo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var check bool
			// get token from header
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return response.RespondError(c, 401, ErrTokenRequired)
			}
			tokenBearer := strings.Split(token, " ")
			if len(tokenBearer) == 2 {
				token = tokenBearer[1]
			}
			secretKey := []byte(config.JwtAuth.JwtSecretKey)
			tokenParsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
				}
				return secretKey, nil
			})
			if err != nil {
				return response.RespondError(c, 401, err)
			}
			if !tokenParsed.Valid {
				return response.RespondError(c, 401, ErrTokenExpired)
			}
			if claims, ok := tokenParsed.Claims.(*jwt.StandardClaims); ok && tokenParsed.Valid {
				claims.ExpiresAt, check = helperJwt.JwtCheckExpiredAt(claims.ExpiresAt, config.JwtAuth.JwtExpireTime)
				if check {
					return response.RespondError(c, 401, ErrUnauthorizedError)
				}
				// get data from database

				// validate token in database
				verifyToken, _ := verifyTokenRepo.FindOne(c.Request().Context(), &verifyTokenController.VerifyTokenParams{
					Identity: claims.Id,
					Token:    token,
				})
				if verifyToken.ID == "" {
					return response.RespondError(c, 401, ErrUnauthorizedError)
				}

			}
			return next(c)
		}
	}
}
