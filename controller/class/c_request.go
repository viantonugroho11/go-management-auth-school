package class

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"

	classEntity "go-management-auth-school/entity/class"
)

type ClassParams struct {
	Name  string `json:"name"`
	Major int    `json:"major"`
	Level string `json:"level"`
	controller.DefaultParameter
}

type ClassRequest struct {
	Name  string `json:"name"`
	Major int    `json:"major"`
	Level string `json:"level"`
}

func (i *ClassRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *ClassRequest) ToService() (res *classEntity.Class) {
	res = &classEntity.Class{
		Name:  i.Name,
		MajorID: i.Major,
		Level: i.Level,
	}
	return
}


