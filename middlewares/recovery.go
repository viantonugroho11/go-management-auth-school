package middlewares

import (
	"go-management-auth-school/response"
	"log"
	"net/http"
	"runtime/debug"
)


func Recovery() func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					log.Fatalf("%s", debug.Stack())
					response.RespondError(w, http.StatusInternalServerError, ErrUnknownError)
					return
				}
			}()
			nextHandler.ServeHTTP(w, r)
		})
	}
}