package middlewares

import (
	// "go-management-auth-school/response"
	"log"
	// "net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)


func Recovery() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				err := recover()
				if err != nil {
					log.Fatalf("%s", debug.Stack())
					return
				}
			}()
			return next(c)
		}
	}
}

// func Recovery() func(nextHandler http.Handler) http.Handler {
// 	return func(nextHandler http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			defer func() {
// 				err := recover()
// 				if err != nil {
// 					log.Fatalf("%s", debug.Stack())
// 					response.RespondError(w, http.StatusInternalServerError, ErrUnknownError)
// 					return
// 				}
// 			}()
// 			nextHandler.ServeHTTP(w, r)
// 		})
// 	}
// }