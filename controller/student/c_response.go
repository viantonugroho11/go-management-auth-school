package student


import (
	studentEntity "go-management-auth-school/entity/student"
)

type StudentResponse struct {
	ID string `json:"id"`
	Nik string `json:"nik"`
	Nisn string `json:"nisn"`
}

func FromServices(data []studentEntity.Student) (response []StudentResponse) {
	for _, v := range data {
		response = append(response, StudentResponse{
			ID: v.ID,
			Nik: v.Nik,
			Nisn: v.Nisn,
		})
	}
	return
}

