package student

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
	studentEntity "go-management-auth-school/entity/student"
)

type StudentParams struct {
	ID         string `json:"id"`
	Nik        string `json:"nik"`
	Nisn       string `json:"nisn"`
	Nis        string `json:"nis"`
	IdentityID string `json:"identity_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	JoinDate   string `json:"join_date"`
	controller.DefaultParameter
}

type StudentRequest struct {
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Email string `json:"Email"`
	Nisn string `json:"Nisn" validate:"required"`
	Nis string `json:"Nis"`
	Nik string `json:"Nik"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	DateOfBirth string `json:"DateOfBirth"`

	Phone string `json:"Phone"`
	Address string `json:"Address"`
	Gender string `json:"gender"`
	Religion string `json:"Religion"`
	Image string `json:"Image"`
	Status int `json:"Status"`

	IsActive int `json:"IsActive"`
	ProvinceID int `json:"ProvinceID"`
	CityID int `json:"CityID"`
	SubdistrictID int `json:"SubdistrictID"`

	WardID int `json:"WardID"`
	Rt int `json:"Rt"`
	Rw int `json:"Rw"`
	// PostalCode int `json:"PostalCode"`

	Height int `json:"height"`
	Weight int `json:"weight"`
	BloodType string `json:"BloodType"`
	Disability int `json:"Disability"`
	DisabilityInfo string `json:"DisabilityInfo"`

	JoinDate string `json:"JoinDate"`
	Details string `json:"Details"`
}

func (i *StudentRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (i *StudentRequest) ToEntity() (res *studentEntity.Student) {
	res = &studentEntity.Student{
		FirstName: i.FirstName,
		LastName: i.LastName,
		Email: i.Email,
		Nisn: i.Nisn,
		Nik: i.Nik,
		PlaceOfBirth: i.PlaceOfBirth,
		DateOfBirth: i.DateOfBirth,
		Phone: i.Phone,
		Address: i.Address,
	}
	return

}
