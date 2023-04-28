package lesson

import (
	"context"

	lessonController "go-management-auth-school/controller/lesson"
	lessonEntity "go-management-auth-school/entity/lesson"

	"github.com/jmoiron/sqlx"
)

type LessonRepo interface {
	FindAll(ctx context.Context, params *lessonController.LessonParams) (data []lessonEntity.Lesson, err error)
	SelectAll(ctx context.Context, parameter *lessonController.LessonParams) (data []lessonEntity.Lesson, err error)
	FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error)
	Create(ctx context.Context, tx *sqlx.Tx, params *lessonEntity.Lesson) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
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
	data, err = service.lessonRepo.SelectAll(ctx, parameter)
	if err != nil {
		return data, err
	}
	return
}

func (service lessonService) FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error) {
	data, err = service.lessonRepo.FindOne(ctx, params)
	if err != nil {
		return data, err
	}
	return
}

func (service lessonService) Create(ctx context.Context, params *lessonEntity.Lesson) (err error) {
	tx, err := service.lessonRepo.CreateTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = service.lessonRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return
}
