package auth

import (
	"fmt"
	"github.com/go-playground/validator"
	userEntity "go-management-auth-school/entity/user"
)


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RePassword string `json:"rePassword" validate:"required,eqfield=Password"`
	IdentityID string `json:"identityId" validate:"required"`
	DeviceID string `json:"deviceId"`
}

func (i *RegisterRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *RegisterRequest) ToEntity() (res *userEntity.User) {
	res = &userEntity.User{
		Username: i.Username,
		Password: i.Password,
		IdentityID: i.IdentityID,
		DeviceID: i.DeviceID,
	}
	return res
}
