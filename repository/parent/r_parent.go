package parent

import "github.com/jmoiron/sqlx"


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