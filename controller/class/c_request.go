package class

import "go-management-auth-school/controller"

type ClassParams struct {
	Name  string `json:"name"`
	Major int    `json:"major"`
	Level string `json:"level"`
	controller.DefaultParameter
}
