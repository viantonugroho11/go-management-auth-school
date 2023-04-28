package parent

import (
	"fmt"

	"go-management-auth-school/controller"
	parentEntity "go-management-auth-school/entity/parent"

	"github.com/go-playground/validator"
)

type ParentParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	StudentID string `json:"student_id"`
	NIK       string `json:"nik"`
	Gender    int    `json:"gender"`
	Type      string `json:"type"`
	controller.DefaultParameter
}

type ParentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	WorkID    int    `json:"work_id"`
	Income    int    `json:"income"`
	Image     string `json:"image"`
	StudentID string `json:"student_id"`
}

func (i *ParentRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *ParentRequest) ToService() (res *parentEntity.Parent) {
	res = &parentEntity.Parent{
		FirstName: i.FirstName,
		LastName:  i.LastName,
		WorkID:    i.WorkID,
		Gender:    i.Gender,
		Phone:     i.Phone,
		Image:     i.Image,
		Type:      i.Type,
		Income:    i.Income,
		StudentID: i.StudentID,
	}
	return
}
