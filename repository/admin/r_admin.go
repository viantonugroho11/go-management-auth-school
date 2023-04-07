package admin

import "github.com/jmoiron/sqlx"


type adminRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewAdminRepo(dbMaster ,dbSlave *sqlx.DB) *adminRepo {
	return &adminRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}