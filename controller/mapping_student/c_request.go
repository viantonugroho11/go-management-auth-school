package mapping_student

import "go-management-auth-school/controller"

type MappingStudentParams struct {
	Indentity string `json:"indentity"`
	Type      string `json:"type"`
	ClassID   string `json:"class_id"`
	controller.DefaultParameter
}
