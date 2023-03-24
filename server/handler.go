package server

import (
	"fmt"
	"go-management-auth-school/response"
	"net/http"
)

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.RespondError(w, http.StatusNotFound, fmt.Errorf("404 not found"))
	})
}

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.RespondError(w, http.StatusMethodNotAllowed, fmt.Errorf("405 method not allowed"))
	})
}