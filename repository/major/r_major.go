package major

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	majorController "go-management-auth-school/controller/major"
	majorEntity "go-management-auth-school/entity/major"
)

type majorRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMajorRepo(dbMaster, dbSlave *sqlx.DB) *majorRepo {
	return &majorRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo majorRepo) buildingParams(ctx context.Context, parameter *majorController.MajorParams) (conditionString string, conditionParam []interface{}) {

	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}

	return
}

func (repo majorRepo) FindAll(ctx context.Context, params *majorController.MajorParams) (data []majorEntity.Major, err error) {
	// build query here
	return
}

func (repo majorRepo) SelectAll(ctx context.Context, parameter *majorController.MajorParams) (data []majorEntity.Major, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := majorEntity.Select + ` WHERE def.deleted_at IS NULL ` + whereStatment

	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d := majorEntity.Major{}
		err = d.ScanRows(rows, nil)
		if err != sql.ErrNoRows {
			return nil, err
		}
		data = append(data, d)
	}

	err = rows.Err()
	if err != sql.ErrNoRows {
		return
	}

	return
}

func (repo majorRepo) FindOne(ctx context.Context, params *majorController.MajorParams) (data majorEntity.Major, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := majorEntity.Select + ` WHERE def.deleted_at IS NULL ` + whereStatment

	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows {
		return data, err
	}

	return
}

func (repo majorRepo) Create(ctx context.Context, tx *sqlx.Tx, params *majorEntity.Major) (err error) {
	queries := InsertMajor
	_, err = tx.ExecContext(ctx, queries, params.Name)
	if err != nil {
		return err
	}
	return
}

func (repo majorRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	return
}
