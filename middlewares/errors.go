package middlewares

import "errors"

var (
	ErrUnknownError          = errors.New("ERR_UNKNOWN")
	ErrUnauthorizedError     = errors.New("ERR_UNAUTHORIZED")
	ErrUserNotActivatedError = errors.New("ERR_USER_NOT_ACTIVATED")
	ErrTimeoutError          = errors.New("ERR_TIMEOUT")
)