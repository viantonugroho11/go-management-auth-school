package mapping_course

import (
	"context"

	mapCourseController "go-management-auth-school/controller/mapping_course"
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
)




type MpCourseRepo interface {
	FindAll(ctx context.Context, params *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	SelectAll(ctx context.Context, parameter *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	FindOne(ctx context.Context, params *mapCourseController.MappingCourseParams) (data mapCourseEntity.MappingCourse, err error)
}


type mpCourseService struct {
	mpCourseRepo MpCourseRepo
}

func NewMappingCourseService(repo MpCourseRepo) *mpCourseService {
	return &mpCourseService{
		mpCourseRepo: repo,
	}
}

func (service mpCourseService) FindAll(ctx context.Context, params *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error) {
	//logic here
	return
}

func (service mpCourseService) SelectAll(ctx context.Context, parameter *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error) {
	//logic here
	return
}

func (service mpCourseService) FindOne(ctx context.Context, params *mapCourseController.MappingCourseParams) (data mapCourseEntity.MappingCourse, err error) {
	//logic here
	return
}
