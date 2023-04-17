package mapping_student

import (
	"context"
	 "github.com/jmoiron/sqlx"

mapStudentController "go-management-auth-school/controller/mapping_student"
mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

type mpStudentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMpStudentRepo(dbMaster ,dbSlave *sqlx.DB) *mpStudentRepo {
	return &mpStudentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo mpStudentRepo) FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	// build query here
	return
}

func (repo mpStudentRepo) SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	// build query here
	return
}

func (repo mpStudentRepo) FindOne(ctx context.Context, params *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error) {
	// build query here
	return
}