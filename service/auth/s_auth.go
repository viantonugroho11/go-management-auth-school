package auth

import (
	"context"
	"go-management-auth-school/config"
	authLoginRequest "go-management-auth-school/controller/auth"
	mappingCourseServices "go-management-auth-school/controller/mapping_course"
	mapStudent "go-management-auth-school/controller/mapping_student"
	studentServices "go-management-auth-school/controller/student"
	userController "go-management-auth-school/controller/user"

	//entity
	authEntity "go-management-auth-school/entity/auth"
	userEntity "go-management-auth-school/entity/user"

	// "go-management-auth-school/entity/class"
	jwthelper "go-management-auth-school/helper/jwt"
	userRepo "go-management-auth-school/service/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
}

type authService struct {
	userRepo userRepo.UserRepo
	config config.Config
	mapCourseService mappingCourseServices.MappingCourseService
	studentService studentServices.StudentService
	mapStudentService mapStudent.MappingStudentService
}

func NewAuthService(repo userRepo.UserRepo, config config.Config, 
	mapCourse mappingCourseServices.MappingCourseService, studentService studentServices.StudentService,
	mapStudent mapStudent.MappingStudentService) *authService {
	return &authService{
		userRepo: repo,
		config: config,
		mapCourseService: mapCourse,
		studentService: studentService,
		mapStudentService: mapStudent,
	}
}

func (service authService) Login(ctx context.Context, parameter *authLoginRequest.LoginRequest) (data authEntity.Auth, err error) {
	dataUser,err := service.userRepo.FindOne(ctx, &userController.UserParams{
		Username: parameter.Username,
	})
	if err != nil {
		return
	}
	if dataUser.ID == "" {
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(parameter.Password))
	if err != nil {
		return
	}

	// generate token
	dataStudent, err := service.studentService.FindOne(ctx, &studentServices.StudentParams{
		IdentityID: dataUser.IdentityID,
	})
	if err != nil {
		return
	}
	if dataStudent.ID == "" {
		return
	}

	sessionID := uuid.New().String()
	refreshTokenExpireTime := time.Now().Add(time.Hour * time.Duration(service.config.JwtAuth.JwtRefreshExpireTime))
	tokenExpireTime := time.Now().Add(time.Hour * time.Duration(service.config.JwtAuth.JwtExpireTime))

	jwtClaims := jwt.StandardClaims{
			ExpiresAt: tokenExpireTime.Unix(),
			Id: dataUser.IdentityID,
		}
	
		refreshJwtClaims := jwt.StandardClaims{
			ExpiresAt:refreshTokenExpireTime.Unix(),
			Id: dataUser.IdentityID,
		}

 token := jwthelper.JwtGenerator(jwtClaims , service.config.JwtAuth.JwtSecretKey)
 refreshToken := jwthelper.JwtGenerator(refreshJwtClaims , service.config.JwtAuth.JwtRefreshSecretKey)

 data = authEntity.Auth{
	Indentity: dataUser.IdentityID,
	IsActive: dataUser.Status,
	// Type: dataUser.,
	ExpiredAt: tokenExpireTime.Format("2006-01-02 15:04:05"),
	Token: token,
	RefreshExpiredAt: refreshTokenExpireTime.Format("2006-01-02 15:04:05"),
	RefreshToken: refreshToken,
	SessionToken: sessionID,
 }
	
	return
}

// func (service authService) RefreshToken(ctx context.Context, parameter *authLoginRequest.RefreshTokenRequest) (data authEntity.Auth, err error) {
// 	return
// }

// func (service authService) Logout(ctx context.Context, parameter *authLoginRequest.LogoutRequest) (err error) {
// 	return
// }

// func (service authService) ValidateToken(ctx context.Context, parameter *authLoginRequest.ValidateTokenRequest) (err error) {
// 	return
// }

// func (service authService) ValidateRefreshToken(ctx context.Context, parameter *authLoginRequest.ValidateRefreshTokenRequest) (err error) {
// 	return
// }

// func (service authService) ValidateSessionToken(ctx context.Context, parameter *authLoginRequest.ValidateSessionTokenRequest) (err error) {
// 	return
// }

func (service authService) RegisterStudent(ctx context.Context, input *userEntity.User) (data string, err error) {
	// check student
	_, err = service.studentService.FindOne(ctx, &studentServices.StudentParams{
		IdentityID: input.IdentityID,
	})
	if err != nil {
		return
	}

	// check username
	_, err = service.userRepo.FindOne(ctx, &userController.UserParams{
		Username: input.Username,
	})
	if err != nil {
		return
	}

	

	return
}