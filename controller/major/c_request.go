package major

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
	majorEntity "go-management-auth-school/entity/major"
)

type MajorParams struct {
	controller.DefaultParameter
}

type MajorRequest struct {
	Name string `json:"name"`
}

func (i *MajorRequest) Validate() error {
		err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *MajorRequest) ToEntity() (res *majorEntity.Major) {
	res = &majorEntity.Major{
		Name: i.Name,
	}
	return
}

