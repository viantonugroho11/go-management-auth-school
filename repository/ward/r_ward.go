package ward

import "github.com/jmoiron/sqlx"


type wardRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewWardRepo(dbMaster ,dbSlave *sqlx.DB) *wardRepo {
	return &wardRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}