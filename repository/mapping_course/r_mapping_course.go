package mapping_course

import (
	"context"

	"github.com/jmoiron/sqlx"

	mpCourseController "go-management-auth-school/controller/mapping_course"
	mpCourseEntity "go-management-auth-school/entity/mapping_course"
)

type mpCourseRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMpCourseRepo(dbMaster ,dbSlave *sqlx.DB) *mpCourseRepo {
	return &mpCourseRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo mpCourseRepo) FindAll(ctx context.Context, params *mpCourseController.MappingCourseParams) (data []mpCourseEntity.MappingCourse, err error) {
	// build query here
	return
}

func (repo mpCourseRepo) SelectAll(ctx context.Context, parameter *mpCourseController.MappingCourseParams) (data []mpCourseEntity.MappingCourse, err error) {
	// build query here
	return
}

func (repo mpCourseRepo) FindOne(ctx context.Context, params *mpCourseController.MappingCourseParams) (data mpCourseEntity.MappingCourse, err error) {
	// build query here
	return
}