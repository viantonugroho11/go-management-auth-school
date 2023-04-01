package auth

import (
	"context"
	"go-management-auth-school/config"
	authLoginRequest "go-management-auth-school/controller/auth"
	authEntity "go-management-auth-school/entity/auth"
	userController "go-management-auth-school/controller/user"
	jwthelper "go-management-auth-school/helper/jwt"
	userRepo "go-management-auth-school/service/user"

	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
}

type authService struct {
	userRepo userRepo.UserRepo
	config config.Config
}

func NewAuthService(repo userRepo.UserRepo) *authService {
	return &authService{
		userRepo: repo,
	}
}

func (service authService) Login(ctx context.Context, parameter *authLoginRequest.LoginRequest) (data userEntity.User, err error) {
	dataUser,err := service.userRepo.FindOne(ctx, &userController.UserParams{
		Username: parameter.Username,
	})
	if err != nil {
		return
	}
	if dataUser.ID == "" {
		err = errors.New("Username not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(parameter.Password))
	if err != nil {
		return
	}

	

	jwthelper.JwtGenerator(authEntity.JwtCustomClaimsStudent{},service.config.JwtAuth.JwtSecretKey)
	
}