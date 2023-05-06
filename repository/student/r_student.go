package student

import (
	"context"
	"database/sql"

	studentRequest "go-management-auth-school/controller/student"
	studentEntity "go-management-auth-school/entity/student"

	// "go-management-auth-school/helper/database"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewStudentRepo(dbMaster, dbSlave *sqlx.DB) *studentRepo {
	return &studentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo studentRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	return
}

func (repo studentRepo) SelectAll(ctx context.Context, parameter *studentRequest.StudentParams) (data []studentEntity.Student, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := studentEntity.SelectStudentStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + studentEntity.GroupStatement +
		` ORDER BY def.id` + parameter.OrderBy

	// query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows && err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := studentEntity.Student{}
		err = temp.ScanRows(rows, nil)
		if err != sql.ErrNoRows && err != nil {
			return
		}
		data = append(data, temp)
	}

	err = rows.Err()
	if err != sql.ErrNoRows && err != nil {
		return
	}

	return
}

func (repo studentRepo) FindOne(ctx context.Context, parameter *studentRequest.StudentParams) (data studentEntity.Student, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := studentEntity.SelectStudentStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment

	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows && err != nil {
		return
	}

	return
}

func (repo studentRepo) Create(ctx context.Context, tx *sqlx.Tx, input *studentEntity.Student) (res string, err error) {
	query := InsertStudent
	//convert string to date
	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, query,
		uuidRandom,
		input.FirstName,
		input.LastName,
		input.Email,
		input.Nisn,
		input.Nis,
		input.Nik,
		input.PlaceOfBirth,
		input.DateOfBirth,
		input.Phone,
		input.Address,
		input.Gender,
		input.Religion,
		input.Image,
		input.Status,
		input.IsActive,
		input.ProvinceID,
		input.CityID,
		input.SubdistrictID,
		input.WardID,
		input.Rt,
		input.Rw,
		input.Height,
		input.Weight,
		input.BloodType,
		input.Disability,
		input.DisabilityInfo,
		input.JoinDate,
		input.Details,
	)
	if err != nil {
		return
	}

	return
}
