package class

import "github.com/jmoiron/sqlx"

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