package server

import (
	"fmt"
	"go-management-auth-school/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func NotFoundHandler() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		response.RespondError(w, http.StatusNotFound, fmt.Errorf("404 not found"))
// 	})
// }

// func NotFoundHandler() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return response.RespondError(c, http.StatusNotFound, fmt.Errorf("404 not found"))
// 	}
// }

func NotFoundHandler(c echo.Context) error {
	return response.RespondError(c, http.StatusNotFound, fmt.Errorf("404 not found"))
}

func MethodNotAllowedHandler(c echo.Context) error {
	return response.RespondError(c, http.StatusMethodNotAllowed, fmt.Errorf("405 method not allowed"))
}

// func MethodNotAllowedHandler() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		response.RespondError(w, http.StatusMethodNotAllowed, fmt.Errorf("405 method not allowed"))
// 	})
// }
