package mapping_course

import (
)

type MappingCourseService interface {
}

type mappingCourseController struct {
	mappingCourseService MappingCourseService
}

func NewMappingCourseController(mappingCourseService MappingCourseService) mappingCourseController {
	return mappingCourseController{
		mappingCourseService: mappingCourseService,
	}
}