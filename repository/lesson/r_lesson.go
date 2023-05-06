package lesson

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	lessonController "go-management-auth-school/controller/lesson"
	lessonEntity "go-management-auth-school/entity/lesson"
)

type lessonRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewLessonRepo(dbMaster, dbSlave *sqlx.DB) *lessonRepo {
	return &lessonRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo lessonRepo) buildingParams(ctx context.Context, parameter *lessonController.LessonParams) (conditionString string, conditionParam []interface{}) {
	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}

	return
}

func (repo lessonRepo) FindAll(ctx context.Context, params *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	return
}

func (repo lessonRepo) SelectAll(ctx context.Context, parameter *lessonController.LessonParams) (data []lessonEntity.Lesson, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := lessonEntity.SelectStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment

	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d := lessonEntity.Lesson{}
		err = d.ScanRows(rows, nil)
		if err != sql.ErrNoRows {
			return nil, err
		}
		data = append(data, d)
	}

	return
}

func (repo lessonRepo) FindOne(ctx context.Context, params *lessonController.LessonParams) (data lessonEntity.Lesson, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := lessonEntity.SelectStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment

	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows {
		return data, err
	}
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows {
		return data, err
	}

	return
}

func (repo lessonRepo) Create(ctx context.Context, tx *sqlx.Tx, params *lessonEntity.Lesson) (err error) {
	queries := InsertLesson
	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, queries, uuidRandom, params.Name, params.Type)
	if err != nil {
		return err
	}
	return
}

func (repo lessonRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	return
}
