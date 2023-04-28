package mapping_student

import (
	"context"
	"errors"

	"go-management-auth-school/controller"
	classService "go-management-auth-school/controller/class"
	mapStudentController "go-management-auth-school/controller/mapping_student"
	studentService "go-management-auth-school/controller/student"
	teacherService "go-management-auth-school/controller/teacher"
	mapStudentEntity "go-management-auth-school/entity/mapping_student"

	"github.com/jmoiron/sqlx"
)

type MpStudentRepo interface {
	FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	FindOne(ctx context.Context, params *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *mapStudentEntity.MappingStudentReq) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type mpStudentService struct {
	mpStudentRepo   MpStudentRepo
	studentServices studentService.StudentService
	teacherServices teacherService.TeacherService
	classServices   classService.ClassService
}

func NewMappingStudentService(repo MpStudentRepo, studentServices studentService.StudentService,
	teacherServices teacherService.TeacherService, classServices classService.ClassService) *mpStudentService {
	return &mpStudentService{
		mpStudentRepo:   repo,
		studentServices: studentServices,
		teacherServices: teacherServices,
		classServices:   classServices,
	}
}

func (service mpStudentService) FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	return
}

func (service mpStudentService) SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	data, err = service.mpStudentRepo.SelectAll(ctx, parameter)
	if err != nil {
		return
	}
	return
}

func (service mpStudentService) FindOne(ctx context.Context, params *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error) {

	data, err = service.mpStudentRepo.FindOne(ctx, params)
	if err != nil {
		return
	}

	return
}

func (service mpStudentService) Create(ctx context.Context, params *mapStudentEntity.MappingStudentReq) (err error) {

	check, err := service.checkIdentity(params, ctx)
	if check {
		return err
	}
	// check class
	checkClass, err := service.classServices.FindOne(ctx, &classService.ClassParams{
		DefaultParameter: controller.DefaultParameter{
			ID: params.ClassID,
		},
	})
	if err != nil {
		return err
	}
	if checkClass.ID == 0 {
		return errors.New("class not found")
	}

	tx, err := service.mpStudentRepo.CreateTx(ctx)
	if err != nil {
		return err
	}
	// create mapping student
	err = service.mpStudentRepo.Create(ctx, tx, params)
	if err != nil {
		return err
	}
	err = tx.Commit()

	return
}

func (service mpStudentService) checkIdentity(params *mapStudentEntity.MappingStudentReq, ctx context.Context) (bool, error) {
	switch params.Type {
	case "0":
		checkIdentityStudent, err := service.studentServices.FindOne(ctx, &studentService.StudentParams{
			IdentityID: params.Indentity,
		})
		if err != nil {
			return true, err
		}
		if checkIdentityStudent.ID == "" {
			return true, errors.New("student not found")
		}
	case "1":
		checkIdentityTeacher, err := service.teacherServices.FindOne(ctx, &teacherService.TeacherParams{
			IdentityID: params.Indentity,
		})
		if err != nil {
			return true, err
		}
		if checkIdentityTeacher.ID == "" {
			return true, errors.New("teacher not found")
		}
	default:
		return true, errors.New("type not found")
	}
	return false, nil
}
