package mapping_student

import (
	"context"

	mapStudentController "go-management-auth-school/controller/mapping_student"
	mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

type MpStudentRepo interface {
	FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	FindOne(ctx context.Context, params *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error)
}

type mpStudentService struct {
	mpStudentRepo MpStudentRepo
}

func NewMappingStudentService(repo MpStudentRepo) *mpStudentService {
	return &mpStudentService{
		mpStudentRepo: repo,
	}
}

func (service mpStudentService) FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	return
}

func (service mpStudentService) SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	return
}

func (service mpStudentService) FindOne(ctx context.Context, params *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error) {
	return
}
