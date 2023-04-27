package user

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
)

type UserParams struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	IdentityID string `json:"identity_id"`
	Permission string `json:"permission"`
	controller.DefaultParameter
}

type UserRequset struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"rePassword" validate:"required,eqfield=Password"`
	IdentityID string `json:"identityId" validate:"required"`
	DeviceID   string `json:"deviceId"`
}

func (u *UserRequset) Validate() error {
	err := validator.New().Struct(u)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}
