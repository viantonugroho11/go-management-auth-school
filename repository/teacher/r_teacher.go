package teacher

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	teacherController "go-management-auth-school/controller/teacher"
	teacherEntity "go-management-auth-school/entity/teacher"
)

type teacherRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewTeacherRepo(dbMaster, dbSlave *sqlx.DB) *teacherRepo {
	return &teacherRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo teacherRepo) buildingParams(ctx context.Context, parameter *teacherController.TeacherParams) (conditionString string, conditionParam []interface{}) {

	if parameter.FirstName != "" {
		conditionString += " AND def.first_name = ?"
		conditionParam = append(conditionParam, parameter.FirstName)
	}
	if parameter.LastName != "" {
		conditionString += " AND def.last_name = ?"
		conditionParam = append(conditionParam, parameter.LastName)
	}
	if parameter.Gender != "" {
		conditionString += " AND def.gender = ?"
		conditionParam = append(conditionParam, parameter.Gender)
	}

	return
}

func (repo teacherRepo) FindAll(ctx context.Context, params *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	// whereStatment, conditionParam := repo.buildingParams(ctx, params)
	// query := teacherEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + teacherEntity.GroupStatement +
	// 	` ORDER BY def.id` + params.OrderBy

	// // query = database.SubstitutePlaceholder(query, 1)
	// rows := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	// err = data.
	// if err != nil {
	// 	return
	// }
	return
}

func (repo teacherRepo) SelectAll(ctx context.Context, parameter *teacherController.TeacherParams) (data []teacherEntity.Teacher, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := teacherEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + teacherEntity.GroupStatement +
		` ORDER BY def.id` + parameter.OrderBy

	// query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := teacherEntity.Teacher{}
		err = temp.ScanRows(rows, nil)
		if err != sql.ErrNoRows {
			return nil, err
		}
		data = append(data, temp)
	}

	return
}

func (repo teacherRepo) FindOne(ctx context.Context, params *teacherController.TeacherParams) (data teacherEntity.Teacher, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := teacherEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + teacherEntity.GroupStatement +
		` ORDER BY def.id` + params.OrderBy

	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows {
		return
	}
	return
}

func (repo teacherRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != sql.ErrNoRows {
		return
	}
	return
}

// create
func (repo teacherRepo) Create(ctx context.Context, tx *sqlx.Tx, params *teacherEntity.Teacher) (err error) {
	queries := InsertTeacher

	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, queries, uuidRandom,
		params.FirstName,
		params.LastName,
		params.Email,
		params.Nik,
		params.PlaceOfBirth,
		params.DateOfBirth,
		params.Phone,
		params.Address,
		params.Gender,
		params.Religion,
		params.Image,
		1, // 1, status
		1, // 1, is_active
		params.ProvinceID,
		params.CityID,
		params.SubdistrictID,
		params.WardID,
		params.RT,
		params.RW,
	)
	if err != nil {
		return err
	}
	return
}
