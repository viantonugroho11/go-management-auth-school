package auth

import (
	"context"
	"go-management-auth-school/config"
	authLoginRequest "go-management-auth-school/controller/auth"
	mappingCourseServices "go-management-auth-school/controller/mapping_course"
	mapStudent "go-management-auth-school/controller/mapping_student"
	studentServices "go-management-auth-school/controller/student"
	userController "go-management-auth-school/controller/user"
	authEntity "go-management-auth-school/entity/auth"
	"go-management-auth-school/entity/class"
	jwthelper "go-management-auth-school/helper/jwt"
	userRepo "go-management-auth-school/service/user"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(parameter.Password))
	if err != nil {
		return
	}

	// generate token
	dataStudent, err := service.studentService.FindOne(ctx, &studentServices.StudentParams{
		IdentityID: dataUser.IdentifyID,
	})
	if err != nil {
		return
	}
	if dataStudent.ID == "" {
		err = errors.New("Student not found")
		return
	}

	dataMapStudent, err := service.mapStudentService.FindOne(ctx, &mapStudent.MappingStudentParams{
		Indentity: dataStudent.Nis,
	})
	if err != nil || dataMapStudent.ID == 0 {
		return
	}

	dataMapCourse, err := service.mapCourseService.FindAll(ctx, &mappingCourseServices.MappingCourseParams{
		ClassID: dataMapStudent.ClassID,
	})
	if err != nil {
		return
	}
	if len(dataMapCourse) == 0 {
		err = errors.New("Course not found")
		return
	}

	jwtClaims := authEntity.JwtCustomClaimsStudent{
		Username: dataUser.Username,
		Firstname: dataStudent.FirstName,
		Lastname: dataStudent.LastName,
		Indentity: dataUser.IdentifyID,
		Type: "student",
		Phone: dataStudent.Phone,
		Class: class.Class{
			// ID: dataMapStudent.ClassID,
		},
		MappingCourse: dataMapCourse,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(service.config.JwtAuth.JwtExpireTime)).Unix(),
			Id: dataUser.ID,
		},
	}

	jwthelper.JwtGenerator(jwtClaims , service.config.JwtAuth.JwtSecretKey)
	
	return
}