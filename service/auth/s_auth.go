package auth

import (
	"context"
	"go-management-auth-school/config"
	authLoginRequest "go-management-auth-school/controller/auth"
	authEntity "go-management-auth-school/entity/auth"
	userController "go-management-auth-school/controller/user"
	studentServices "go-management-auth-school/controller/student"
	mappingCourseServices "go-management-auth-school/controller/mapping_course"
	mapStudent "go-management-auth-school/controller/mapping_student"
	jwthelper "go-management-auth-school/helper/jwt"
	userRepo "go-management-auth-school/service/user"

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

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(parameter.Password))
	if err != nil {
		return
	}

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

	dataMapCourse, err := service.mapCourseService.GetMappingCourseByStudentID(ctx, dataUser.ID)



	// service.mapCourseService.GetMappingCourseByStudentID(ctx, dataUser.ID)
	// service.



	jwthelper.JwtGenerator(authEntity.JwtCustomClaimsStudent{},service.config.JwtAuth.JwtSecretKey)
	
}