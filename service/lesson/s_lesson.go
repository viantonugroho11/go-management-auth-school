package lesson

import (
	"context"

	lessonController "go-management-auth-school/controller/lesson"
	lessonEntity "go-management-auth-school/entity/lesson"
)

type LessonRepo interface {
	FindAll(ctx context.Context, params *lessonController.LessonParams) (data []lessonEntity.Lesson, err error)
	SelectAll(ctx context.Context, parameter *lessonController.LessonParams) (data []lessonEntity.Lesson, err error)
	FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error)
}

type lessonService struct {
	lessonRepo LessonRepo
}

func NewLessonService(repo LessonRepo) *lessonService {
	return &lessonService{
		lessonRepo: repo,
	}
}

func (service lessonService) FindAll(ctx context.Context, params *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	return
}

func (service lessonService) SelectAll(ctx context.Context, parameter *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	return
}

func (service lessonService) FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error) {
	return
}