package verify_token

import (

	"context"

	verifyTokenEntity "go-management-auth-school/entity/verify_token"
)

type VerifyTokenService interface {
	FindOne(ctx context.Context, parameter *VerifyTokenParams) (data verifyTokenEntity.VerifyToken, err error)
	Create(ctx context.Context, parameter *verifyTokenEntity.VerifyToken) (err error)
	Delete(ctx context.Context, parameter *verifyTokenEntity.VerifyToken) (err error)
}

type verifyTokenController struct {
	verifyTokenServices VerifyTokenService
}

func NewVerifyTokenController(service VerifyTokenService) *verifyTokenController {
	return &verifyTokenController{
		verifyTokenServices: service,
	}
}