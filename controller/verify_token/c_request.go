package verify_token

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
	verifyTokenEntity "go-management-auth-school/entity/verify_token"
)


type VerifyTokenParams struct {
	Identity string `json:"identity"`
	Token string `json:"token"`
	controller.DefaultParameter
}

type VerifyTokenRequest struct {
	Identity string `json:"identity" validate:"required"`
	Token string `json:"token" validate:"required"`
	Expired string `json:"expired" validate:"required"`
}

func (u *VerifyTokenRequest) Validate() error {
	err := validator.New().Struct(u)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

//to Service
func (u *VerifyTokenRequest) ToService() *verifyTokenEntity.VerifyToken {
	return &verifyTokenEntity.VerifyToken{
		Identity: u.Identity,
		Token: u.Token,
		ExpiredAt: u.Expired,
	}
}


