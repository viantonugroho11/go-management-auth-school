package major

import "github.com/jmoiron/sqlx"


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