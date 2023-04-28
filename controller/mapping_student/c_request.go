package mapping_student

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"

	mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

type MappingStudentParams struct {
	Indentity string `json:"indentity"`
	Type      string `json:"type"`
	ClassID   string `json:"class_id"`
	controller.DefaultParameter
}

type MappingStudentRequest struct {
	Indentity string `json:"indentity"`
	Type      string `json:"type"`
	ClassID   string `json:"class_id"`
}

func (m *MappingStudentRequest) Validate() error {
	err := validator.New().Struct(m)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (m *MappingStudentRequest) ToService() (res *mapStudentEntity.MappingStudentReq) {
	res = &mapStudentEntity.MappingStudentReq{
		Indentity: m.Indentity,
		Type:      m.Type,
		ClassID:   m.ClassID,
	}
	return
}
