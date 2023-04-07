package mapping_course

import "github.com/jmoiron/sqlx"

type mpCourseRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMpCourseRepo(dbMaster ,dbSlave *sqlx.DB) *mpCourseRepo {
	return &mpCourseRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}