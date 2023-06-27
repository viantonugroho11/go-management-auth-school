package mapping_course

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	mpCourseController "go-management-auth-school/controller/mapping_course"
	mpCourseEntity "go-management-auth-school/entity/mapping_course"
)

type mpCourseRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMpCourseRepo(dbMaster, dbSlave *sqlx.DB) *mpCourseRepo {
	return &mpCourseRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo mpCourseRepo) buildingParams(ctx context.Context, parameter *mpCourseController.MappingCourseParams) (conditionString string, conditionParam []interface{}) {

	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}
	if parameter.ClassID != "" {
		conditionString += " AND def.class_id = ?"
		conditionParam = append(conditionParam, parameter.ClassID)
	}

	return
}
func (repo mpCourseRepo) FindAll(ctx context.Context, params *mpCourseController.MappingCourseParams) (data []mpCourseEntity.MappingCourse, err error) {
	// build query here
	return
}

func (repo mpCourseRepo) SelectAll(ctx context.Context, parameter *mpCourseController.MappingCourseParams) (data []mpCourseEntity.MappingCourse, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := mpCourseEntity.SelectMapCourse + ` WHERE def.deleted_at IS NULL ` + whereStatment

	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows && err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		d := mpCourseEntity.MappingCourse{}
		err = d.ScanRows(rows, nil)
		if err != sql.ErrNoRows && err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	err = rows.Err()
	if err != sql.ErrNoRows && err != nil {
		return
	}

	return
}

func (repo mpCourseRepo) FindOne(ctx context.Context, params *mpCourseController.MappingCourseParams) (data mpCourseEntity.MappingCourse, err error) {

	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := mpCourseEntity.SelectMapCourse + ` WHERE def.deleted_at IS NULL ` + whereStatment

	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows && err != nil {
		return
	}
	return
}

func (repo mpCourseRepo) Create(ctx context.Context, tx *sqlx.Tx, params *mpCourseEntity.MappingCourseReq) (err error) {

	queries := InsertMapCourse
	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, queries,
		uuidRandom,
		params.ClassID,
		params.TeacherID,
		params.LessonID,
	)
	if err != nil {
		return err
	}
	return
}

func (repo mpCourseRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	return
}

func (repo mpCourseRepo) Update(ctx context.Context, tx *sqlx.Tx, params *mpCourseEntity.MappingCourseReq) (err error) {
	
	queries := `UPDATE ` + mpCourseEntity.Table + ` SET class_id = ?, teacher_id = ?, lesson_id = ? WHERE id = ?`
	if tx == nil {
		_, err = repo.DbMaster.ExecContext(ctx, queries,
			params.ClassID,
			params.TeacherID,
			params.LessonID,
			params.ID,
		)
		if err != nil {
			return err
		}
		return
	}
	_, err = tx.ExecContext(ctx, queries,
		params.ClassID,
		params.TeacherID,
		params.LessonID,
		params.ID,
	)
	if err != nil {
		return err
	}
	return
}

func (repo mpCourseRepo) Delete(ctx context.Context, tx *sqlx.Tx, params *mpCourseEntity.MappingCourseReq) (err error) {
	
	queries := `UPDATE ` + mpCourseEntity.Table + ` SET deleted_at = NOW() WHERE id = ?`
	if tx == nil {
		_, err = repo.DbMaster.ExecContext(ctx, queries,
			params.ID,
		)
		if err != nil {
			return err
		}
		return
	}
	_, err = tx.ExecContext(ctx, queries,
		params.ID,
	)
	if err != nil {
		return err
	}
	return
}
