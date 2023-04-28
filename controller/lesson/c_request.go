package lesson

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
	lessonEntity "go-management-auth-school/entity/lesson"
)

type LessonParams struct {
	controller.DefaultParameter
}

type LessonRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (i *LessonRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *LessonRequest) ToService() (res *lessonEntity.Lesson) {
	res = &lessonEntity.Lesson{
		Name: i.Name,
		Type: i.Type,
	}
	return
}
