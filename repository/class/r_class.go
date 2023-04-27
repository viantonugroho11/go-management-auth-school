package class

import (
	"context"

	"github.com/jmoiron/sqlx"

	classController "go-management-auth-school/controller/class"
	classEntity "go-management-auth-school/entity/class"
)

type classRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewClassRepo(dbMaster ,dbSlave *sqlx.DB) *classRepo {
	return &classRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo classRepo) FindAll(ctx context.Context, params *classController.ClassParams) (data []classEntity.Class, err error) {
	// build query here
	return
}

func (repo classRepo) SelectAll(ctx context.Context, parameter *classController.ClassParams) (data []classEntity.Class, err error) {
	// build query here
	return
}

func (repo classRepo) FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error) {
	// build query here
	return
}