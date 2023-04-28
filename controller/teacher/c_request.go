package teacher

import "go-management-auth-school/controller"

type TeacherParams struct {
	IdentityID string `json:"identity_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`

	controller.DefaultParameter
}
