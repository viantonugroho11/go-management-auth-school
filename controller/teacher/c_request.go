package teacher

import "go-management-auth-school/controller"

type TeacherParams struct {
	IdentityID string `json:"identity_id"`
	NIK        string `json:"nik"`
	Religion   string `json:"religion"`
	Gender     string `json:"gender"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`

	controller.DefaultParameter
}


type TeacherRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email" validate:"required"`
	Nik           string `json:"nik" validate:"required"`
	PlaceOfBirth  string `json:"place_of_birth"`
	DateOfBirth   string `json:"date_of_birth"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Gender        int    `json:"gender"`
	Religion      string `json:"religion"`
	Image         string `json:"image"`
	ProvinceID    int    `json:"province_id"`
	CityID        int    `json:"city_id"`
	SubdistrictID int    `json:"subdistrict_id"`
	WardID        int    `json:"ward_id"`
	RT            int    `json:"rt"`
	RW            int    `json:"rw"`
}

