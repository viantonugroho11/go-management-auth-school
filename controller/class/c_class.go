package class


import (
)

type ClassService interface {
	FindOne(id int)
}

type classController struct {
	classService ClassService
}

func NewClassController(classService ClassService) classController {
	return classController{
		classService: classService,
	}
}