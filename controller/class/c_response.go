package class

import (
	classEntity "go-management-auth-school/entity/class"
)

type ClassResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Major string `json:"major"`
	Level int `json:"level"`
}

func FromServices(res []classEntity.Class) (data []ClassResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res classEntity.Class) (data ClassResponse) {
	data = ClassResponse{
		ID:    res.ID,
		Name:  res.Name,
		Major: res.Major.Name,
		Level: res.Level,
	}
	return
}
