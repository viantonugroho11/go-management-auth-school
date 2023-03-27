package user

import (
	"go-management-auth-school/controller"
)


type UserParams struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	IdentityID string `json:"identity_id"`
	Permission string `json:"permission"`
	controller.DefaultParameter
}
