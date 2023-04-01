package student

import (
	"go-management-auth-school/controller"
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
