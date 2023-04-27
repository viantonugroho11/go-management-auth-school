package teacher

import (
	
	"context"

	teacherController "go-management-auth-school/controller/teacher"
	teacherEntity "go-management-auth-school/entity/teacher"
)

type TeacherRepo interface {
	FindAll(ctx context.Context, params *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error)
	SelectAll(ctx context.Context, parameter *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error)
	FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error)
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
	return
}

func (service teacherService) FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error) {
	return
}