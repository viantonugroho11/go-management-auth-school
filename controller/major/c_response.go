package major

import (
	majorEntity "go-management-auth-school/entity/major"
)

type MajorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromServices(res []majorEntity.Major) (data []MajorResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res majorEntity.Major) (data MajorResponse) {
	data = MajorResponse{
		ID:   res.ID,
		Name: res.Name,
	}
	return
}
