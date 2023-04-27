package parent

import (
	"context"
	"errors"

	parentController "go-management-auth-school/controller/parent"
	studentService "go-management-auth-school/controller/student"
	parentEntity "go-management-auth-school/entity/parent"

	"github.com/jmoiron/sqlx"
)

type ParentRepo interface {
	FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error)
	SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error)
	FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *parentEntity.Parent) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type parentService struct {
	parentRepo ParentRepo
	studentServices studentService.StudentService
}

func NewParentService(repo ParentRepo, studentServices studentService.StudentService) *parentService {
	return &parentService{
		parentRepo: repo,
		studentServices: studentServices,
	}
}

func (service parentService) FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	return
}

func (service parentService) SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	return
}

func (service parentService) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	return
}

//create parent
func (service parentService) Create(ctx context.Context, params *parentEntity.Parent) (err error) {
	// checkStudent
	checkStudent , err := service.studentServices.FindOne(ctx, &studentService.StudentParams{})
	if err != nil {
		return
	}
	if checkStudent.ID == "" {
		return errors.New("Student not found")
	}

	checkNikParent, err := service.parentRepo.FindOne(ctx, &parentController.ParentParams{
		NIK: params.NIK,
	})
	if err != nil {
		return
	}
	if checkNikParent.ID != "" {
		return errors.New("NIK already exist")
	}
	// create parent
	err = service.parentRepo.Create(ctx,nil, params)
	return
}
