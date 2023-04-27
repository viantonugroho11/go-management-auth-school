package parent

import (
	"context"

	"github.com/jmoiron/sqlx"

	parentController "go-management-auth-school/controller/parent"
	parentEntity "go-management-auth-school/entity/parent"
)


type parentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewParentRepo(dbMaster ,dbSlave *sqlx.DB) *parentRepo {
	return &parentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo parentRepo) FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	// build query here
	return
}

func (repo parentRepo) SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	// build query here
	return
}

func (repo parentRepo) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	// build query here
	return
}