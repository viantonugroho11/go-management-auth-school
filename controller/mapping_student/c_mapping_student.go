package mapping_student

import (
)

type MappingStudentService interface {
}

type mappingStudentController struct {
	mappingStudentService MappingStudentService
}

func NewMappingStudentController(mappingStudentService MappingStudentService) mappingStudentController {
	return mappingStudentController{
		mappingStudentService: mappingStudentService,
	}
}