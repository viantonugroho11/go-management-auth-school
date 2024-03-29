package auth

import (
	"context"
	"errors"
	configurable "go-management-auth-school/config"
	authLoginRequest "go-management-auth-school/controller/auth"
	mappingCourseServices "go-management-auth-school/controller/mapping_course"
	mapStudent "go-management-auth-school/controller/mapping_student"
	studentServices "go-management-auth-school/controller/student"
	userController "go-management-auth-school/controller/user"
	"strings"

	//entity
	authEntity "go-management-auth-school/entity/auth"
	userEntity "go-management-auth-school/entity/user"

	//repo
	verifyTokenController "go-management-auth-school/controller/verify_token"
	verifyTokenEntity "go-management-auth-school/entity/verify_token"
	verifyTokenRepo "go-management-auth-school/service/verify_token"

	//controller
	// studentController "go-management-auth-school/controller/student"
	// teacherController "go-management-auth-school/controller/teacher"

	// "go-management-auth-school/entity/class"
	jwthelper "go-management-auth-school/helper/jwt"
	strHelper "go-management-auth-school/helper/str"
	timeHelper "go-management-auth-school/helper/time"

	// validasiHelper "go-management-auth-school/helper/validasi"
	userRepo "go-management-auth-school/service/user"
	"time"

	//helper
	helperStr "go-management-auth-school/helper/str"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
}

type AuthService struct {
	userRepo          userRepo.UserRepo
	config            configurable.Config
	mapCourseService  mappingCourseServices.MappingCourseService
	studentService    studentServices.StudentService
	mapStudentService mapStudent.MappingStudentService
	userService       userController.UserService
	verifyTokenRepo   verifyTokenRepo.VerifyTokenRepo
}

func NewAuthService(repo userRepo.UserRepo, config configurable.Config,
	mapCourseService mappingCourseServices.MappingCourseService, studentService studentServices.StudentService,
	mapStudentService mapStudent.MappingStudentService, userService userController.UserService,
	verifyTokenRepository verifyTokenRepo.VerifyTokenRepo) *AuthService {
	return &AuthService{
		userRepo:          repo,
		config:            config,
		mapCourseService:  mapCourseService,
		studentService:    studentService,
		mapStudentService: mapStudentService,
		userService:       userService,
		verifyTokenRepo:   verifyTokenRepository,
	}
}

func (service AuthService) Login(ctx context.Context, parameter *authLoginRequest.LoginRequest) (data authEntity.Auth, err error) {
	dataUser, err := service.userRepo.FindOne(ctx, &userController.UserParams{
		Username: parameter.Username,
	})

	if err != nil {
		return data, err
	}
	if dataUser.ID == "" {
		return data,errors.New("user not found")
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
	refreshTokenExpireTime, tokenExpireTime, token, refreshToken := service.generateTokenJwt(dataUser)

	// create tx
	tx, err := service.userRepo.CreateTx(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback()

	// update last login
	err = service.userRepo.UpdateLastLogin(ctx, tx, &userEntity.User{
		IdentityID: dataUser.IdentityID,
	})
	if err != nil {
		return
	}

	//check token exist
	verifyToken, err := service.verifyTokenRepo.FindOne(ctx, &verifyTokenController.VerifyTokenParams{
		Identity: dataUser.IdentityID,
	})
	if verifyToken.ID != "" {
		err = service.verifyTokenRepo.Delete(ctx, nil, &verifyTokenEntity.VerifyToken{
			Identity: dataUser.IdentityID,
		})
		if err != nil {
			return
		}
	}

	// insert verify token
	err = service.verifyTokenRepo.Create(ctx, nil, &verifyTokenEntity.VerifyToken{
		Identity:  dataUser.IdentityID,
		Token:     token,
		ExpiredAt: strHelper.ConvertTimeToString(tokenExpireTime),
	})

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	data = authEntity.Auth{
		Identity: dataUser.IdentityID,
		IsActive: dataUser.Status,
		// Type: dataUser.,
		ExpiredAt:        tokenExpireTime.Format(timeHelper.DATE_TIME_FORMAT),
		Token:            token,
		RefreshExpiredAt: refreshTokenExpireTime.Format(timeHelper.DATE_TIME_FORMAT),
		RefreshToken:     refreshToken,
		SessionToken:     sessionID,
	}

	return
}

func (service AuthService) generateTokenJwt(dataUser userEntity.User) (time.Time, time.Time, string, string) {
	refreshTokenExpireTime := time.Now().Add(time.Hour * time.Duration(service.config.JwtAuth.JwtRefreshExpireTime))
	tokenExpireTime := time.Now().Add(time.Hour * time.Duration(service.config.JwtAuth.JwtExpireTime))

	jwtClaims := jwt.StandardClaims{
		ExpiresAt: tokenExpireTime.Unix(),
		Id:        dataUser.IdentityID,
	}

	refreshJwtClaims := jwt.StandardClaims{
		ExpiresAt: refreshTokenExpireTime.Unix(),
		Id:        dataUser.IdentityID,
	}

	token := jwthelper.JwtGenerator(jwtClaims, service.config.JwtAuth.JwtSecretKey)
	refreshToken := jwthelper.JwtGenerator(refreshJwtClaims, service.config.JwtAuth.JwtRefreshSecretKey)
	return refreshTokenExpireTime, tokenExpireTime, token, refreshToken
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

func (service AuthService) RegisterStudent(ctx context.Context, input *userEntity.User) (data authEntity.Auth, err error) {
	// check student
	checkData, err := service.studentService.FindOne(ctx, &studentServices.StudentParams{
		IdentityID: input.IdentityID,
	})
	if err != nil {
		return
	}
	if checkData.ID == "" {
		return
	}
	input.Permission = 0

	// service user
	dataUser, err := service.userService.Create(ctx, input)
	if err != nil {
		return
	}
	sessionID := uuid.New().String()
	refreshTokenExpireTime, tokenExpireTime, token, refreshToken := service.generateTokenJwt(dataUser)

	data = authEntity.Auth{
		Identity: dataUser.IdentityID,
		IsActive: dataUser.Status,
		// Type: dataUser.,
		ExpiredAt:        tokenExpireTime.Format(timeHelper.DATE_TIME_FORMAT),
		Token:            token,
		RefreshExpiredAt: refreshTokenExpireTime.Format(timeHelper.DATE_TIME_FORMAT),
		RefreshToken:     refreshToken,
		SessionToken:     sessionID,
	}

	return
}

func (service AuthService) ValidateToken(ctx context.Context, token string) (data authEntity.AuthValidate, err error) {

	var check bool
	tokenBearer := strings.Split(token, " ")
	if len(tokenBearer) == 2 {
		token = tokenBearer[1]
	}

	// verify JWT token
	secretKey := []byte(service.config.JwtAuth.JwtSecretKey)
	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
		}
		return secretKey, nil
	})

	if err != nil {
		// fmt.Println("Error while parsing token: ", err)
		return
	}

	if !tokenParsed.Valid {
		// fmt.Println("Token is invalid")
		return
	}

	if claims, ok := tokenParsed.Claims.(*jwt.StandardClaims); ok && tokenParsed.Valid {
		// get data from JWT token
		id := claims.Id
		data.Identity = id
		data.ExpiredAt = claims.ExpiresAt
		claims.ExpiresAt, check = jwthelper.JwtCheckExpiredAt(claims.ExpiresAt, service.config.JwtAuth.JwtExpireTime)
		if check {
			return data, errors.New("token expired")
		}

		// check token in database
		verifyToken, _ := service.verifyTokenRepo.FindOne(ctx, &verifyTokenController.VerifyTokenParams{
			Identity: id,
			Token:    token,
		})
		if verifyToken.ID == "" {
			return data, errors.New("invalid token")
		}

		// get data from database
		data.User, err = service.userRepo.FindOne(ctx, &userController.UserParams{
			IdentityID: id,
		})
		permission := helperStr.IsPermission(data.User.Permission)

		switch permission {
		case "student":
			data.Student, err = service.studentService.FindOne(ctx, &studentServices.StudentParams{
				Nis: data.User.IdentityID,
			})
		default:
			err = errors.New("invalid permission")
		}

		// get mapping student
	} else {
		return data, errors.New("invalid token")
	}
	return data, err
}

// refresh token
func (service AuthService) RefreshToken(ctx context.Context, token string) (data authEntity.Auth, err error) {
return
}
