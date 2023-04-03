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

type UserRequset struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RePassword string `json:"rePassword" validate:"required,eqfield=Password"`
	IdentityID string `json:"identityId" validate:"required"`
	DeviceID string `json:"deviceId"`
}
