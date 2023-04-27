package teacher

import (
	"context"

	"github.com/jmoiron/sqlx"

	teacherController "go-management-auth-school/controller/teacher"
	teacherEntity "go-management-auth-school/entity/teacher"
)

type teacherRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewTeacherRepo(dbMaster, dbSlave *sqlx.DB) *teacherRepo {
	return &teacherRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo teacherRepo) FindAll(ctx context.Context, params *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	return
}

func (repo teacherRepo) SelectAll(ctx context.Context, parameter *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	return
}

func (repo teacherRepo) FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error) {
	return
}
