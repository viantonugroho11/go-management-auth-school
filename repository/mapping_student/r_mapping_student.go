package mapping_student

import "github.com/jmoiron/sqlx"


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