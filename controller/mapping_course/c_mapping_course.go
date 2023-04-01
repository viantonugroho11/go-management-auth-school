package mapping_course

import (
	"context"
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
)

type MappingCourseService interface {
	FindAll(ctx context.Context, params *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	SelectAll(ctx context.Context, parameter *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	FindOne(ctx context.Context, params *MappingCourseParams) (data mapCourseEntity.MappingCourse, err error)
}

type mappingCourseController struct {
	mappingCourseService MappingCourseService
}

func NewMappingCourseController(mappingCourseService MappingCourseService) mappingCourseController {
	return mappingCourseController{
		mappingCourseService: mappingCourseService,
	}
}