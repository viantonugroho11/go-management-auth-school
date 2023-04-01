package mapping_student

import (
	"context"

	mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

type MappingStudentService interface {
	FindAll(ctx context.Context, params *MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	SelectAll(ctx context.Context, parameter *MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	FindOne(ctx context.Context, params *MappingStudentParams) (data mapStudentEntity.MappingStudent, err error)
}

type mappingStudentController struct {
	mappingStudentService MappingStudentService
}

func NewMappingStudentController(mappingStudentService MappingStudentService) mappingStudentController {
	return mappingStudentController{
		mappingStudentService: mappingStudentService,
	}
}