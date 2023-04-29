package student

import (
	"fmt"
	parentEntity "go-management-auth-school/entity/parent"
	studentEntity "go-management-auth-school/entity/student"
	helperStr "go-management-auth-school/helper/str"
)

type StudentResponse struct {
	ID             string           `json:"id"`
	Nik            string           `json:"nik"`
	Nisn           string           `json:"nisn"`
	FullName       string           `json:"full_name"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Email          string           `json:"email"`
	PlaceOfBirth   string           `json:"place_of_birth"`
	DateOfBirth    string           `json:"date_of_birth"`
	Phone          string           `json:"phone"`
	Gender         string           `json:"gender"`
	Religion       string           `json:"religion"`
	Image          string           `json:"image"`
	Status         string           `json:"status"`
	IsActive       string           `json:"is_active"`
	Address        string           `json:"address"`
	Height         string           `json:"height"`
	Weight         string           `json:"weight"`
	BloodType      string           `json:"blood_type"`
	Disability     string           `json:"disability"`
	DisabilityInfo string           `json:"disability_info"`
	Details        string           `json:"details"`
	RT             string           `json:"rt"`
	RW             string           `json:"rw"`
	Parent         []ParentResponse `json:"parent"`
}

type ParentResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	WorkID   int    `json:"work_id"`
	WorkName string `json:"work_name"`
	Income   int    `json:"income"`
	Image    string `json:"image"`
}

func FromServices(res []studentEntity.Student) (data []StudentResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res studentEntity.Student) (data StudentResponse) {
	data = StudentResponse{
		ID:             res.ID,
		Nik:            res.Nik,
		Nisn:           res.Nisn,
		FullName:       helperStr.GetFullNameStudent(res),
		FirstName:      res.FirstName,
		LastName:       res.LastName,
		Email:          res.Email,
		PlaceOfBirth:   res.PlaceOfBirth,
		DateOfBirth:    res.DateOfBirth,
		Phone:          res.Phone,
		Gender:         res.Gender,
		Religion:       res.Religion,
		Image:          res.Image,
		Status:         helperStr.IsTypeStudent(res.Status),
		IsActive:       helperStr.IsActiveConvert(res.IsActive),
		Address:        res.Address,
		Height:         helperStr.IntToString(res.Height),
		Weight:         helperStr.IntToString(res.Weight),
		BloodType:      res.BloodType,
		Disability:     helperStr.IsTypeStudent(res.Disability),
		DisabilityInfo: res.DisabilityInfo,
		Details:        fmt.Sprintf("%v", res.Details),
		RT:             helperStr.IntToString(res.Rt),
		RW:             helperStr.IntToString(res.Rw),
		Parent:         FromServicesParent(res.Parent),
	}
	return
}

func FromServicesParent(res []parentEntity.Parent) (data []ParentResponse) {
	for _, v := range res {
		data = append(data, FromServiceParent(v))
	}
	return
}

func FromServiceParent(res parentEntity.Parent) (data ParentResponse) {
	data = ParentResponse{
		ID:       res.ID,
		Name:     helperStr.GetFullNameParent(res),
		Type:     helperStr.IsTypeParent(helperStr.StringToInt(res.Type)),
		Gender:   helperStr.IsGender(res.Gender),
		Phone:    res.Phone,
		WorkID:   res.WorkID,
		WorkName: res.WorkName,
		Income:   res.Income,
		Image:    res.Image,
	}
	return
}
