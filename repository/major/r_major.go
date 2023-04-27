package major

import (
	"context"

	"github.com/jmoiron/sqlx"

	majorController "go-management-auth-school/controller/major"
	majorEntity "go-management-auth-school/entity/major"
)


type majorRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMajorRepo(dbMaster ,dbSlave *sqlx.DB) *majorRepo {
	return &majorRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo majorRepo) FindAll(ctx context.Context, params *majorController.MajorParams) (data []majorEntity.Major, err error) {
	// build query here
	return
}

func (repo majorRepo) SelectAll(ctx context.Context, parameter *majorController.MajorParams) (data []majorEntity.Major, err error) {
	// build query here
	return
}

func (repo majorRepo) FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error) {
	// build query here
	return
}