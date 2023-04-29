package teacher

import (
	"context"

	teacherController "go-management-auth-school/controller/teacher"
	teacherEntity "go-management-auth-school/entity/teacher"

	"github.com/jmoiron/sqlx"
)

type TeacherRepo interface {
	FindAll(ctx context.Context, params *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error)
	SelectAll(ctx context.Context, parameter *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error)
	FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *teacherEntity.Teacher) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type teacherService struct {
	teacherRepo TeacherRepo
}

func NewTeacherService(repo TeacherRepo) *teacherService {
	return &teacherService{
		teacherRepo: repo,
	}
}

func (service teacherService) FindAll(ctx context.Context, params *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	return
}

func (service teacherService) SelectAll(ctx context.Context, parameter *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	data ,err = service.teacherRepo.SelectAll(ctx, parameter)
	if err != nil {
		return data, err
	}
	return
}

func (service teacherService) FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error) {
	data ,err = service.teacherRepo.FindOne(ctx, params)
	if err != nil {
		return data, err
	}
	return
}

func (service teacherService) Create(ctx context.Context, params *teacherEntity.Teacher) (err error) {
	tx ,err := service.teacherRepo.CreateTx(ctx)
	if err != nil {
		return err
	}

	err = service.teacherRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return
}
