package auth

import (
	"context"
	userRepo "go-management-auth-school/service/user"
	authLoginRequest "go-management-auth-school/controller/auth"
	helper "go-management-auth-school/helper"
)

type authRepository interface {
}

type authService struct {
	userRepo userRepo.UserRepo
}

func NewAuthService(repo userRepo.UserRepo) *authService {
	return &authService{
		userRepo: repo,
	}
}

func (service authService) Login(ctx context.Context, parameter *authLoginRequest.LoginRequest) (data userEntity.User, err error) {
	
}