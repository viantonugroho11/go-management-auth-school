package lesson

import (
	lessonEntity "go-management-auth-school/entity/lesson"
)

type LessonResponse struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func FromServices(res []lessonEntity.Lesson) (data []LessonResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res lessonEntity.Lesson) (data LessonResponse) {
	data = LessonResponse{
		ID:   res.ID,
		Name: res.Name,
		Type: res.Type,
	}
	return
}
