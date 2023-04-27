package lesson

import (
	"context"

	"github.com/jmoiron/sqlx"

	lessonController "go-management-auth-school/controller/lesson"
	lessonEntity "go-management-auth-school/entity/lesson"
)


type lessonRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewLessonRepo(dbMaster ,dbSlave *sqlx.DB) *lessonRepo {
	return &lessonRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo lessonRepo) FindAll(ctx context.Context, params *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	return
}

func (repo lessonRepo) SelectAll(ctx context.Context, parameter *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	return
}

func (repo lessonRepo) FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error) {
	return
}