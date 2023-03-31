package class

import (
	majorEntity "go-management-auth-school/entity/major"
)

type Class struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Major majorEntity.Major `json:"major"`
	Level string `json:"level"`
}