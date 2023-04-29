package parent

import (
	"context"
	"errors"
	"fmt"

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
	parentRepo      ParentRepo
	studentServices studentService.StudentService
}

func NewParentService(repo ParentRepo, studentServices studentService.StudentService) *parentService {
	return &parentService{
		parentRepo:      repo,
		studentServices: studentServices,
	}
}

func (service parentService) FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	return
}

func (service parentService) SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	data, err = service.parentRepo.SelectAll(ctx, parameter)
	if err != nil {
		return
	}
	return
}

func (service parentService) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	data, err = service.parentRepo.FindOne(ctx, params)
	if err != nil {
		return
	}
	return
}

// create parent
func (service parentService) Create(ctx context.Context, params *parentEntity.Parent) (err error) {
	// checkStudent
	checkStudent, _ := service.studentServices.FindOne(ctx, &studentService.StudentParams{})
	if checkStudent.ID == "" {
		return errors.New("Student not found")
	}

	checkNikParent, _ := service.parentRepo.FindOne(ctx, &parentController.ParentParams{
		NIK: params.NIK,
	})
	fmt.Println("checkNikParent", checkNikParent)
	if checkNikParent.ID != "" {
		return errors.New("NIK already exist")
	}

	tx, err := service.parentRepo.CreateTx(ctx)
	if err != nil {
		return
	}
	// create parent
	err = service.parentRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
