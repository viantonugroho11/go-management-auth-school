package lesson

import "github.com/jmoiron/sqlx"


type lessonRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewLessonRepo(dbMaster ,dbSlave *sqlx.DB) *lessonRepo {
	return &lessonRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}